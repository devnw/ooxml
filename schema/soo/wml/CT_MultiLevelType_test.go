package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_MultiLevelTypeConstructor(t *testing.T) {
	v := wml.NewCT_MultiLevelType()
	if v == nil {
		t.Errorf("wml.NewCT_MultiLevelType must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_MultiLevelType should validate: %s", err)
	}
}

func TestCT_MultiLevelTypeMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_MultiLevelType()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_MultiLevelType()
	xml.Unmarshal(buf, v2)
}
