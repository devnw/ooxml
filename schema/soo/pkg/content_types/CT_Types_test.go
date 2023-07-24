package content_types_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pkg/content_types"
)

func TestCT_TypesConstructor(t *testing.T) {
	v := content_types.NewCT_Types()
	if v == nil {
		t.Errorf("content_types.NewCT_Types must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed content_types.CT_Types should validate: %s", err)
	}
}

func TestCT_TypesMarshalUnmarshal(t *testing.T) {
	v := content_types.NewCT_Types()
	buf, _ := xml.Marshal(v)
	v2 := content_types.NewCT_Types()
	xml.Unmarshal(buf, v2)
}
