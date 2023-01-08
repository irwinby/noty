package model

import (
	"time"
)

type (
	MessageDataTransferObject struct {
		ID        string    `json:"id"`
		Text      string    `json:"text"`
		CreatedAt time.Time `json:"created_at"`
		UpdatedAt time.Time `json:"updated_at"`
	}

	MessageBusinessObject struct {
		ID        string
		Text      string
		CreatedAt time.Time
		UpdatedAt time.Time
	}

	MessageDataAccessObject struct {
		ID        string    `db:"pid"`
		Text      string    `db:"text"`
		CreatedAt time.Time `db:"created_at"`
		UpdatedAt time.Time `db:"updated_at"`
	}
)

func (o *MessageBusinessObject) SetID(id string) {
	o.ID = id
}

func (o *MessageBusinessObject) SetCreatedAt(time time.Time) {
	o.CreatedAt = time
}

func (o *MessageBusinessObject) SetUpdatedAt(time time.Time) {
	o.UpdatedAt = time
}

func MapMessageBusinessObjectListToMessageDataTransferObjectList(messages []MessageBusinessObject) []MessageDataTransferObject {
	_messages := make([]MessageDataTransferObject, 0, len(messages))

	for _, message := range messages {
		_messages = append(_messages, MessageDataTransferObject(message))
	}

	return _messages
}

func MapMessageDataAccessObjectListToMessageBusinessObjectList(messages []MessageDataAccessObject) []MessageBusinessObject {
	_messages := make([]MessageBusinessObject, 0, len(messages))

	for _, message := range messages {
		_messages = append(_messages, MessageBusinessObject(message))
	}

	return _messages
}
