package dao

import (
	"context"
	"demo/internal/user/dao/query"
	"demo/internal/user/dto"
	"gorm.io/gorm"
)

type UserSaver struct {
	db *gorm.DB
}

func NewUserSaver(db *gorm.DB) *UserSaver {
	return &UserSaver{db: db}
}

func (f *UserSaver) Create(ctx context.Context, user dto.User) error {
	result := f.db.Table(query.UserTable).WithContext(ctx).Create(&user)

	return result.Error
}

func (f *UserSaver) Update(ctx context.Context, user dto.User) error {
	result := f.db.Table(query.UserTable).WithContext(ctx).Save(&user)

	return result.Error
}
