package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestEG_LineJoinPropertiesConstructor(t *testing.T) {
	v := dml.NewEG_LineJoinProperties()
	if v == nil {
		t.Errorf("dml.NewEG_LineJoinProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.EG_LineJoinProperties should validate: %s", err)
	}
}

func TestEG_LineJoinPropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewEG_LineJoinProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewEG_LineJoinProperties()
	xml.Unmarshal(buf, v2)
}
