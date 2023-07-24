package spreadsheet

import (
	"go.devnw.com/ooxml"
	"go.devnw.com/ooxml/schema/soo/sml"
)

// IconScale maps values to icons.
type IconScale struct {
	x *sml.CT_IconSet
}

// X returns the inner wrapped XML type.
func (c IconScale) X() *sml.CT_IconSet {
	return c.x
}

// SetIcons sets the icon set to use for display.
func (c IconScale) SetIcons(t sml.ST_IconSetType) {
	c.x.IconSetAttr = t
}

// AddFormatValue adds a format value to be used in determining which icons to display.
func (c IconScale) AddFormatValue(t sml.ST_CfvoType, val string) {
	v := sml.NewCT_Cfvo()
	v.TypeAttr = t
	v.ValAttr = office.String(val)
	c.x.Cfvo = append(c.x.Cfvo, v)
}
