package dml

import (
	"encoding/xml"
	"fmt"
)

// ST_AdjAngle is a union type
type ST_AdjAngle struct {
	ST_Angle         *int32
	ST_GeomGuideName *string
}

func (m *ST_AdjAngle) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_AdjAngle) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_Angle != nil {
		e.EncodeToken(xml.CharData(fmt.Sprintf("%d", *m.ST_Angle)))
	}
	if m.ST_GeomGuideName != nil {
		e.EncodeToken(xml.CharData(*m.ST_GeomGuideName))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_AdjAngle) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_Angle != nil {
		mems = append(mems, "ST_Angle")
	}
	if m.ST_GeomGuideName != nil {
		mems = append(mems, "ST_GeomGuideName")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_AdjAngle) String() string {
	if m.ST_Angle != nil {
		return fmt.Sprintf("%v", *m.ST_Angle)
	}
	if m.ST_GeomGuideName != nil {
		return fmt.Sprintf("%v", *m.ST_GeomGuideName)
	}
	return ""
}
