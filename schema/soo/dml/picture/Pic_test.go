package picture_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/picture"
)

func TestPicConstructor(t *testing.T) {
	v := picture.NewPic()
	if v == nil {
		t.Errorf("picture.NewPic must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed picture.Pic should validate: %s", err)
	}
}

func TestPicMarshalUnmarshal(t *testing.T) {
	v := picture.NewPic()
	buf, _ := xml.Marshal(v)
	v2 := picture.NewPic()
	xml.Unmarshal(buf, v2)
}
