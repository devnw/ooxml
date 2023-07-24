package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_SDDescriptionConstructor(t *testing.T) {
	v := diagram.NewCT_SDDescription()
	if v == nil {
		t.Errorf("diagram.NewCT_SDDescription must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_SDDescription should validate: %s", err)
	}
}

func TestCT_SDDescriptionMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_SDDescription()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_SDDescription()
	xml.Unmarshal(buf, v2)
}
