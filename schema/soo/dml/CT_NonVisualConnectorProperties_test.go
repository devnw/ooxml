package dml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml"
)

func TestCT_NonVisualConnectorPropertiesConstructor(t *testing.T) {
	v := dml.NewCT_NonVisualConnectorProperties()
	if v == nil {
		t.Errorf("dml.NewCT_NonVisualConnectorProperties must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed dml.CT_NonVisualConnectorProperties should validate: %s", err)
	}
}

func TestCT_NonVisualConnectorPropertiesMarshalUnmarshal(t *testing.T) {
	v := dml.NewCT_NonVisualConnectorProperties()
	buf, _ := xml.Marshal(v)
	v2 := dml.NewCT_NonVisualConnectorProperties()
	xml.Unmarshal(buf, v2)
}
