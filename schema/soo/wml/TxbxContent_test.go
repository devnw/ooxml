package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestTxbxContentConstructor(t *testing.T) {
	v := wml.NewTxbxContent()
	if v == nil {
		t.Errorf("wml.NewTxbxContent must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.TxbxContent should validate: %s", err)
	}
}

func TestTxbxContentMarshalUnmarshal(t *testing.T) {
	v := wml.NewTxbxContent()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewTxbxContent()
	xml.Unmarshal(buf, v2)
}
