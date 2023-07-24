package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_FontCollectionConstructor(t *testing.T) {
	v := dml.NewCT_FontCollection()
	if v == nil {
		t.Errorf("dml.NewCT_FontCollection must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_FontCollection should validate: %s", err)
	}
}

func TestCT_FontCollectionMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_FontCollection()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_FontCollection()
	xml.Unmarshal(buf, v2)
}
