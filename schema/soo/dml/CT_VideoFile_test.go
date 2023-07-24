package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_VideoFileConstructor(t *testing.T) {
	v := dml.NewCT_VideoFile()
	if v == nil {
		t.Errorf("dml.NewCT_VideoFile must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_VideoFile should validate: %s", err)
	}
}

func TestCT_VideoFileMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_VideoFile()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_VideoFile()
	xml.Unmarshal(buf, v2)
}
