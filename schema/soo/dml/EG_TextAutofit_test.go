package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestEG_TextAutofitConstructor(t *testing.T) {
	v := dml.NewEG_TextAutofit()
	if v == nil {
		t.Errorf("dml.NewEG_TextAutofit must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_TextAutofit should validate: %s", err)
	}
}

func TestEG_TextAutofitMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_TextAutofit()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_TextAutofit()
	xml.Unmarshal(buf, v2)
}
