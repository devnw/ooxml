package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_FileRecoveryPrConstructor(t *testing.T) {
	v := sml.NewCT_FileRecoveryPr()
	if v == nil {
		t.Errorf("sml.NewCT_FileRecoveryPr must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_FileRecoveryPr should validate: %s", err)
	}
}

func TestCT_FileRecoveryPrMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_FileRecoveryPr()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_FileRecoveryPr()
	xml.Unmarshal(buf, v2)
}
