package ports

import (
	"artforintrovert_test/internal/domain/models"
	"context"
)

type Storage interface {
	//Получаем из БД записи(все или по фильтру name)
	GetList(ctx context.Context, name string) ([]models.Data, error)
	//Обновляем запись в БД
	UpdateData(ctx context.Context, data *models.DataJSON) error
	//Удаляем запись в БД
	DeleteData(ctx context.Context, data *models.DataJSON) error
}
