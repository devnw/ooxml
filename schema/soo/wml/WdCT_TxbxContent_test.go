package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestWdCT_TxbxContentConstructor(t *testing.T) {
	v := wml.NewWdCT_TxbxContent()
	if v == nil {
		t.Errorf("wml.NewWdCT_TxbxContent must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.WdCT_TxbxContent should validate: %s", err)
	}
}

func TestWdCT_TxbxContentMarshalUnmarshal(t *testing.T) {
	v := wml.NewWdCT_TxbxContent()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewWdCT_TxbxContent()
	xml.Unmarshal(buf, v2)
}
