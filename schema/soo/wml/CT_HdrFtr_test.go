package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_HdrFtrConstructor(t *testing.T) {
	v := wml.NewCT_HdrFtr()
	if v == nil {
		t.Errorf("wml.NewCT_HdrFtr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_HdrFtr should validate: %s", err)
	}
}

func TestCT_HdrFtrMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_HdrFtr()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_HdrFtr()
	xml.Unmarshal(buf, v2)
}
