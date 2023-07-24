package content_types_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pkg/content_types"
)

func TestDefaultConstructor(t *testing.T) {
	v := content_types.NewDefault()
	if v == nil {
		t.Errorf("content_types.NewDefault must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed content_types.Default should validate: %s", err)
	}
}

func TestDefaultMarshalUnmarshal(t *testing.T) {
	v := content_types.NewDefault()
	buf, _ := xml.Marshal(v)
	v2 := content_types.NewDefault()
	xml.Unmarshal(buf, v2)
}
