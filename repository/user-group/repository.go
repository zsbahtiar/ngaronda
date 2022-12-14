package user_group

import (
	"context"
	"github.com/zsbahtiar/ngaronda/core/entity"
	"github.com/zsbahtiar/ngaronda/core/repository"
	"gorm.io/gorm"
)

type userGroup struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) repository.UserGroupRepository {
	return &userGroup{db: db}
}

func (u *userGroup) UpsertUserGroup(ctx context.Context, request *entity.UpsertUserGroupRequest) error {
	var err error
	tx := u.db.WithContext(ctx).Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()
	err = tx.Delete(userGroup{}, "group_id = ?", request.GroupID).Error
	if err != nil {
		return err
	}

	err = tx.Table(tableName).Create(parseUpsertUserGroupToDto(request)).Error
	if err != nil {
		return err
	}

	return tx.Commit().Error

}
