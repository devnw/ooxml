package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_ColumnConstructor(t *testing.T) {
	v := wml.NewCT_Column()
	if v == nil {
		t.Errorf("wml.NewCT_Column must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Column should validate: %s", err)
	}
}

func TestCT_ColumnMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Column()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Column()
	xml.Unmarshal(buf, v2)
}
