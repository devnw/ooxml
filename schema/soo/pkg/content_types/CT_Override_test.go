package content_types_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pkg/content_types"
)

func TestCT_OverrideConstructor(t *testing.T) {
	v := content_types.NewCT_Override()
	if v == nil {
		t.Errorf("content_types.NewCT_Override must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed content_types.CT_Override should validate: %s", err)
	}
}

func TestCT_OverrideMarshalUnmarshal(t *testing.T) {
	v := content_types.NewCT_Override()
	buf, _ := xml.Marshal(v)
	v2 := content_types.NewCT_Override()
	xml.Unmarshal(buf, v2)
}
