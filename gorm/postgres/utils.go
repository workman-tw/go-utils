package postgres

import (
	"fmt"

	"gorm.io/gorm"
)

// HandleTransaction - if an error occurred, rollback a subset of operations
func HandleTransaction(tx *gorm.DB) {
	if r := recover(); r != nil {
		tx.Rollback()
	}
}

// Rollback - handle rollback error
func Rollback(tx *gorm.DB, unexpectedErr error) error {
	rollBackErr := tx.Rollback().Error

	if unexpectedErr == nil {
		return rollBackErr
	}

	if rollBackErr == nil {
		return unexpectedErr
	}

	return fmt.Errorf("unexpected error: %v, rollback error: %v", unexpectedErr, rollBackErr)
}
