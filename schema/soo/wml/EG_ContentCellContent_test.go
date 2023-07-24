package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestEG_ContentCellContentConstructor(t *testing.T) {
	v := wml.NewEG_ContentCellContent()
	if v == nil {
		t.Errorf("wml.NewEG_ContentCellContent must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.EG_ContentCellContent should validate: %s", err)
	}
}

func TestEG_ContentCellContentMarshalUnmarshal(t *testing.T) {
	v := wml.NewEG_ContentCellContent()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewEG_ContentCellContent()
	xml.Unmarshal(buf, v2)
}
