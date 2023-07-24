package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestCT_PivotCacheRecordsConstructor(t *testing.T) {
	v := sml.NewCT_PivotCacheRecords()
	if v == nil {
		t.Errorf("sml.NewCT_PivotCacheRecords must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.CT_PivotCacheRecords should validate: %s", err)
	}
}

func TestCT_PivotCacheRecordsMarshalUnmarshal(t *testing.T) {
	v := sml.NewCT_PivotCacheRecords()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewCT_PivotCacheRecords()
	xml.Unmarshal(buf, v2)
}
