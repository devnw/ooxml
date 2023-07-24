package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_ResizeHandlesConstructor(t *testing.T) {
	v := diagram.NewCT_ResizeHandles()
	if v == nil {
		t.Errorf("diagram.NewCT_ResizeHandles must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_ResizeHandles should validate: %s", err)
	}
}

func TestCT_ResizeHandlesMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_ResizeHandles()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_ResizeHandles()
	xml.Unmarshal(buf, v2)
}
