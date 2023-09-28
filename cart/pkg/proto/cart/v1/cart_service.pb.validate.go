// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: cart_service.proto

package cart

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

// Validate checks the field values on ListItem with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListItem) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListItem with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListItemMultiError, or nil
// if none found.
func (m *ListItem) ValidateAll() error {
	return m.validate(true)
}

func (m *ListItem) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Sku

	// no validation rules for Count

	// no validation rules for Name

	// no validation rules for Price

	if len(errors) > 0 {
		return ListItemMultiError(errors)
	}

	return nil
}

// ListItemMultiError is an error wrapping multiple validation errors returned
// by ListItem.ValidateAll() if the designated constraints aren't met.
type ListItemMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListItemMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListItemMultiError) AllErrors() []error { return m }

// ListItemValidationError is the validation error returned by
// ListItem.Validate if the designated constraints aren't met.
type ListItemValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListItemValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListItemValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListItemValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListItemValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListItemValidationError) ErrorName() string { return "ListItemValidationError" }

// Error satisfies the builtin error interface
func (e ListItemValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListItem.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListItemValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListItemValidationError{}

// Validate checks the field values on ListRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListRequestMultiError, or
// nil if none found.
func (m *ListRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ListRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for User

	if len(errors) > 0 {
		return ListRequestMultiError(errors)
	}

	return nil
}

// ListRequestMultiError is an error wrapping multiple validation errors
// returned by ListRequest.ValidateAll() if the designated constraints aren't met.
type ListRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListRequestMultiError) AllErrors() []error { return m }

// ListRequestValidationError is the validation error returned by
// ListRequest.Validate if the designated constraints aren't met.
type ListRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListRequestValidationError) ErrorName() string { return "ListRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListRequestValidationError{}

// Validate checks the field values on ListResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ListResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ListResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ListResponseMultiError, or
// nil if none found.
func (m *ListResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ListResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	for idx, item := range m.GetItems() {
		_, _ = idx, item

		if all {
			switch v := interface{}(item).(type) {
			case interface{ ValidateAll() error }:
				if err := v.ValidateAll(); err != nil {
					errors = append(errors, ListResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			case interface{ Validate() error }:
				if err := v.Validate(); err != nil {
					errors = append(errors, ListResponseValidationError{
						field:  fmt.Sprintf("Items[%v]", idx),
						reason: "embedded message failed validation",
						cause:  err,
					})
				}
			}
		} else if v, ok := interface{}(item).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return ListResponseValidationError{
					field:  fmt.Sprintf("Items[%v]", idx),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	// no validation rules for TotalPrice

	if len(errors) > 0 {
		return ListResponseMultiError(errors)
	}

	return nil
}

// ListResponseMultiError is an error wrapping multiple validation errors
// returned by ListResponse.ValidateAll() if the designated constraints aren't met.
type ListResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ListResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ListResponseMultiError) AllErrors() []error { return m }

// ListResponseValidationError is the validation error returned by
// ListResponse.Validate if the designated constraints aren't met.
type ListResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListResponseValidationError) ErrorName() string { return "ListResponseValidationError" }

// Error satisfies the builtin error interface
func (e ListResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListResponseValidationError{}

// Validate checks the field values on ClearRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ClearRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ClearRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ClearRequestMultiError, or
// nil if none found.
func (m *ClearRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ClearRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for User

	if len(errors) > 0 {
		return ClearRequestMultiError(errors)
	}

	return nil
}

// ClearRequestMultiError is an error wrapping multiple validation errors
// returned by ClearRequest.ValidateAll() if the designated constraints aren't met.
type ClearRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ClearRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ClearRequestMultiError) AllErrors() []error { return m }

// ClearRequestValidationError is the validation error returned by
// ClearRequest.Validate if the designated constraints aren't met.
type ClearRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClearRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClearRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClearRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClearRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClearRequestValidationError) ErrorName() string { return "ClearRequestValidationError" }

// Error satisfies the builtin error interface
func (e ClearRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClearRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClearRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClearRequestValidationError{}

// Validate checks the field values on ClearResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ClearResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ClearResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ClearResponseMultiError, or
// nil if none found.
func (m *ClearResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ClearResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ClearResponseMultiError(errors)
	}

	return nil
}

// ClearResponseMultiError is an error wrapping multiple validation errors
// returned by ClearResponse.ValidateAll() if the designated constraints
// aren't met.
type ClearResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ClearResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ClearResponseMultiError) AllErrors() []error { return m }

// ClearResponseValidationError is the validation error returned by
// ClearResponse.Validate if the designated constraints aren't met.
type ClearResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ClearResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ClearResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ClearResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ClearResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ClearResponseValidationError) ErrorName() string { return "ClearResponseValidationError" }

// Error satisfies the builtin error interface
func (e ClearResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sClearResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ClearResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ClearResponseValidationError{}

// Validate checks the field values on ItemAddRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *ItemAddRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ItemAddRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in ItemAddRequestMultiError,
// or nil if none found.
func (m *ItemAddRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ItemAddRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for User

	// no validation rules for Sku

	if m.GetCount() < 1 {
		err := ItemAddRequestValidationError{
			field:  "Count",
			reason: "value must be greater than or equal to 1",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return ItemAddRequestMultiError(errors)
	}

	return nil
}

// ItemAddRequestMultiError is an error wrapping multiple validation errors
// returned by ItemAddRequest.ValidateAll() if the designated constraints
// aren't met.
type ItemAddRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ItemAddRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ItemAddRequestMultiError) AllErrors() []error { return m }

// ItemAddRequestValidationError is the validation error returned by
// ItemAddRequest.Validate if the designated constraints aren't met.
type ItemAddRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ItemAddRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ItemAddRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ItemAddRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ItemAddRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ItemAddRequestValidationError) ErrorName() string { return "ItemAddRequestValidationError" }

// Error satisfies the builtin error interface
func (e ItemAddRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sItemAddRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ItemAddRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ItemAddRequestValidationError{}

// Validate checks the field values on ItemAddResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ItemAddResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ItemAddResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ItemAddResponseMultiError, or nil if none found.
func (m *ItemAddResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ItemAddResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ItemAddResponseMultiError(errors)
	}

	return nil
}

// ItemAddResponseMultiError is an error wrapping multiple validation errors
// returned by ItemAddResponse.ValidateAll() if the designated constraints
// aren't met.
type ItemAddResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ItemAddResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ItemAddResponseMultiError) AllErrors() []error { return m }

// ItemAddResponseValidationError is the validation error returned by
// ItemAddResponse.Validate if the designated constraints aren't met.
type ItemAddResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ItemAddResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ItemAddResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ItemAddResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ItemAddResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ItemAddResponseValidationError) ErrorName() string { return "ItemAddResponseValidationError" }

// Error satisfies the builtin error interface
func (e ItemAddResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sItemAddResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ItemAddResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ItemAddResponseValidationError{}

// Validate checks the field values on ItemDeleteRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ItemDeleteRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ItemDeleteRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ItemDeleteRequestMultiError, or nil if none found.
func (m *ItemDeleteRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *ItemDeleteRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for User

	// no validation rules for Sku

	if len(errors) > 0 {
		return ItemDeleteRequestMultiError(errors)
	}

	return nil
}

// ItemDeleteRequestMultiError is an error wrapping multiple validation errors
// returned by ItemDeleteRequest.ValidateAll() if the designated constraints
// aren't met.
type ItemDeleteRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ItemDeleteRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ItemDeleteRequestMultiError) AllErrors() []error { return m }

// ItemDeleteRequestValidationError is the validation error returned by
// ItemDeleteRequest.Validate if the designated constraints aren't met.
type ItemDeleteRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ItemDeleteRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ItemDeleteRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ItemDeleteRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ItemDeleteRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ItemDeleteRequestValidationError) ErrorName() string {
	return "ItemDeleteRequestValidationError"
}

// Error satisfies the builtin error interface
func (e ItemDeleteRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sItemDeleteRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ItemDeleteRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ItemDeleteRequestValidationError{}

// Validate checks the field values on ItemDeleteResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the first error encountered is returned, or nil if there are no violations.
func (m *ItemDeleteResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ItemDeleteResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ItemDeleteResponseMultiError, or nil if none found.
func (m *ItemDeleteResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *ItemDeleteResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return ItemDeleteResponseMultiError(errors)
	}

	return nil
}

// ItemDeleteResponseMultiError is an error wrapping multiple validation errors
// returned by ItemDeleteResponse.ValidateAll() if the designated constraints
// aren't met.
type ItemDeleteResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ItemDeleteResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ItemDeleteResponseMultiError) AllErrors() []error { return m }

// ItemDeleteResponseValidationError is the validation error returned by
// ItemDeleteResponse.Validate if the designated constraints aren't met.
type ItemDeleteResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ItemDeleteResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ItemDeleteResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ItemDeleteResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ItemDeleteResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ItemDeleteResponseValidationError) ErrorName() string {
	return "ItemDeleteResponseValidationError"
}

// Error satisfies the builtin error interface
func (e ItemDeleteResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sItemDeleteResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ItemDeleteResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ItemDeleteResponseValidationError{}

// Validate checks the field values on CheckoutRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CheckoutRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CheckoutRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CheckoutRequestMultiError, or nil if none found.
func (m *CheckoutRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *CheckoutRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for User

	if len(errors) > 0 {
		return CheckoutRequestMultiError(errors)
	}

	return nil
}

// CheckoutRequestMultiError is an error wrapping multiple validation errors
// returned by CheckoutRequest.ValidateAll() if the designated constraints
// aren't met.
type CheckoutRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CheckoutRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CheckoutRequestMultiError) AllErrors() []error { return m }

// CheckoutRequestValidationError is the validation error returned by
// CheckoutRequest.Validate if the designated constraints aren't met.
type CheckoutRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckoutRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckoutRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckoutRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckoutRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckoutRequestValidationError) ErrorName() string { return "CheckoutRequestValidationError" }

// Error satisfies the builtin error interface
func (e CheckoutRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckoutRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckoutRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckoutRequestValidationError{}

// Validate checks the field values on CheckoutResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *CheckoutResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on CheckoutResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// CheckoutResponseMultiError, or nil if none found.
func (m *CheckoutResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *CheckoutResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for OrderId

	if len(errors) > 0 {
		return CheckoutResponseMultiError(errors)
	}

	return nil
}

// CheckoutResponseMultiError is an error wrapping multiple validation errors
// returned by CheckoutResponse.ValidateAll() if the designated constraints
// aren't met.
type CheckoutResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m CheckoutResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m CheckoutResponseMultiError) AllErrors() []error { return m }

// CheckoutResponseValidationError is the validation error returned by
// CheckoutResponse.Validate if the designated constraints aren't met.
type CheckoutResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CheckoutResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CheckoutResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CheckoutResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CheckoutResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CheckoutResponseValidationError) ErrorName() string { return "CheckoutResponseValidationError" }

// Error satisfies the builtin error interface
func (e CheckoutResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCheckoutResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CheckoutResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CheckoutResponseValidationError{}
