// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: protodep.proto

package protodep

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

// define the regex for a UUID once up-front
var _protodep_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Config with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Config) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetLocal()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConfigValidationError{
				field:  "Local",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	for idx, item := range m.GetImports() {
		_, _ = idx, item

		if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ConfigValidationError{
					field:  fmt.Sprintf("Imports[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

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

// Validate checks the field values on Import with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Import) Validate() error {
	if m == nil {
		return nil
	}

	switch m.ImportType.(type) {

	case *Import_GoMod:

		if v, ok := interface{}(m.GetGoMod()).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ImportValidationError{
					field:  "GoMod",
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	return nil
}

// ImportValidationError is the validation error returned by Import.Validate if
// the designated constraints aren't met.
type ImportValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ImportValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ImportValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ImportValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ImportValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ImportValidationError) ErrorName() string { return "ImportValidationError" }

// Error satisfies the builtin error interface
func (e ImportValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sImport.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ImportValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ImportValidationError{}

// Validate checks the field values on Local with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *Local) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// LocalValidationError is the validation error returned by Local.Validate if
// the designated constraints aren't met.
type LocalValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LocalValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LocalValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LocalValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LocalValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LocalValidationError) ErrorName() string { return "LocalValidationError" }

// Error satisfies the builtin error interface
func (e LocalValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLocal.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LocalValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LocalValidationError{}

// Validate checks the field values on GoModImport with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *GoModImport) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Package

	return nil
}

// GoModImportValidationError is the validation error returned by
// GoModImport.Validate if the designated constraints aren't met.
type GoModImportValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GoModImportValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GoModImportValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GoModImportValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GoModImportValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GoModImportValidationError) ErrorName() string { return "GoModImportValidationError" }

// Error satisfies the builtin error interface
func (e GoModImportValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGoModImport.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GoModImportValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GoModImportValidationError{}
