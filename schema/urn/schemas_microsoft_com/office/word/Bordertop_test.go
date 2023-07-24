package word_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/office/word"
)

func TestBordertopConstructor(t *testing.T) {
	v := word.NewBordertop()
	if v == nil {
		t.Errorf("word.NewBordertop must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed word.Bordertop should validate: %s", err)
	}
}

func TestBordertopMarshalUnmarshal(t *testing.T) {
	v := word.NewBordertop()
	buf, _ := xml.Marshal(v)
	v2 := word.NewBordertop()
	xml.Unmarshal(buf, v2)
}
