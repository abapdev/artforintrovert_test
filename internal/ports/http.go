package ports

import (
	"artforintrovert_test/internal/domain/models"
	"context"
)

type Service interface {
	List(ctx context.Context) ([]models.Data, error)
	Update(ctx context.Context, data *models.Data) error
	Delete(ctx context.Context, data *models.Data) error
}
