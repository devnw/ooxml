package office_test

import "testing"
import "go.devnw.com/ooxml"

func TestWMLFilenames(t *testing.T) {
	td := []struct {
		Idx    int
		Type   string
		ExpAbs string
	}{
		{0, office.CorePropertiesType, "docProps/core.xml"},
		{0, office.ExtendedPropertiesType, "docProps/app.xml"},
		{0, office.ThumbnailType, "docProps/thumbnail.jpeg"},
		{0, office.StylesType, "word/styles.xml"},

		{0, office.OfficeDocumentType, "word/document.xml"},
		{0, office.FontTableType, "word/fontTable.xml"},
		{0, office.EndNotesType, "word/endnotes.xml"},
		{0, office.FootNotesType, "word/footnotes.xml"},
		{0, office.NumberingType, "word/numbering.xml"},
		{0, office.WebSettingsType, "word/webSettings.xml"},
		{0, office.SettingsType, "word/settings.xml"},
		{23, office.HeaderType, "word/header23.xml"},
		{15, office.FooterType, "word/footer15.xml"},
		{1, office.ThemeType, "word/theme/theme1.xml"},
	}
	for _, tc := range td {
		abs := office.AbsoluteFilename(office.DocTypeDocument, tc.Type, tc.Idx)
		if abs != tc.ExpAbs {
			t.Errorf("expected absolute filename of %s for document %s/%d, got %s", tc.ExpAbs, tc.Type, tc.Idx, abs)
		}
	}
}

func TestSMLFilenames(t *testing.T) {
	td := []struct {
		Idx    int
		Type   string
		ExpAbs string
	}{
		{0, office.CorePropertiesType, "docProps/core.xml"},
		{0, office.ExtendedPropertiesType, "docProps/app.xml"},
		{0, office.ThumbnailType, "docProps/thumbnail.jpeg"},
		{0, office.StylesType, "xl/styles.xml"},

		{0, office.OfficeDocumentType, "xl/workbook.xml"},
		{15, office.ChartType, "xl/charts/chart15.xml"},
		{12, office.DrawingType, "xl/drawings/drawing12.xml"},
		{13, office.TableType, "xl/tables/table13.xml"},
		{2, office.CommentsType, "xl/comments2.xml"},
		{15, office.WorksheetType, "xl/worksheets/sheet15.xml"},
		{2, office.VMLDrawingType, "xl/drawings/vmlDrawing2.vml"},
		{0, office.SharedStringsType, "xl/sharedStrings.xml"},
		{1, office.ThemeType, "xl/theme/theme1.xml"},
		{2, office.ImageType, "xl/media/image2.png"},
	}
	for _, tc := range td {
		abs := office.AbsoluteFilename(office.DocTypeSpreadsheet, tc.Type, tc.Idx)
		if abs != tc.ExpAbs {
			t.Errorf("expected absolute filename of %s for document %s/%d, got %s", tc.ExpAbs, tc.Type, tc.Idx, abs)
		}
	}
}

func TestPMLFilenames(t *testing.T) {
	td := []struct {
		Idx    int
		Type   string
		ExpAbs string
	}{
		{0, office.CorePropertiesType, "docProps/core.xml"},
		{0, office.ExtendedPropertiesType, "docProps/app.xml"},
		{0, office.ThumbnailType, "docProps/thumbnail.jpeg"},
		{0, office.StylesType, "ppt/styles.xml"},

		{0, office.OfficeDocumentType, "ppt/presentation.xml"},
		{4, office.SlideType, "ppt/slides/slide4.xml"},
		{5, office.SlideLayoutType, "ppt/slideLayouts/slideLayout5.xml"},
		{6, office.SlideMasterType, "ppt/slideMasters/slideMaster6.xml"},
		{7, office.ThemeType, "ppt/theme/theme7.xml"},
	}
	for _, tc := range td {
		abs := office.AbsoluteFilename(office.DocTypePresentation, tc.Type, tc.Idx)
		if abs != tc.ExpAbs {
			t.Errorf("expected absolute filename of %s for document %s/%d, got %s", tc.ExpAbs, tc.Type, tc.Idx, abs)
		}
	}
}
