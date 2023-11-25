package exerror

import "context"

type ErrorMeta struct {
	IsImportant bool
	IsTemporary bool

	fields map[string]any
}

var defaultMeta = ErrorMeta{
	IsImportant: false,
	IsTemporary: false,
	fields:      nil,
}

func (e *ErrorMeta) Fields() map[string]any {
	if e.fields == nil {
		return map[string]any{}
	}

	return e.fields
}

type ExtendedError struct {
	ctx context.Context

	error error

	meta ErrorMeta
}

type ErrorMetaOpt func(*ExtendedError)

func New(ctx context.Context, err error, opts ...ErrorMetaOpt) *ExtendedError {
	error := &ExtendedError{
		ctx:   ctx,
		error: err,
		meta:  defaultMeta,
	}

	for _, opt := range opts {
		opt(error)
	}

	return error
}

func (e *ExtendedError) Error() error {
	return e.error
}

func (e *ExtendedError) Meta() ErrorMeta {
	return e.meta
}

func Important() ErrorMetaOpt {
	return func(e *ExtendedError) { e.meta.IsImportant = true }
}

func Temporary() ErrorMetaOpt {
	return func(e *ExtendedError) { e.meta.IsTemporary = true }
}
