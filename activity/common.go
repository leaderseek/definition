package activity

import (
	"database/sql"

	"github.com/friendsofgo/errors"
)

func ErrorWithRollback(tx *sql.Tx, err error, message string) error {
	if rErr := tx.Rollback(); rErr != nil {
		return errors.Wrapf(err, "%s, rollback error = %v", message, rErr)
	}

	return errors.Wrap(err, message)
}

func NewErrorWithRollback(tx *sql.Tx, message string) error {
	if rErr := tx.Rollback(); rErr != nil {
		return errors.Wrapf(rErr, "%s, rollback error", message)
	}

	return errors.New(message)
}
