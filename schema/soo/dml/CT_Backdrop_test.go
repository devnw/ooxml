package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_BackdropConstructor(t *testing.T) {
	v := dml.NewCT_Backdrop()
	if v == nil {
		t.Errorf("dml.NewCT_Backdrop must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_Backdrop should validate: %s", err)
	}
}

func TestCT_BackdropMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_Backdrop()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_Backdrop()
	xml.Unmarshal(buf, v2)
}
