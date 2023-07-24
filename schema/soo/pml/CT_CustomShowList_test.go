package pml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pml"
)

func TestCT_CustomShowListConstructor(t *testing.T) {
	v := pml.NewCT_CustomShowList()
	if v == nil {
		t.Errorf("pml.NewCT_CustomShowList must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed pml.CT_CustomShowList should validate: %s", err)
	}
}

func TestCT_CustomShowListMarshalUnmarshal(t *testing.T) {
	v := pml.NewCT_CustomShowList()
	buf, _ := xml.Marshal(v)
	v2 := pml.NewCT_CustomShowList()
	xml.Unmarshal(buf, v2)
}
