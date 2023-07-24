package document

import (
	"go.devnw.com/ooxml"
	"go.devnw.com/ooxml/measurement"
	"go.devnw.com/ooxml/schema/soo/wml"
)

// TableStyleProperties are table properties as defined in a style.
type TableStyleProperties struct {
	x *wml.CT_TblPrBase
}

// X returns the inner wrapped XML type.
func (t TableStyleProperties) X() *wml.CT_TblPrBase {
	return t.x
}

// SetRowBandSize sets the number of Rows in the row band
func (t TableStyleProperties) SetRowBandSize(rows int64) {
	t.x.TblStyleRowBandSize = wml.NewCT_DecimalNumber()
	t.x.TblStyleRowBandSize.ValAttr = rows
}

// SetColumnBandSize sets the number of Columns in the column band
func (t TableStyleProperties) SetColumnBandSize(cols int64) {
	t.x.TblStyleColBandSize = wml.NewCT_DecimalNumber()
	t.x.TblStyleColBandSize.ValAttr = cols
}

// SetTableIndent sets the Table Indent from the Leading Margin
func (t TableStyleProperties) SetTableIndent(ind measurement.Distance) {
	t.x.TblInd = wml.NewCT_TblWidth()
	t.x.TblInd.TypeAttr = wml.ST_TblWidthDxa
	t.x.TblInd.WAttr = &wml.ST_MeasurementOrPercent{}
	t.x.TblInd.WAttr.ST_DecimalNumberOrPercent = &wml.ST_DecimalNumberOrPercent{}
	t.x.TblInd.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = office.Int64(int64(ind / measurement.Dxa))
}

// SetCellSpacingAuto sets the cell spacing within a table to automatic.
func (t TableStyleProperties) SetCellSpacingAuto() {
	t.x.TblCellSpacing = wml.NewCT_TblWidth()
	t.x.TblCellSpacing.TypeAttr = wml.ST_TblWidthAuto
}

// SetCellSpacingPercent sets the cell spacing within a table to a percent width.
func (t TableStyleProperties) SetCellSpacingPercent(pct float64) {
	t.x.TblCellSpacing = wml.NewCT_TblWidth()
	t.x.TblCellSpacing.TypeAttr = wml.ST_TblWidthPct
	t.x.TblCellSpacing.WAttr = &wml.ST_MeasurementOrPercent{}
	t.x.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent = &wml.ST_DecimalNumberOrPercent{}
	t.x.TblCellSpacing.WAttr.ST_DecimalNumberOrPercent.ST_UnqualifiedPercentage = office.Int64(int64(pct * 50))
}

// Borders allows manipulation of the table borders.
func (t TableStyleProperties) Borders() TableBorders {
	if t.x.TblBorders == nil {
		t.x.TblBorders = wml.NewCT_TblBorders()
	}
	return TableBorders{t.x.TblBorders}
}
