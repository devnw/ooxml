package word_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/office/word"
)

func TestCT_WrapConstructor(t *testing.T) {
	v := word.NewCT_Wrap()
	if v == nil {
		t.Errorf("word.NewCT_Wrap must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed word.CT_Wrap should validate: %s", err)
	}
}

func TestCT_WrapMarshalUnmarshal(t *testing.T) {
	v := word.NewCT_Wrap()
	buf, _ := xml.Marshal(v)
	v2 := word.NewCT_Wrap()
	xml.Unmarshal(buf, v2)
}
