package web

import "fmt"

// BindError - unexpected error occurred in binding data
func BindError(err error) error {
	return fmt.Errorf("failed to bind data: %v", err)
}

// ValidateError - unexpected error occurred in validating data
func ValidateError(err error) error {
	return fmt.Errorf("failed to validate data: %v", err)
}

// CreateError - unexpected error occurred in creating a new record
func CreateError(err error) error {
	return fmt.Errorf("failed to create a new record: %v", err)
}

// UpdateError - unexpected error occurred in creating a new record
func UpdateError(err error) error {
	return fmt.Errorf("failed to update record: %v", err)
}

// RemoveError - unexpected error occurred in removing a new record
func RemoveError(err error) error {
	return fmt.Errorf("failed to remove record: %v", err)
}

// FindError - unexpected error occurred in finding record
func FindError(err error) error {
	return fmt.Errorf("failed to find records: %v", err)
}
