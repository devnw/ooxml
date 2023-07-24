package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TextNoAutofitConstructor(t *testing.T) {
	v := dml.NewCT_TextNoAutofit()
	if v == nil {
		t.Errorf("dml.NewCT_TextNoAutofit must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TextNoAutofit should validate: %s", err)
	}
}

func TestCT_TextNoAutofitMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TextNoAutofit()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TextNoAutofit()
	xml.Unmarshal(buf, v2)
}
