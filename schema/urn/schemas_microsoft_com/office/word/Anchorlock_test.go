package word_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/office/word"
)

func TestAnchorlockConstructor(t *testing.T) {
	v := word.NewAnchorlock()
	if v == nil {
		t.Errorf("word.NewAnchorlock must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed word.Anchorlock should validate: %s", err)
	}
}

func TestAnchorlockMarshalUnmarshal(t *testing.T) {
	v := word.NewAnchorlock()
	buf, _ := xml.Marshal(v)
	v2 := word.NewAnchorlock()
	xml.Unmarshal(buf, v2)
}
