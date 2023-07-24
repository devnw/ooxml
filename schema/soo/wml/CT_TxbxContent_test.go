package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_TxbxContentConstructor(t *testing.T) {
	v := wml.NewCT_TxbxContent()
	if v == nil {
		t.Errorf("wml.NewCT_TxbxContent must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_TxbxContent should validate: %s", err)
	}
}

func TestCT_TxbxContentMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_TxbxContent()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_TxbxContent()
	xml.Unmarshal(buf, v2)
}
