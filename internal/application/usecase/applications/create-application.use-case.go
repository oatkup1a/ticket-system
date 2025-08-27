package applications

import (
	"context"

	"github.com/google/uuid"
	"oatkup.sys/internal/entity"
	"oatkup.sys/internal/repo/applications"
)

type CreateApplicationUsecase struct{ repo applications.Repository }

func NewCreateApplicationUsecase(r applications.Repository) *CreateApplicationUsecase {
	return &CreateApplicationUsecase{repo: r}
}

func (uc *CreateApplicationUsecase) Apply(ctx context.Context, req *ApplicationRequest) (*ApplicationModel, error) {
	app := &entity.Application{
		ID:          uuid.NewString(),
		Name:        req.Name,
		Description: req.Description,
	}
	if err := uc.repo.Create(ctx, app); err != nil {
		return nil, err
	}
	return &ApplicationModel{
		ID:          app.ID,
		Name:        app.Name,
		Description: app.Description,
	}, nil
}
