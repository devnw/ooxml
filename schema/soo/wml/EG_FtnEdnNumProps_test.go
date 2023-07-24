package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestEG_FtnEdnNumPropsConstructor(t *testing.T) {
	v := wml.NewEG_FtnEdnNumProps()
	if v == nil {
		t.Errorf("wml.NewEG_FtnEdnNumProps must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.EG_FtnEdnNumProps should validate: %s", err)
	}
}

func TestEG_FtnEdnNumPropsMarshalUnmarshal(t *testing.T) {
	v := wml.NewEG_FtnEdnNumProps()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewEG_FtnEdnNumProps()
	xml.Unmarshal(buf, v2)
}
