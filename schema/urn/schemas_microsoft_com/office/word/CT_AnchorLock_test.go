package word_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/office/word"
)

func TestCT_AnchorLockConstructor(t *testing.T) {
	v := word.NewCT_AnchorLock()
	if v == nil {
		t.Errorf("word.NewCT_AnchorLock must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed word.CT_AnchorLock should validate: %s", err)
	}
}

func TestCT_AnchorLockMarshalUnmarshal(t *testing.T) {
	v := word.NewCT_AnchorLock()
	buf, _ := xml.Marshal(v)
	v2 := word.NewCT_AnchorLock()
	xml.Unmarshal(buf, v2)
}
