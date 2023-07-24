package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestOMathParaConstructor(t *testing.T) {
	v := math.NewOMathPara()
	if v == nil {
		t.Errorf("math.NewOMathPara must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.OMathPara should validate: %s", err)
	}
}

func TestOMathParaMarshalUnmarshal(t *testing.T) {
	v := math.NewOMathPara()
	buf, _ := xml.Marshal(v)
	v2 := math.NewOMathPara()
	xml.Unmarshal(buf, v2)
}
