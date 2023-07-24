package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_CnfConstructor(t *testing.T) {
	v := wml.NewCT_Cnf()
	if v == nil {
		t.Errorf("wml.NewCT_Cnf must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Cnf should validate: %s", err)
	}
}

func TestCT_CnfMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Cnf()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Cnf()
	xml.Unmarshal(buf, v2)
}
