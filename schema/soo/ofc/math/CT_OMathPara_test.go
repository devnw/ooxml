package math_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/math"
)

func TestCT_OMathParaConstructor(t *testing.T) {
	v := math.NewCT_OMathPara()
	if v == nil {
		t.Errorf("math.NewCT_OMathPara must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed math.CT_OMathPara should validate: %s", err)
	}
}

func TestCT_OMathParaMarshalUnmarshal(t *testing.T) {
	v := math.NewCT_OMathPara()
	buf, _ := xml.Marshal(v)
	v2 := math.NewCT_OMathPara()
	xml.Unmarshal(buf, v2)
}
