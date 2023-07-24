package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestCT_OdsoFieldMapDataConstructor(t *testing.T) {
	v := wml.NewCT_OdsoFieldMapData()
	if v == nil {
		t.Errorf("wml.NewCT_OdsoFieldMapData must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.CT_OdsoFieldMapData should validate: %s", err)
	}
}

func TestCT_OdsoFieldMapDataMarshalUnmarshal(t *testing.T) {
	v := wml.NewCT_OdsoFieldMapData()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewCT_OdsoFieldMapData()
	xml.Unmarshal(buf, v2)
}
