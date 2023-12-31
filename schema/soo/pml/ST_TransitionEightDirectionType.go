package pml

import (
	"encoding/xml"
	"fmt"
)

// ST_TransitionEightDirectionType is a union type
type ST_TransitionEightDirectionType struct {
	ST_TransitionSideDirectionType   ST_TransitionSideDirectionType
	ST_TransitionCornerDirectionType ST_TransitionCornerDirectionType
}

func (m *ST_TransitionEightDirectionType) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_TransitionEightDirectionType) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_TransitionSideDirectionType != ST_TransitionSideDirectionTypeUnset {
		e.EncodeToken(xml.CharData(m.ST_TransitionSideDirectionType.String()))
	}
	if m.ST_TransitionCornerDirectionType != ST_TransitionCornerDirectionTypeUnset {
		e.EncodeToken(xml.CharData(m.ST_TransitionCornerDirectionType.String()))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_TransitionEightDirectionType) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_TransitionSideDirectionType != ST_TransitionSideDirectionTypeUnset {
		mems = append(mems, "ST_TransitionSideDirectionType")
	}
	if m.ST_TransitionCornerDirectionType != ST_TransitionCornerDirectionTypeUnset {
		mems = append(mems, "ST_TransitionCornerDirectionType")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_TransitionEightDirectionType) String() string {
	if m.ST_TransitionSideDirectionType != ST_TransitionSideDirectionTypeUnset {
		return m.ST_TransitionSideDirectionType.String()
	}
	if m.ST_TransitionCornerDirectionType != ST_TransitionCornerDirectionTypeUnset {
		return m.ST_TransitionCornerDirectionType.String()
	}
	return ""
}
