package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_GvmlConnectorConstructor(t *testing.T) {
	v := dml.NewCT_GvmlConnector()
	if v == nil {
		t.Errorf("dml.NewCT_GvmlConnector must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_GvmlConnector should validate: %s", err)
	}
}

func TestCT_GvmlConnectorMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_GvmlConnector()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_GvmlConnector()
	xml.Unmarshal(buf, v2)
}
