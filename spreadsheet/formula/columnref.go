package formula

import "go.devnw.com/ooxml/spreadsheet/reference"

// updateColumnToLeft gets a column reference string representation like JJ, parses it and makes a string representation of a new reference with respect to the update type in the case of a column to the left of this reference was removed (e.g. JI).
func updateColumnToLeft(column string, colIdxToRemove uint32) string {
	colIdx := reference.ColumnToIndex(column)
	if colIdx == colIdxToRemove {
		return "#REF!"
	} else if colIdx > colIdxToRemove {
		return reference.IndexToColumn(colIdx - 1)
	} else {
		return column
	}
}
