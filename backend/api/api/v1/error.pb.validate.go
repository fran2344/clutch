// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/v1/error.proto

package apiv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = ptypes.DynamicAny{}
)

// Validate checks the field values on ErrorDetails with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ErrorDetails) Validate() error {
	if m == nil {
		return nil
	}

	for idx, item := range m.GetWrapped() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ErrorDetailsValidationError{
					field:  fmt.Sprintf("Wrapped[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ErrorDetailsValidationError is the validation error returned by
// ErrorDetails.Validate if the designated constraints aren't met.
type ErrorDetailsValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ErrorDetailsValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ErrorDetailsValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ErrorDetailsValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ErrorDetailsValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ErrorDetailsValidationError) ErrorName() string { return "ErrorDetailsValidationError" }

// Error satisfies the builtin error interface
func (e ErrorDetailsValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sErrorDetails.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ErrorDetailsValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ErrorDetailsValidationError{}
