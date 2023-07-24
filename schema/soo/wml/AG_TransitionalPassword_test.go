package wml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/wml"
)

func TestAG_TransitionalPasswordConstructor(t *testing.T) {
	v := wml.NewAG_TransitionalPassword()
	if v == nil {
		t.Errorf("wml.NewAG_TransitionalPassword must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed wml.AG_TransitionalPassword should validate: %s", err)
	}
}

func TestAG_TransitionalPasswordMarshalUnmarshal(t *testing.T) {
	v := wml.NewAG_TransitionalPassword()
	buf, _ := xml.Marshal(v)
	v2 := wml.NewAG_TransitionalPassword()
	xml.Unmarshal(buf, v2)
}
