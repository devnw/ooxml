package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_DirContentRunConstructor(t *testing.T) {
	v := wml.NewCT_DirContentRun()
	if v == nil {
		t.Errorf("wml.NewCT_DirContentRun must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_DirContentRun should validate: %s", err)
	}
}

func TestCT_DirContentRunMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_DirContentRun()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_DirContentRun()
	xml.Unmarshal(buf, v2)
}
