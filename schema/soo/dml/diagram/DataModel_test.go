package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestDataModelConstructor(t *testing.T) {
	v := diagram.NewDataModel()
	if v == nil {
		t.Errorf("diagram.NewDataModel must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.DataModel should validate: %s", err)
	}
}

func TestDataModelMarshalUnmarshal(t *testing.T) {
	v := diagram.NewDataModel()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewDataModel()
	xml.Unmarshal(buf, v2)
}
