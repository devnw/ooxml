package word_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/office/word"
)

func TestBorderleftConstructor(t *testing.T) {
	v := word.NewBorderleft()
	if v == nil {
		t.Errorf("word.NewBorderleft must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed word.Borderleft should validate: %s", err)
	}
}

func TestBorderleftMarshalUnmarshal(t *testing.T) {
	v := word.NewBorderleft()
	buf, _ := xml.Marshal(v)
	v2 := word.NewBorderleft()
	xml.Unmarshal(buf, v2)
}
