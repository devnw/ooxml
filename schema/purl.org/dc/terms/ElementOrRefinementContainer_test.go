package terms_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/purl.org/dc/terms"
)

func TestElementOrRefinementContainerConstructor(t *testing.T) {
	v := terms.NewElementOrRefinementContainer()
	if v == nil {
		t.Errorf("terms.NewElementOrRefinementContainer must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.ElementOrRefinementContainer should validate: %s", err)
	}
}

func TestElementOrRefinementContainerMarshalUnmarshal(t *testing.T) {
	v := terms.NewElementOrRefinementContainer()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewElementOrRefinementContainer()
	xml.Unmarshal(buf, v2)
}
