package excel_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/urn/schemas_microsoft_com/office/excel"
)

func TestCT_ClientDataConstructor(t *testing.T) {
	v := excel.NewCT_ClientData()
	if v == nil {
		t.Errorf("excel.NewCT_ClientData must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed excel.CT_ClientData should validate: %s", err)
	}
}

func TestCT_ClientDataMarshalUnmarshal(t *testing.T) {
	v := excel.NewCT_ClientData()
	buf, _ := xml.Marshal(v)
	v2 := excel.NewCT_ClientData()
	xml.Unmarshal(buf, v2)
}
