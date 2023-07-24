package content_types_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pkg/content_types"
)

func TestOverrideConstructor(t *testing.T) {
	v := content_types.NewOverride()
	if v == nil {
		t.Errorf("content_types.NewOverride must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed content_types.Override should validate: %s", err)
	}
}

func TestOverrideMarshalUnmarshal(t *testing.T) {
	v := content_types.NewOverride()
	buf, _ := xml.Marshal(v)
	v2 := content_types.NewOverride()
	xml.Unmarshal(buf, v2)
}
