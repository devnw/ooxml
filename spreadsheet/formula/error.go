package formula

import "go.devnw.com/ooxml/spreadsheet/update"

// Error is an error expression.
type Error struct {
	s string
}

// NewError constructs a new error expression from a string.
func NewError(v string) Expression {
	return Error{v}
}

// Eval evaluates and returns the result of an error expression.
func (e Error) Eval(ctx Context, ev Evaluator) Result {
	return MakeErrorResult(e.s)
}

// Reference returns an invalid reference for Error.
func (e Error) Reference(ctx Context, ev Evaluator) Reference {
	return ReferenceInvalid
}

// String returns an empty string for Error.
func (e Error) String() string {
	return ""
}

// Update returns the same object as updating sheet references does not affect Error.
func (e Error) Update(q *update.UpdateQuery) Expression {
	return e
}
