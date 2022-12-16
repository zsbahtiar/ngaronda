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
	err = tx.Delete(userGroup{}, "group_id = ?", request.GroupID).
		Error
	if err != nil {
		return err
	}

	err = tx.Table(tableName).
		Create(parseUpsertUserGroupToDto(request)).
		Error
	if err != nil {
		return err
	}

	return tx.Commit().Error

}

func (u *userGroup) GetUserGroupsByCronType(ctx context.Context, cronType entity.CronType) ([]*entity.UserGroup, error) {
	var userGroups []*UserGroup

	err := u.db.
		WithContext(ctx).
		Table(tableName).
		Find(&userGroups, "cron_type = ?", cronType).Error

	return parseToUserGroupsEntity(userGroups), err
}

func (u *userGroup) UpdateUserGroupsRank(ctx context.Context, ugs []*entity.UserGroup) error {
	var err error
	tx := u.db.WithContext(ctx).Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		}
	}()

	for _, ug := range ugs {
		err = tx.Table(tableName).
			Where("serial = ?", ug.Serial).
			Update("current_rank", ug.CurrentRank).
			Error
		if err != nil {
			return err
		}
	}

	return tx.Commit().Error
}
