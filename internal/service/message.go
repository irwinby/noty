package service

import (
	"context"
	"time"

	"github.com/pkg/errors"

	"github.com/hereisajvi/noty/internal/model"
	"github.com/hereisajvi/noty/internal/storage/postgres"
	"github.com/hereisajvi/noty/pkg/nanoid"
)

var ErrNoMessage = errors.New("no message")

type Message struct {
	repository MessageRepository
}

func NewMessage(repository MessageRepository) *Message {
	return &Message{
		repository: repository,
	}
}

func (m *Message) Create(ctx context.Context, message model.MessageBusinessObject) (model.MessageDataTransferObject, error) {
	id, err := nanoid.New()
	if err != nil {
		return model.MessageDataTransferObject{}, errors.Wrap(err, "failed to create nanoid")
	}

	now := time.Now()

	message.SetID(id)
	message.SetCreatedAt(now)
	message.SetUpdatedAt(now)

	message, err = m.repository.Create(ctx, model.MessageDataAccessObject(message))
	if err != nil {
		return model.MessageDataTransferObject{}, errors.Wrap(err, "failed to create message")
	}

	return model.MessageDataTransferObject(message), nil
}

func (m *Message) Update(ctx context.Context, id string, message model.MessageBusinessObject) (model.MessageDataTransferObject, error) {
	now := time.Now()

	message.SetUpdatedAt(now)

	message, err := m.repository.Update(ctx, id, model.MessageDataAccessObject(message))
	if err != nil {
		if errors.Is(err, postgres.ErrNoMessage) {
			return model.MessageDataTransferObject{}, ErrNoMessage
		}

		return model.MessageDataTransferObject{}, errors.Wrap(err, "failed to update message")
	}

	return model.MessageDataTransferObject(message), nil
}

func (m *Message) Delete(ctx context.Context, id string) error {
	err := m.repository.Delete(ctx, id)
	if err != nil {
		if errors.Is(err, postgres.ErrNoMessage) {
			return ErrNoMessage
		}

		return errors.Wrap(err, "failed to delete message")
	}

	return nil
}

func (m *Message) List(ctx context.Context) ([]model.MessageDataTransferObject, error) {
	messages, err := m.repository.List(ctx)
	if err != nil {
		if errors.Is(err, postgres.ErrNoMessage) {
			return nil, ErrNoMessage
		}

		return nil, errors.Wrap(err, "failed to get list of messages")
	}

	return model.MapMessageBusinessObjectListToMessageDataTransferObjectList(messages), nil
}
