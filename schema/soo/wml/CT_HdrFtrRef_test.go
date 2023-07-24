package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_HdrFtrRefConstructor(t *testing.T) {
	v := wml.NewCT_HdrFtrRef()
	if v == nil {
		t.Errorf("wml.NewCT_HdrFtrRef must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_HdrFtrRef should validate: %s", err)
	}
}

func TestCT_HdrFtrRefMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_HdrFtrRef()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_HdrFtrRef()
	xml.Unmarshal(buf, v2)
}
