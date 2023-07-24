package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestRelIdsConstructor(t *testing.T) {
	v := diagram.NewRelIds()
	if v == nil {
		t.Errorf("diagram.NewRelIds must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.RelIds should validate: %s", err)
	}
}

func TestRelIdsMarshalUnmarshal(t *testing.T) {
	v := diagram.NewRelIds()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewRelIds()
	xml.Unmarshal(buf, v2)
}
