package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_BackgroundFillStyleListConstructor(t *testing.T) {
	v := dml.NewCT_BackgroundFillStyleList()
	if v == nil {
		t.Errorf("dml.NewCT_BackgroundFillStyleList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_BackgroundFillStyleList should validate: %s", err)
	}
}

func TestCT_BackgroundFillStyleListMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_BackgroundFillStyleList()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_BackgroundFillStyleList()
	xml.Unmarshal(buf, v2)
}
