package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestAG_BlobConstructor(t *testing.T) {
	v := dml.NewAG_Blob()
	if v == nil {
		t.Errorf("dml.NewAG_Blob must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.AG_Blob should validate: %s", err)
	}
}

func TestAG_BlobMarshalUnmarshal(t *testing.T) {
	v := dml.NewAG_Blob()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewAG_Blob()
	xml.Unmarshal(buf, v2)
}
