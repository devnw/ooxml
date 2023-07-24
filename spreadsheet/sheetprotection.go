package spreadsheet

import (
	"go.devnw.com/ooxml"
	"go.devnw.com/ooxml/schema/soo/sml"
)

type SheetProtection struct {
	x *sml.CT_SheetProtection
}

// X returns the inner wrapped XML type.
func (p SheetProtection) X() *sml.CT_SheetProtection {
	return p.x
}

// IsSheetLocked returns whether the sheet is locked.
func (p SheetProtection) IsSheetLocked() bool {
	return p.x.SheetAttr != nil && *p.x.SheetAttr
}

// LockSheet controls the locking of the sheet.
func (p SheetProtection) LockSheet(b bool) {
	if !b {
		p.x.SheetAttr = nil
	} else {
		p.x.SheetAttr = office.Bool(true)
	}
}

// IsSheetLocked returns whether the sheet objects are locked.
func (p SheetProtection) IsObjectLocked() bool {
	return p.x.ObjectsAttr != nil && *p.x.ObjectsAttr
}

// LockObject controls the locking of the sheet objects.
func (p SheetProtection) LockObject(b bool) {
	if !b {
		p.x.ObjectsAttr = nil
	} else {
		p.x.ObjectsAttr = office.Bool(true)
	}
}

// PasswordHash returns the hash of the workbook password.
func (p SheetProtection) PasswordHash() string {
	if p.x.PasswordAttr == nil {
		return ""
	}
	return *p.x.PasswordAttr
}

// SetPassword sets the password hash to a hash of the input password.
func (p SheetProtection) SetPassword(pw string) {
	p.SetPasswordHash(PasswordHash(pw))
}

// SetPasswordHash sets the password hash to the input.
func (p SheetProtection) SetPasswordHash(pwHash string) {
	p.x.PasswordAttr = office.String(pwHash)
}
