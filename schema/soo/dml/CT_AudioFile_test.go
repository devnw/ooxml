package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_AudioFileConstructor(t *testing.T) {
	v := dml.NewCT_AudioFile()
	if v == nil {
		t.Errorf("dml.NewCT_AudioFile must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_AudioFile should validate: %s", err)
	}
}

func TestCT_AudioFileMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_AudioFile()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_AudioFile()
	xml.Unmarshal(buf, v2)
}
