package spreadsheet

import (
	"go.devnw.com/ooxml"
	"go.devnw.com/ooxml/schema/soo/sml"
)

type WorkbookProtection struct {
	x *sml.CT_WorkbookProtection
}

// X returns the inner wrapped XML type.
func (p WorkbookProtection) X() *sml.CT_WorkbookProtection {
	return p.x
}

// IsStructureLocked returns whether the workbook structure is locked.
func (p WorkbookProtection) IsStructureLocked() bool {
	return p.x.LockStructureAttr != nil && *p.x.LockStructureAttr
}

// LockStructure controls the locking of the workbook structure.
func (p WorkbookProtection) LockStructure(b bool) {
	if !b {
		p.x.LockStructureAttr = nil
	} else {
		p.x.LockStructureAttr = office.Bool(true)
	}
}

// IsWindowLocked returns whether the workbook windows are locked.
func (p WorkbookProtection) IsWindowLocked() bool {
	return p.x.LockWindowsAttr != nil && *p.x.LockWindowsAttr
}

// LockWindow controls the locking of the workbook windows.
func (p WorkbookProtection) LockWindow(b bool) {
	if !b {
		p.x.LockWindowsAttr = nil
	} else {
		p.x.LockWindowsAttr = office.Bool(true)
	}
}

// PasswordHash returns the hash of the workbook password.
func (p WorkbookProtection) PasswordHash() string {
	if p.x.WorkbookPasswordAttr == nil {
		return ""
	}
	return *p.x.WorkbookPasswordAttr
}

// SetPassword sets the password hash to a hash of the input password.
func (p WorkbookProtection) SetPassword(pw string) {
	p.SetPasswordHash(PasswordHash(pw))
}

// SetPasswordHash sets the password hash to the input.
func (p WorkbookProtection) SetPasswordHash(pwHash string) {
	p.x.WorkbookPasswordAttr = office.String(pwHash)
}
