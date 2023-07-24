package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestPivotCacheRecordsConstructor(t *testing.T) {
	v := sml.NewPivotCacheRecords()
	if v == nil {
		t.Errorf("sml.NewPivotCacheRecords must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.PivotCacheRecords should validate: %s", err)
	}
}

func TestPivotCacheRecordsMarshalUnmarshal(t *testing.T) {
	v := sml.NewPivotCacheRecords()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewPivotCacheRecords()
	xml.Unmarshal(buf, v2)
}
