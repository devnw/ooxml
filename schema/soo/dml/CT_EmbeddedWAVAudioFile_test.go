package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_EmbeddedWAVAudioFileConstructor(t *testing.T) {
	v := dml.NewCT_EmbeddedWAVAudioFile()
	if v == nil {
		t.Errorf("dml.NewCT_EmbeddedWAVAudioFile must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_EmbeddedWAVAudioFile should validate: %s", err)
	}
}

func TestCT_EmbeddedWAVAudioFileMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_EmbeddedWAVAudioFile()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_EmbeddedWAVAudioFile()
	xml.Unmarshal(buf, v2)
}
