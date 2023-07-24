package terms_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/purl.org/dc/terms"
)

func TestPeriodConstructor(t *testing.T) {
	v := terms.NewPeriod()
	if v == nil {
		t.Errorf("terms.NewPeriod must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed terms.Period should validate: %s", err)
	}
}

func TestPeriodMarshalUnmarshal(t *testing.T) {
	v := terms.NewPeriod()
	buf, _ := xml.Marshal(v)
	v2 := terms.NewPeriod()
	xml.Unmarshal(buf, v2)
}
