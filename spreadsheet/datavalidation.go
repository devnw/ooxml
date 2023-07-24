package spreadsheet

import (
	"go.devnw.com/ooxml"
	"go.devnw.com/ooxml/schema/soo/sml"
)

// DataValidation controls cell validation
type DataValidation struct {
	x *sml.CT_DataValidation
}

// X returns the inner wrapped XML type.
func (d DataValidation) X() *sml.CT_DataValidation {
	return d.x
}

func (d DataValidation) clear() {
	d.x.Formula1 = office.String("0")
	d.x.Formula2 = office.String("0")
}

// SetAllowBlank controls if blank values are accepted.
func (d DataValidation) SetAllowBlank(b bool) {
	if !b {
		d.x.AllowBlankAttr = nil
	} else {
		d.x.AllowBlankAttr = office.Bool(true)
	}
}

func (d DataValidation) SetList() DataValidationList {
	d.clear()
	d.x.TypeAttr = sml.ST_DataValidationTypeList
	d.x.OperatorAttr = sml.ST_DataValidationOperatorEqual
	return DataValidationList{d.x}
}

func (d DataValidation) SetComparison(t DVCompareType, op DVCompareOp) DataValidationCompare {
	d.clear()
	d.x.TypeAttr = sml.ST_DataValidationType(t)
	d.x.OperatorAttr = sml.ST_DataValidationOperator(op)
	return DataValidationCompare{d.x}
}

// SetRange sets the cell or range of cells that the validation should apply to.
// It can be a single cell (e.g. "A1") or a range of cells (e.g. "A1:B5")
func (d DataValidation) SetRange(cellRange string) {
	d.x.SqrefAttr = sml.ST_Sqref{cellRange}
}
