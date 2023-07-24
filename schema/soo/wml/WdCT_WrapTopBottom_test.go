package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestWdCT_WrapTopBottomConstructor(t *testing.T) {
	v := wml.NewWdCT_WrapTopBottom()
	if v == nil {
		t.Errorf("wml.NewWdCT_WrapTopBottom must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdCT_WrapTopBottom should validate: %s", err)
	}
}

func TestWdCT_WrapTopBottomMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdCT_WrapTopBottom()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdCT_WrapTopBottom()
	xml.Unmarshal(buf, v2)
}
