package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_AudioCDConstructor(t *testing.T) {
	v := dml.NewCT_AudioCD()
	if v == nil {
		t.Errorf("dml.NewCT_AudioCD must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_AudioCD should validate: %s", err)
	}
}

func TestCT_AudioCDMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_AudioCD()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_AudioCD()
	xml.Unmarshal(buf, v2)
}
