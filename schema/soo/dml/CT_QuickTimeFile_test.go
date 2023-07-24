package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_QuickTimeFileConstructor(t *testing.T) {
	v := dml.NewCT_QuickTimeFile()
	if v == nil {
		t.Errorf("dml.NewCT_QuickTimeFile must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_QuickTimeFile should validate: %s", err)
	}
}

func TestCT_QuickTimeFileMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_QuickTimeFile()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_QuickTimeFile()
	xml.Unmarshal(buf, v2)
}
