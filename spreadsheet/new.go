package spreadsheet

import (
	"runtime"

	"go.devnw.com/ooxml"
	"go.devnw.com/ooxml/common"
	"go.devnw.com/ooxml/schema/soo/sml"
)

// New constructs a new workbook.
func New() *Workbook {
	wb := &Workbook{}
	wb.x = sml.NewWorkbook()

	runtime.SetFinalizer(wb, workbookFinalizer)

	wb.AppProperties = common.NewAppProperties()
	wb.CoreProperties = common.NewCoreProperties()
	wb.StyleSheet = NewStyleSheet(wb)

	wb.Rels = common.NewRelationships()
	wb.wbRels = common.NewRelationships()

	wb.Rels.AddRelationship(office.RelativeFilename(office.DocTypeSpreadsheet, "", office.ExtendedPropertiesType, 0), office.ExtendedPropertiesType)
	wb.Rels.AddRelationship(office.RelativeFilename(office.DocTypeSpreadsheet, "", office.CorePropertiesType, 0), office.CorePropertiesType)
	wb.Rels.AddRelationship(office.RelativeFilename(office.DocTypeSpreadsheet, "", office.OfficeDocumentType, 0), office.OfficeDocumentType)
	wb.wbRels.AddRelationship(office.RelativeFilename(office.DocTypeSpreadsheet, office.OfficeDocumentType, office.StylesType, 0), office.StylesType)

	wb.ContentTypes = common.NewContentTypes()
	wb.ContentTypes.AddDefault("vml", office.VMLDrawingContentType)
	wb.ContentTypes.AddOverride(office.AbsoluteFilename(office.DocTypeSpreadsheet, office.OfficeDocumentType, 0), "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet.main+xml")
	wb.ContentTypes.AddOverride(office.AbsoluteFilename(office.DocTypeSpreadsheet, office.StylesType, 0), office.SMLStyleSheetContentType)

	wb.SharedStrings = NewSharedStrings()
	wb.ContentTypes.AddOverride(office.AbsoluteFilename(office.DocTypeSpreadsheet, office.SharedStringsType, 0), office.SharedStringsContentType)
	wb.wbRels.AddRelationship(office.RelativeFilename(office.DocTypeSpreadsheet, office.OfficeDocumentType, office.SharedStringsType, 0), office.SharedStringsType)

	return wb
}
