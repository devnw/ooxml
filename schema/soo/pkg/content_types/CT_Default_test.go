package content_types_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pkg/content_types"
)

func TestCT_DefaultConstructor(t *testing.T) {
	v := content_types.NewCT_Default()
	if v == nil {
		t.Errorf("content_types.NewCT_Default must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed content_types.CT_Default should validate: %s", err)
	}
}

func TestCT_DefaultMarshalUnmarshal(t *testing.T) {
	v := content_types.NewCT_Default()
	buf, _ := xml.Marshal(v)
	v2 := content_types.NewCT_Default()
	xml.Unmarshal(buf, v2)
}
