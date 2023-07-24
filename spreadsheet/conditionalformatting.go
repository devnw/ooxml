package spreadsheet

import "go.devnw.com/ooxml/schema/soo/sml"

// ConditionalFormatting controls the formatting styles and rules for a range of
// cells with the same conditional formatting.
type ConditionalFormatting struct {
	x *sml.CT_ConditionalFormatting
}

// X returns the inner wrapped XML type.
func (c ConditionalFormatting) X() *sml.CT_ConditionalFormatting {
	return c.x
}

// AddRule adds and returns a new rule that can be configured.
func (c ConditionalFormatting) AddRule() ConditionalFormattingRule {
	rule := sml.NewCT_CfRule()
	c.x.CfRule = append(c.x.CfRule, rule)
	r := ConditionalFormattingRule{rule}
	r.InitializeDefaults()
	r.SetPriority(int32(len(c.x.CfRule) + 1))
	return r
}
