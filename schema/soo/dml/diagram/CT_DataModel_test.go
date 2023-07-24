package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_DataModelConstructor(t *testing.T) {
	v := diagram.NewCT_DataModel()
	if v == nil {
		t.Errorf("diagram.NewCT_DataModel must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_DataModel should validate: %s", err)
	}
}

func TestCT_DataModelMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_DataModel()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_DataModel()
	xml.Unmarshal(buf, v2)
}
