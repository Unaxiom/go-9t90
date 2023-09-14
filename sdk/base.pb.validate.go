// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: base.proto

package sdk

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

// define the regex for a UUID once up-front
var _base_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on Empty with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Empty) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Empty with the rules defined in the
// proto definition for this message. If any rules are violated, the result is
// a list of violation errors wrapped in EmptyMultiError, or nil if none found.
func (m *Empty) ValidateAll() error {
	return m.validate(true)
}

func (m *Empty) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return EmptyMultiError(errors)
	}

	return nil
}

// EmptyMultiError is an error wrapping multiple validation errors returned by
// Empty.ValidateAll() if the designated constraints aren't met.
type EmptyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m EmptyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m EmptyMultiError) AllErrors() []error { return m }

// EmptyValidationError is the validation error returned by Empty.Validate if
// the designated constraints aren't met.
type EmptyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e EmptyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e EmptyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e EmptyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e EmptyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e EmptyValidationError) ErrorName() string { return "EmptyValidationError" }

// Error satisfies the builtin error interface
func (e EmptyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sEmpty.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = EmptyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = EmptyValidationError{}

// Validate checks the field values on BooleanResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *BooleanResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on BooleanResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// BooleanResponseMultiError, or nil if none found.
func (m *BooleanResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *BooleanResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Value

	if len(errors) > 0 {
		return BooleanResponseMultiError(errors)
	}

	return nil
}

// BooleanResponseMultiError is an error wrapping multiple validation errors
// returned by BooleanResponse.ValidateAll() if the designated constraints
// aren't met.
type BooleanResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m BooleanResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m BooleanResponseMultiError) AllErrors() []error { return m }

// BooleanResponseValidationError is the validation error returned by
// BooleanResponse.Validate if the designated constraints aren't met.
type BooleanResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e BooleanResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e BooleanResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e BooleanResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e BooleanResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e BooleanResponseValidationError) ErrorName() string { return "BooleanResponseValidationError" }

// Error satisfies the builtin error interface
func (e BooleanResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sBooleanResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = BooleanResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = BooleanResponseValidationError{}

// Validate checks the field values on CountRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CountRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CountRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CountRequestMultiError, or
// nil if none found.
func (m *CountRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CountRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for IsActive

	if len(errors) > 0 {
		return CountRequestMultiError(errors)
	}

	return nil
}

// CountRequestMultiError is an error wrapping multiple validation errors
// returned by CountRequest.ValidateAll() if the designated constraints aren't met.
type CountRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CountRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CountRequestMultiError) AllErrors() []error { return m }

// CountRequestValidationError is the validation error returned by
// CountRequest.Validate if the designated constraints aren't met.
type CountRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CountRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CountRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CountRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CountRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CountRequestValidationError) ErrorName() string { return "CountRequestValidationError" }

// Error satisfies the builtin error interface
func (e CountRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCountRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CountRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CountRequestValidationError{}

// Validate checks the field values on CountResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *CountResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CountResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in CountResponseMultiError, or
// nil if none found.
func (m *CountResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CountResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Count

	if len(errors) > 0 {
		return CountResponseMultiError(errors)
	}

	return nil
}

// CountResponseMultiError is an error wrapping multiple validation errors
// returned by CountResponse.ValidateAll() if the designated constraints
// aren't met.
type CountResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CountResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CountResponseMultiError) AllErrors() []error { return m }

// CountResponseValidationError is the validation error returned by
// CountResponse.Validate if the designated constraints aren't met.
type CountResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CountResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CountResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CountResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CountResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CountResponseValidationError) ErrorName() string { return "CountResponseValidationError" }

// Error satisfies the builtin error interface
func (e CountResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCountResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CountResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CountResponseValidationError{}

// Validate checks the field values on URLResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *URLResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on URLResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in URLResponseMultiError, or
// nil if none found.
func (m *URLResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *URLResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Url

	if len(errors) > 0 {
		return URLResponseMultiError(errors)
	}

	return nil
}

// URLResponseMultiError is an error wrapping multiple validation errors
// returned by URLResponse.ValidateAll() if the designated constraints aren't met.
type URLResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m URLResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m URLResponseMultiError) AllErrors() []error { return m }

// URLResponseValidationError is the validation error returned by
// URLResponse.Validate if the designated constraints aren't met.
type URLResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e URLResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e URLResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e URLResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e URLResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e URLResponseValidationError) ErrorName() string { return "URLResponseValidationError" }

// Error satisfies the builtin error interface
func (e URLResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sURLResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = URLResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = URLResponseValidationError{}

// Validate checks the field values on Metadata with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Metadata) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Metadata with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in MetadataMultiError, or nil
// if none found.
func (m *Metadata) ValidateAll() error {
	return m.validate(true)
}

func (m *Metadata) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Uuid

	// no validation rules for IsActive

	// no validation rules for CreatedAt

	// no validation rules for ModifiedAt

	if len(errors) > 0 {
		return MetadataMultiError(errors)
	}

	return nil
}

// MetadataMultiError is an error wrapping multiple validation errors returned
// by Metadata.ValidateAll() if the designated constraints aren't met.
type MetadataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m MetadataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m MetadataMultiError) AllErrors() []error { return m }

// MetadataValidationError is the validation error returned by
// Metadata.Validate if the designated constraints aren't met.
type MetadataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e MetadataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e MetadataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e MetadataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e MetadataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e MetadataValidationError) ErrorName() string { return "MetadataValidationError" }

// Error satisfies the builtin error interface
func (e MetadataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sMetadata.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = MetadataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = MetadataValidationError{}

// Validate checks the field values on Identifier with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *Identifier) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on Identifier with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in IdentifierMultiError, or
// nil if none found.
func (m *Identifier) ValidateAll() error {
	return m.validate(true)
}

func (m *Identifier) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if err := m._validateUuid(m.GetUuid()); err != nil {
		err = IdentifierValidationError{
			field:  "Uuid",
			reason: "value must be a valid UUID",
			cause:  err,
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return IdentifierMultiError(errors)
	}

	return nil
}

func (m *Identifier) _validateUuid(uuid string) error {
	if matched := _base_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// IdentifierMultiError is an error wrapping multiple validation errors
// returned by Identifier.ValidateAll() if the designated constraints aren't met.
type IdentifierMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m IdentifierMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m IdentifierMultiError) AllErrors() []error { return m }

// IdentifierValidationError is the validation error returned by
// Identifier.Validate if the designated constraints aren't met.
type IdentifierValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e IdentifierValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e IdentifierValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e IdentifierValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e IdentifierValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e IdentifierValidationError) ErrorName() string { return "IdentifierValidationError" }

// Error satisfies the builtin error interface
func (e IdentifierValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sIdentifier.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = IdentifierValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = IdentifierValidationError{}

// Validate checks the field values on SearchKeyRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *SearchKeyRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on SearchKeyRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// SearchKeyRequestMultiError, or nil if none found.
func (m *SearchKeyRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *SearchKeyRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for SearchKey

	if len(errors) > 0 {
		return SearchKeyRequestMultiError(errors)
	}

	return nil
}

// SearchKeyRequestMultiError is an error wrapping multiple validation errors
// returned by SearchKeyRequest.ValidateAll() if the designated constraints
// aren't met.
type SearchKeyRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m SearchKeyRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m SearchKeyRequestMultiError) AllErrors() []error { return m }

// SearchKeyRequestValidationError is the validation error returned by
// SearchKeyRequest.Validate if the designated constraints aren't met.
type SearchKeyRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchKeyRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchKeyRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchKeyRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchKeyRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchKeyRequestValidationError) ErrorName() string { return "SearchKeyRequestValidationError" }

// Error satisfies the builtin error interface
func (e SearchKeyRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchKeyRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchKeyRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchKeyRequestValidationError{}
