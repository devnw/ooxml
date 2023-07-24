package spreadsheet

import (
	"strings"

	"go.devnw.com/ooxml"
	"go.devnw.com/ooxml/schema/soo/sml"
)

// DataValidationList is just a view on a DataValidation configured as a list.
// It presents a drop-down combo box for spreadsheet users to select values. The
// contents of the dropdown can either pull from a rang eof cells (SetRange) or
// specified directly (SetValues).
type DataValidationList struct {
	x *sml.CT_DataValidation
}

// SetRange sets the range that contains the possible values. This is incompatible with SetValues.
func (d DataValidationList) SetRange(cellRange string) {
	d.x.Formula1 = office.String(cellRange)
	d.x.Formula2 = office.String("0")
}

// SetValues sets the possible values. This is incompatible with SetRange.
func (d DataValidationList) SetValues(values []string) {
	d.x.Formula1 = office.String("\"" + strings.Join(values, ",") + "\"")
	d.x.Formula2 = office.String("0")
}
