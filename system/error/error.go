package error

import "fmt"

// SystemError type
type SystemError struct {
	StatusCode int
	Err        error
}

// Error string
func (r *SystemError) Error() string {
	return fmt.Sprintf("Status %d: Error => %v", r.StatusCode, r.Err)
}
