package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_JcTableConstructor(t *testing.T) {
	v := wml.NewCT_JcTable()
	if v == nil {
		t.Errorf("wml.NewCT_JcTable must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_JcTable should validate: %s", err)
	}
}

func TestCT_JcTableMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_JcTable()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_JcTable()
	xml.Unmarshal(buf, v2)
}
