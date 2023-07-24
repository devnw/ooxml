package document

import (
	"go.devnw.com/ooxml"
	"go.devnw.com/ooxml/schema/soo/wml"
)

// Settings controls the document settings.
type Settings struct {
	x *wml.Settings
}

// NewSettings constructs a new empty Settings
func NewSettings() Settings {
	s := wml.NewSettings()
	s.Compat = wml.NewCT_Compat()
	stng := wml.NewCT_CompatSetting()
	stng.NameAttr = office.String("compatibilityMode")
	stng.UriAttr = office.String("http://schemas.microsoft.com/office/word")
	stng.ValAttr = office.String("15")
	s.Compat.CompatSetting = append(s.Compat.CompatSetting, stng)
	return Settings{s}
}

// X returns the inner wrapped XML type.
func (s Settings) X() *wml.Settings {
	return s.x
}

// SetUpdateFieldsOnOpen controls if fields are recalculated upon opening the
// document. This is useful for things like a table of contents as the library
// only adds the field code and relies on Word/LibreOffice to actually compute
// the content.
func (s Settings) SetUpdateFieldsOnOpen(b bool) {
	if !b {
		s.x.UpdateFields = nil
	} else {
		s.x.UpdateFields = wml.NewCT_OnOff()
	}
}

// RemoveMailMerge removes any mail merge settings
func (s Settings) RemoveMailMerge() {
	s.x.MailMerge = nil
}
