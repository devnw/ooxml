package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_QueryTableRefreshConstructor(t *testing.T) {
	v := sml.NewCT_QueryTableRefresh()
	if v == nil {
		t.Errorf("sml.NewCT_QueryTableRefresh must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_QueryTableRefresh should validate: %s", err)
	}
}

func TestCT_QueryTableRefreshMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_QueryTableRefresh()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_QueryTableRefresh()
	xml.Unmarshal(buf, v2)
}
