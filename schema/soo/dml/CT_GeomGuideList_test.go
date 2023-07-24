package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_GeomGuideListConstructor(t *testing.T) {
	v := dml.NewCT_GeomGuideList()
	if v == nil {
		t.Errorf("dml.NewCT_GeomGuideList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GeomGuideList should validate: %s", err)
	}
}

func TestCT_GeomGuideListMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GeomGuideList()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GeomGuideList()
	xml.Unmarshal(buf, v2)
}
