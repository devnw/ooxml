package lockedCanvas_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/lockedCanvas"
)

func TestLockedCanvasConstructor(t *testing.T) {
	v := lockedCanvas.NewLockedCanvas()
	if v == nil {
		t.Errorf("lockedCanvas.NewLockedCanvas must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed lockedCanvas.LockedCanvas should validate: %s", err)
	}
}

func TestLockedCanvasMarshalUnmarshal(t *testing.T) {
	v := lockedCanvas.NewLockedCanvas()
	buf, _ := xml.Marshal(v)
	v2 := lockedCanvas.NewLockedCanvas()
	xml.Unmarshal(buf, v2)
}
