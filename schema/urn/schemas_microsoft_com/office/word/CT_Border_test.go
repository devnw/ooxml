package word_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/office/word"
)

func TestCT_BorderConstructor(t *testing.T) {
	v := word.NewCT_Border()
	if v == nil {
		t.Errorf("word.NewCT_Border must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed word.CT_Border should validate: %s", err)
	}
}

func TestCT_BorderMarshalUnmarshal(t *testing.T) {
	v := word.NewCT_Border()
	buf, _ := xml.Marshal(v)
	v2 := word.NewCT_Border()
	xml.Unmarshal(buf, v2)
}
