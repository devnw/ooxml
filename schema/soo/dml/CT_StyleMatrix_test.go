package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_StyleMatrixConstructor(t *testing.T) {
	v := dml.NewCT_StyleMatrix()
	if v == nil {
		t.Errorf("dml.NewCT_StyleMatrix must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_StyleMatrix should validate: %s", err)
	}
}

func TestCT_StyleMatrixMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_StyleMatrix()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_StyleMatrix()
	xml.Unmarshal(buf, v2)
}
