package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_StyleMatrixReferenceConstructor(t *testing.T) {
	v := dml.NewCT_StyleMatrixReference()
	if v == nil {
		t.Errorf("dml.NewCT_StyleMatrixReference must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_StyleMatrixReference should validate: %s", err)
	}
}

func TestCT_StyleMatrixReferenceMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_StyleMatrixReference()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_StyleMatrixReference()
	xml.Unmarshal(buf, v2)
}
