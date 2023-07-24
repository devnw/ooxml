package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_ColumnsConstructor(t *testing.T) {
	v := wml.NewCT_Columns()
	if v == nil {
		t.Errorf("wml.NewCT_Columns must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_Columns should validate: %s", err)
	}
}

func TestCT_ColumnsMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_Columns()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_Columns()
	xml.Unmarshal(buf, v2)
}
