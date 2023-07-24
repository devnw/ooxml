package common

import (
	"fmt"

	"go.devnw.com/ooxml/schema/soo/pkg/relationships"
)

// Relationship is a relationship within a .rels file.
type Relationship struct {
	x *relationships.Relationship
}

// NewRelationship constructs a new relationship.
func NewRelationship() Relationship {
	return Relationship{relationships.NewRelationship()}
}

// X returns the inner wrapped XML type.
func (r Relationship) X() *relationships.Relationship {
	return r.x
}

// ID returns the ID of a relationship.
func (r Relationship) ID() string {
	return r.x.IdAttr
}

// SetTarget set the target (path) of a relationship.
func (r Relationship) SetTarget(s string) {
	r.x.TargetAttr = s
}

// Target returns the target (path) of a relationship.
func (r Relationship) Target() string {
	return r.x.TargetAttr
}

// Type returns the type of a relationship.
func (r Relationship) Type() string {
	return r.x.TypeAttr
}

func (r Relationship) String() string {
	return fmt.Sprintf("{ID: %s Target: %s Type: %s}", r.ID(), r.Target(), r.Type())
}
