package postgres

import (
	"context"
	"database/sql"

	"github.com/pkg/errors"

	"github.com/hereisajvi/noty/internal/model"
)

var ErrNoMessage = errors.New("no message")

type MessageRepository struct {
	postgres Postgres
}

func NewMessageRepository(postgres Postgres) *MessageRepository {
	return &MessageRepository{
		postgres: postgres,
	}
}

func (r *MessageRepository) Create(ctx context.Context, message model.MessageDataAccessObject) (model.MessageBusinessObject, error) {
	query := `
	INSERT INTO messages (
		pid,
		text,
		created_at,
		updated_at
	)
	VALUES (
		:pid,
		:text,
		:created_at,
		:updated_at
	)
	RETURNING
		pid,
		text,
		created_at,
		updated_at
	`

	rows, err := r.postgres.NamedQueryContext(ctx, query, message)
	if err != nil {
		return model.MessageBusinessObject{}, errors.Wrap(err, "failed to execute insert message query")
	}

	for rows.Next() {
		err := rows.StructScan(&message)
		if err != nil {
			return model.MessageBusinessObject{}, errors.Wrap(err, "failed to scan result to struct")
		}
	}

	return model.MessageBusinessObject(message), nil
}

func (r *MessageRepository) Update(ctx context.Context, id string, message model.MessageDataAccessObject) (model.MessageBusinessObject, error) {
	query := `
	UPDATE messages
	SET
		text = :text,
		updated_at = :updated_at
	WHERE
		pid = :pid
	RETURNING
		*
	`

	rows, err := r.postgres.NamedQueryContext(ctx, query, map[string]interface{}{
		"text":       message.Text,
		"updated_at": message.UpdatedAt,
		"pid":        id,
	})
	if err != nil {
		return model.MessageBusinessObject{}, errors.Wrap(err, "failed to execute update message query")
	}

	for rows.Next() {
		err := rows.StructScan(&message)
		if err != nil {
			return model.MessageBusinessObject{}, errors.Wrap(err, "failed to scan result to struct")
		}
	}

	return model.MessageBusinessObject(message), nil
}

func (r *MessageRepository) Delete(ctx context.Context, id string) error {
	query := `DELETE FROM messages WHERE pid = $1
	`

	_, err := r.postgres.ExecContext(ctx, query, id)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return ErrNoMessage
		}

		return errors.Wrap(err, "failed to execute delete message query")
	}

	return nil
}

func (r *MessageRepository) List(ctx context.Context) ([]model.MessageBusinessObject, error) {
	query := `
	SELECT
		pid,
		text,
		created_at,
		updated_at
	FROM
		messages
	`

	var messages []model.MessageDataAccessObject

	err := r.postgres.SelectContext(ctx, &messages, query)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return nil, ErrNoMessage
		}

		return nil, errors.Wrap(err, "failed to execute select messages query")
	}

	return model.MapMessageDataAccessObjectListToMessageBusinessObjectList(messages), nil
}
