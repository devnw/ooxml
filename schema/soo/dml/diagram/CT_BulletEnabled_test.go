package diagram_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/diagram"
)

func TestCT_BulletEnabledConstructor(t *testing.T) {
	v := diagram.NewCT_BulletEnabled()
	if v == nil {
		t.Errorf("diagram.NewCT_BulletEnabled must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed diagram.CT_BulletEnabled should validate: %s", err)
	}
}

func TestCT_BulletEnabledMarshalUnmarshal(t *testing.T) {
	v := diagram.NewCT_BulletEnabled()
	buf, _ := xml.Marshal(v)
	v2 := diagram.NewCT_BulletEnabled()
	xml.Unmarshal(buf, v2)
}
