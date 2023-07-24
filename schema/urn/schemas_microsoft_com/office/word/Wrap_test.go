package word_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/office/word"
)

func TestWrapConstructor(t *testing.T) {
	v := word.NewWrap()
	if v == nil {
		t.Errorf("word.NewWrap must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed word.Wrap should validate: %s", err)
	}
}

func TestWrapMarshalUnmarshal(t *testing.T) {
	v := word.NewWrap()
	buf, _ := xml.Marshal(v)
	v2 := word.NewWrap()
	xml.Unmarshal(buf, v2)
}
