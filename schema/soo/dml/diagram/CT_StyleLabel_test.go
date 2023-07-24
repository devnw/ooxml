package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_StyleLabelConstructor(t *testing.T) {
	v := diagram.NewCT_StyleLabel()
	if v == nil {
		t.Errorf("diagram.NewCT_StyleLabel must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_StyleLabel should validate: %s", err)
	}
}

func TestCT_StyleLabelMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_StyleLabel()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_StyleLabel()
	xml.Unmarshal(buf, v2)
}
