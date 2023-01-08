package nanoid

import (
	"github.com/jaevor/go-nanoid"
	"github.com/pkg/errors"
)

const alphabet = "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz"

const length = 12

func New() (string, error) {
	fn, err := nanoid.CustomUnicode(alphabet, length)
	if err != nil {
		return "", errors.Wrap(err, "failed to create nanoid with custom alphabet and length")
	}

	return fn(), nil
}
