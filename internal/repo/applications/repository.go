package applications

import (
	"context"

	"gorm.io/gorm"
	"oatkup.sys/internal/entity"
)

type Repository interface {
	Create(ctx context.Context, app *entity.Application) error
}

type gormRepo struct{ db *gorm.DB }

func NewGormRepository(db *gorm.DB) Repository { return &gormRepo{db: db} }

func (r *gormRepo) Create(ctx context.Context, app *entity.Application) error {
	return r.db.WithContext(ctx).Create(app).Error
}
