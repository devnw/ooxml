package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_SampleDataConstructor(t *testing.T) {
	v := diagram.NewCT_SampleData()
	if v == nil {
		t.Errorf("diagram.NewCT_SampleData must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_SampleData should validate: %s", err)
	}
}

func TestCT_SampleDataMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_SampleData()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_SampleData()
	xml.Unmarshal(buf, v2)
}
