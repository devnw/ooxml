package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestEG_SectPrContentsConstructor(t *testing.T) {
	v := wml.NewEG_SectPrContents()
	if v == nil {
		t.Errorf("wml.NewEG_SectPrContents must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.EG_SectPrContents should validate: %s", err)
	}
}

func TestEG_SectPrContentsMarshalUnmarshal(t *testing.T) {
	v := wml.NewEG_SectPrContents()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewEG_SectPrContents()
	xml.Unmarshal(buf, v2)
}
