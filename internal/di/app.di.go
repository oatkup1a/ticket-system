package di

import (
	"gorm.io/gorm"

	"oatkup.sys/internal/application/usecase/applications"
	"oatkup.sys/internal/entity"
	repo "oatkup.sys/internal/repo/applications"
	"oatkup.sys/internal/server/rest/handler"
	"oatkup.sys/pkg/database"
)

type AppDeps struct {
	DB          *gorm.DB
	AppRepo     repo.Repository
	CreateAppUC *applications.CreateApplicationUsecase
	AppHandler  *handler.ApplicationHandler
}

func NewAppDeps(dsn string) (*AppDeps, error) {
	db, err := database.NewGormDB(database.Config{DSN: dsn})
	if err != nil {
		return nil, err
	}
	// Auto-migrate model(s)
	if err := db.AutoMigrate(&entity.Application{}); err != nil {
		return nil, err
	}

	repo := repo.NewGormRepository(db)
	createUC := applications.NewCreateApplicationUsecase(repo)
	h := handler.NewApplicationHandler(createUC)

	return &AppDeps{
		DB:          db,
		AppRepo:     repo,
		CreateAppUC: createUC,
		AppHandler:  h,
	}, nil
}
