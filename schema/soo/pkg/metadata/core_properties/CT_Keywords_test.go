package core_properties_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/pkg/metadata/core_properties"
)

func TestCT_KeywordsConstructor(t *testing.T) {
	v := core_properties.NewCT_Keywords()
	if v == nil {
		t.Errorf("core_properties.NewCT_Keywords must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed core_properties.CT_Keywords should validate: %s", err)
	}
}

func TestCT_KeywordsMarshalUnmarshal(t *testing.T) {
	v := core_properties.NewCT_Keywords()
	buf, _ := xml.Marshal(v)
	v2 := core_properties.NewCT_Keywords()
	xml.Unmarshal(buf, v2)
}
