package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_TileInfoPropertiesConstructor(t *testing.T) {
	v := dml.NewCT_TileInfoProperties()
	if v == nil {
		t.Errorf("dml.NewCT_TileInfoProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_TileInfoProperties should validate: %s", err)
	}
}

func TestCT_TileInfoPropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_TileInfoProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_TileInfoProperties()
	xml.Unmarshal(buf, v2)
}
