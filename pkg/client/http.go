package client

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func Do(ctx context.Context, endpoint, method string, req map[string]map[string]string) (map[string]interface{}, error) {
	buff := bytes.Buffer{}
	if bodyMap, isBodyMapExists := req["body"]; isBodyMapExists {
		b, err := json.Marshal(bodyMap)
		if err != nil {
			return nil, err
		}
		_, err = buff.Write(b)
		if err != nil {
			return nil, err
		}
	}
	queryParam := url.Values{}
	if queryParamMap, isQueryParamMapExist := req["query"]; isQueryParamMapExist {
		for k, v := range queryParamMap {
			queryParam.Set(k, v)
		}
	}

	request, err := http.NewRequestWithContext(ctx, method, fmt.Sprintf("%s?%s", endpoint, queryParam.Encode()), &buff)
	if err != nil {
		return nil, err
	}
	if headerMap, isHeaderMapExist := req["header"]; isHeaderMapExist {
		for k, v := range headerMap {
			request.Header.Set(k, v)
		}
	}

	client := http.Client{}
	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	defer response.Body.Close()

	body, err := io.ReadAll(response.Body)
	if err != nil {
		return nil, fmt.Errorf("failed on unmarshal response: %v", err)
	}

	if response.StatusCode < http.StatusOK && response.StatusCode >= http.StatusMultipleChoices {
		return nil, fmt.Errorf("failed: %v", string(body))
	}
	var resp map[string]interface{}
	err = json.Unmarshal(body, &resp)
	if err != nil {
		return nil, fmt.Errorf("failed on unmarshal response: %v", err)
	}

	return resp, nil

}
