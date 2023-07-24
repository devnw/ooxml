package word_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/office/word"
)

func TestBorderbottomConstructor(t *testing.T) {
	v := word.NewBorderbottom()
	if v == nil {
		t.Errorf("word.NewBorderbottom must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed word.Borderbottom should validate: %s", err)
	}
}

func TestBorderbottomMarshalUnmarshal(t *testing.T) {
	v := word.NewBorderbottom()
	buf, _ := xml.Marshal(v)
	v2 := word.NewBorderbottom()
	xml.Unmarshal(buf, v2)
}
