package applications

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"oatkup.sys/internal/entity"
	"oatkup.sys/internal/repo/applications"
)

type CreateApplicationUsecase struct{ repo applications.Repository }

func NewCreateApplicationUsecase(r applications.Repository) *CreateApplicationUsecase {
	return &CreateApplicationUsecase{repo: r}
}

func (uc *CreateApplicationUsecase) Apply(ctx context.Context, req *ApplicationRequest) (*ApplicationModel, error) {
	const errPrefix = "[CreateApplicationUsecase][Apply]"

	app := &entity.Application{
		ID:          uuid.NewString(),
		Name:        req.Name,
		Description: req.Description,
	}
	if err := uc.repo.Create(ctx, app); err != nil {
		return nil, fmt.Errorf("%s: %w", errPrefix, err)
	}
	return &ApplicationModel{
		ID:          app.ID,
		Name:        app.Name,
		Description: app.Description,
	}, nil
}
