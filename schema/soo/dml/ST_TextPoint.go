package dml

import (
	"encoding/xml"
	"fmt"
)

// ST_TextPoint is a union type
type ST_TextPoint struct {
	ST_TextPointUnqualified *int32
	ST_UniversalMeasure     *string
}

func (m *ST_TextPoint) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TextPoint) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_TextPointUnqualified != nil {
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *m.ST_TextPointUnqualified)))
	}
	if m.ST_UniversalMeasure != nil {
		e.EncodeToken(xml.CharData(*m.ST_UniversalMeasure))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_TextPoint) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_TextPointUnqualified != nil {
		mems = append(mems, "ST_TextPointUnqualified")
	}
	if m.ST_UniversalMeasure != nil {
		mems = append(mems, "ST_UniversalMeasure")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_TextPoint) String() string {
	if m.ST_TextPointUnqualified != nil {
		return fmt.Sprintf("%v", *m.ST_TextPointUnqualified)
	}
	if m.ST_UniversalMeasure != nil {
		return fmt.Sprintf("%v", *m.ST_UniversalMeasure)
	}
	return ""
}
