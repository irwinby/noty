package service

import (
	"context"

	"github.com/hereisajvi/noty/internal/model"
)

type MessageRepository interface {
	Create(ctx context.Context, message model.MessageDataAccessObject) (model.MessageBusinessObject, error)
	Update(ctx context.Context, id string, messgae model.MessageDataAccessObject) (model.MessageBusinessObject, error)
	Delete(ctx context.Context, id string) error
	List(ctx context.Context) ([]model.MessageBusinessObject, error)
}
