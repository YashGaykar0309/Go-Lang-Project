package dberrors

type ConflictError struct{}

func (e *ConflictError) Error() string {
	return "Attempted to create a record with existing key."
}

