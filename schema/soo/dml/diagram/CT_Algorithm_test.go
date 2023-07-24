package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_AlgorithmConstructor(t *testing.T) {
	v := diagram.NewCT_Algorithm()
	if v == nil {
		t.Errorf("diagram.NewCT_Algorithm must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_Algorithm should validate: %s", err)
	}
}

func TestCT_AlgorithmMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_Algorithm()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_Algorithm()
	xml.Unmarshal(buf, v2)
}
