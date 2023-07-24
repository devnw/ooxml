package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_FFDataConstructor(t *testing.T) {
	v := wml.NewCT_FFData()
	if v == nil {
		t.Errorf("wml.NewCT_FFData must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_FFData should validate: %s", err)
	}
}

func TestCT_FFDataMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_FFData()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_FFData()
	xml.Unmarshal(buf, v2)
}
