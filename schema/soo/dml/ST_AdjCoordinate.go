package dml

import (
	"encoding/xml"
	"fmt"
)

// ST_AdjCoordinate is a union type
type ST_AdjCoordinate struct {
	ST_Coordinate    *ST_Coordinate
	ST_GeomGuideName *string
}

func (m *ST_AdjCoordinate) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_AdjCoordinate) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_Coordinate != nil {
		e.Encode(m.ST_Coordinate)
	}
	if m.ST_GeomGuideName != nil {
		e.EncodeToken(xml.CharData(*m.ST_GeomGuideName))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_AdjCoordinate) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_Coordinate != nil {
		if err := m.ST_Coordinate.ValidateWithPath(path + "/ST_Coordinate"); err != nil {
			return err
		}
		mems = append(mems, "ST_Coordinate")
	}
	if m.ST_GeomGuideName != nil {
		mems = append(mems, "ST_GeomGuideName")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_AdjCoordinate) String() string {
	if m.ST_Coordinate != nil {
		return m.ST_Coordinate.String()
	}
	if m.ST_GeomGuideName != nil {
		return fmt.Sprintf("%v", *m.ST_GeomGuideName)
	}
	return ""
}
