package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_SdtListItemConstructor(t *testing.T) {
	v := wml.NewCT_SdtListItem()
	if v == nil {
		t.Errorf("wml.NewCT_SdtListItem must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_SdtListItem should validate: %s", err)
	}
}

func TestCT_SdtListItemMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_SdtListItem()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_SdtListItem()
	xml.Unmarshal(buf, v2)
}
