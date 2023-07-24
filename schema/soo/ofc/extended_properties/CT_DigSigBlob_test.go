package extended_properties_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/ofc/extended_properties"
)

func TestCT_DigSigBlobConstructor(t *testing.T) {
	v := extended_properties.NewCT_DigSigBlob()
	if v == nil {
		t.Errorf("extended_properties.NewCT_DigSigBlob must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed extended_properties.CT_DigSigBlob should validate: %s", err)
	}
}

func TestCT_DigSigBlobMarshalUnmarshal(t *testing.T) {
	v := extended_properties.NewCT_DigSigBlob()
	buf, _ := xml.Marshal(v)
	v2 := extended_properties.NewCT_DigSigBlob()
	xml.Unmarshal(buf, v2)
}
