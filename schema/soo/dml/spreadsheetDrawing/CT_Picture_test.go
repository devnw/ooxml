package spreadsheetDrawing_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/dml/spreadsheetDrawing"
)

func TestCT_PictureConstructor(t *testing.T) {
	v := spreadsheetDrawing.NewCT_Picture()
	if v == nil {
		t.Errorf("spreadsheetDrawing.NewCT_Picture must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed spreadsheetDrawing.CT_Picture should validate: %s", err)
	}
}

func TestCT_PictureMarshalUnmarshal(t *testing.T) {
	v := spreadsheetDrawing.NewCT_Picture()
	buf, _ := xml.Marshal(v)
	v2 := spreadsheetDrawing.NewCT_Picture()
	xml.Unmarshal(buf, v2)
}
