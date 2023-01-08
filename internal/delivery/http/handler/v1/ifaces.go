package v1

import (
	"context"

	"github.com/hereisajvi/noty/internal/model"
)

type MessageService interface {
	Create(ctx context.Context, message model.MessageBusinessObject) (model.MessageDataTransferObject, error)
	Update(ctx context.Context, id string, message model.MessageBusinessObject) (model.MessageDataTransferObject, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]model.MessageDataTransferObject, error)
}
