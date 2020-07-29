package error

import (
	"fmt"
	"net/http"
)

// ForbiddenAccess error
func ForbiddenAccess() *SystemError {
	return CustomForbiddenAccess("Forbidden Access")
}

// CustomForbiddenAccess error
func CustomForbiddenAccess(message string) *SystemError {
	return &SystemError{StatusCode: http.StatusForbidden, Err: fmt.Errorf(message)}
}
