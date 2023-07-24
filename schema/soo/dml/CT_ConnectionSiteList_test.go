package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_ConnectionSiteListConstructor(t *testing.T) {
	v := dml.NewCT_ConnectionSiteList()
	if v == nil {
		t.Errorf("dml.NewCT_ConnectionSiteList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_ConnectionSiteList should validate: %s", err)
	}
}

func TestCT_ConnectionSiteListMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_ConnectionSiteList()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_ConnectionSiteList()
	xml.Unmarshal(buf, v2)
}
