package applications

import (
	"context"

	"github.com/google/uuid"
	"oatkup.sys/internal/entity"
	"oatkup.sys/internal/repo/applications"
)

type CreateUsecase struct{ repo applications.Repository }

func NewCreateUsecase(r applications.Repository) *CreateUsecase {
	return &CreateUsecase{repo: r}
}

func (uc *CreateUsecase) Apply(ctx context.Context, req *ApplicationRequest) (*ApplicationModel, error) {
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
