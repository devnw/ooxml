package excel_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/office/excel"
)

func TestClientDataConstructor(t *testing.T) {
	v := excel.NewClientData()
	if v == nil {
		t.Errorf("excel.NewClientData must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed excel.ClientData should validate: %s", err)
	}
}

func TestClientDataMarshalUnmarshal(t *testing.T) {
	v := excel.NewClientData()
	buf, _ := xml.Marshal(v)
	v2 := excel.NewClientData()
	xml.Unmarshal(buf, v2)
}
