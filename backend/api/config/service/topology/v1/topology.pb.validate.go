// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: config/service/topology/v1/topology.proto

package topologyv1

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
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
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on Config with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Config) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Config with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in ConfigMultiError, or nil if none found.
func (m *Config) ValidateAll() error {
	return m.validate(true)
}

func (m *Config) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if all {
		switch v := interface{}(m.GetCache()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ConfigValidationError{
					field:  "Cache",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ConfigValidationError{
					field:  "Cache",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetCache()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfigValidationError{
				field:  "Cache",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ConfigMultiError(errors)
	}

	return nil
}

// ConfigMultiError is an error wrapping multiple validation errors returned by
// Config.ValidateAll() if the designated constraints aren't met.
type ConfigMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConfigMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConfigMultiError) AllErrors() []error { return m }

// ConfigValidationError is the validation error returned by Config.Validate if
// the designated constraints aren't met.
type ConfigValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConfigValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConfigValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConfigValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConfigValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConfigValidationError) ErrorName() string { return "ConfigValidationError" }

// Error satisfies the builtin error interface
func (e ConfigValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConfig.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConfigValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConfigValidationError{}

// Validate checks the field values on Cache with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Cache) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Cache with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in CacheMultiError, or nil if none found.
func (m *Cache) ValidateAll() error {
	return m.validate(true)
}

func (m *Cache) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if d := m.GetTtl(); d != nil {
		dur, err := d.AsDuration(), d.CheckValid()
		if err != nil {
			err = CacheValidationError{
				field:  "Ttl",
				reason: "value is not a valid duration",
				cause:  err,
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		} else {

			gt := time.Duration(1*time.Second + 0*time.Nanosecond)

			if dur <= gt {
				err := CacheValidationError{
					field:  "Ttl",
					reason: "value must be greater than 1s",
				}
				if !all {
					return err
				}
				errors = append(errors, err)
			}

		}
	}

	// no validation rules for BatchInsertSize

	if all {
		switch v := interface{}(m.GetBatchInsertFlush()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, CacheValidationError{
					field:  "BatchInsertFlush",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, CacheValidationError{
					field:  "BatchInsertFlush",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetBatchInsertFlush()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CacheValidationError{
				field:  "BatchInsertFlush",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return CacheMultiError(errors)
	}

	return nil
}

// CacheMultiError is an error wrapping multiple validation errors returned by
// Cache.ValidateAll() if the designated constraints aren't met.
type CacheMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CacheMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CacheMultiError) AllErrors() []error { return m }

// CacheValidationError is the validation error returned by Cache.Validate if
// the designated constraints aren't met.
type CacheValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CacheValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CacheValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CacheValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CacheValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CacheValidationError) ErrorName() string { return "CacheValidationError" }

// Error satisfies the builtin error interface
func (e CacheValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCache.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CacheValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CacheValidationError{}
