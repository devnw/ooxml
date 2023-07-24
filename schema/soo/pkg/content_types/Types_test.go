package content_types_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pkg/content_types"
)

func TestTypesConstructor(t *testing.T) {
	v := content_types.NewTypes()
	if v == nil {
		t.Errorf("content_types.NewTypes must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed content_types.Types should validate: %s", err)
	}
}

func TestTypesMarshalUnmarshal(t *testing.T) {
	v := content_types.NewTypes()
	buf, _ := xml.Marshal(v)
	v2 := content_types.NewTypes()
	xml.Unmarshal(buf, v2)
}
