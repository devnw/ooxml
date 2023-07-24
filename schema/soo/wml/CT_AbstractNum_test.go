package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_AbstractNumConstructor(t *testing.T) {
	v := wml.NewCT_AbstractNum()
	if v == nil {
		t.Errorf("wml.NewCT_AbstractNum must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_AbstractNum should validate: %s", err)
	}
}

func TestCT_AbstractNumMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_AbstractNum()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_AbstractNum()
	xml.Unmarshal(buf, v2)
}
