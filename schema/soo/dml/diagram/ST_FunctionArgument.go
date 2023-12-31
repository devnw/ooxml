package diagram

import (
	"encoding/xml"
	"fmt"
)

// ST_FunctionArgument is a union type
type ST_FunctionArgument struct {
	ST_VariableType ST_VariableType
}

func (m *ST_FunctionArgument) Validate() error {
	return m.ValidateWithPath("")
}

func (m ST_FunctionArgument) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	e.EncodeToken(start)
	if m.ST_VariableType != ST_VariableTypeUnset {
		e.EncodeToken(xml.CharData(m.ST_VariableType.String()))
	}
	return e.EncodeToken(xml.EndElement{Name: start.Name})
}

func (m *ST_FunctionArgument) ValidateWithPath(path string) error {
	mems := []string{}
	if m.ST_VariableType != ST_VariableTypeUnset {
		mems = append(mems, "ST_VariableType")
	}
	if len(mems) > 1 {
		return fmt.Errorf("%s too many members set: %v", path, mems)
	}
	return nil
}

func (m ST_FunctionArgument) String() string {
	if m.ST_VariableType != ST_VariableTypeUnset {
		return m.ST_VariableType.String()
	}
	return ""
}
