package update

// UpdateAction is the type for update types constants.
type UpdateAction byte

const (
	// UpdateActionRemoveColumn means updating references after removing a column.
	UpdateActionRemoveColumn UpdateAction = iota
)

// UpdateQuery contains terms of how to update references after removing row/column.
type UpdateQuery struct {
	// UpdateType is one of the update types like UpdateActionRemoveColumn.
	UpdateType UpdateAction

	// ColumnIdx is the index of the column removed.
	ColumnIdx uint32

	// SheetToUpdate contains the name of the sheet on which removing happened.
	SheetToUpdate string

	// UpdateCurrentSheet is true if references without sheet prefix should be updated as well.
	UpdateCurrentSheet bool
}
