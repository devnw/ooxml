package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestVideoFileConstructor(t *testing.T) {
	v := dml.NewVideoFile()
	if v == nil {
		t.Errorf("dml.NewVideoFile must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.VideoFile should validate: %s", err)
	}
}

func TestVideoFileMarshalUnmarshal(t *testing.T) {
	v := dml.NewVideoFile()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewVideoFile()
	xml.Unmarshal(buf, v2)
}
