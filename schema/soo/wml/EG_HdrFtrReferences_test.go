package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestEG_HdrFtrReferencesConstructor(t *testing.T) {
	v := wml.NewEG_HdrFtrReferences()
	if v == nil {
		t.Errorf("wml.NewEG_HdrFtrReferences must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.EG_HdrFtrReferences should validate: %s", err)
	}
}

func TestEG_HdrFtrReferencesMarshalUnmarshal(t *testing.T) {
	v := wml.NewEG_HdrFtrReferences()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewEG_HdrFtrReferences()
	xml.Unmarshal(buf, v2)
}
