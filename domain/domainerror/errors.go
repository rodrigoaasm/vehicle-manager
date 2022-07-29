package domainerror

type ErrorType string

const (
	TEST = "TEST"

	DATABASE     = "DATABASE"
	DEPENDENCY   = "DEPENDENCY"
	UNKNOWN      = "UNKNOWN"
	INVALID_DATA = "INVALID_DATA"
	NOT_FOUND    = "NOT_FOUND"
	CONFLICT     = "CONFLICT"
)

type DomainError struct {
	Type    ErrorType `json:"type"`
	Message string    `json:"message"`
}

// Error treatment
func New(eType ErrorType, message string) *DomainError {
	return &DomainError{
		Type:    eType,
		Message: message,
	}
}
