package word_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/office/word"
)

func TestBorderrightConstructor(t *testing.T) {
	v := word.NewBorderright()
	if v == nil {
		t.Errorf("word.NewBorderright must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed word.Borderright should validate: %s", err)
	}
}

func TestBorderrightMarshalUnmarshal(t *testing.T) {
	v := word.NewBorderright()
	buf, _ := xml.Marshal(v)
	v2 := word.NewBorderright()
	xml.Unmarshal(buf, v2)
}
