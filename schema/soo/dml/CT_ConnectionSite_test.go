package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_ConnectionSiteConstructor(t *testing.T) {
	v := dml.NewCT_ConnectionSite()
	if v == nil {
		t.Errorf("dml.NewCT_ConnectionSite must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_ConnectionSite should validate: %s", err)
	}
}

func TestCT_ConnectionSiteMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_ConnectionSite()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_ConnectionSite()
	xml.Unmarshal(buf, v2)
}
