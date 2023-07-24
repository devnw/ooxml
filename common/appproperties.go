package common

import (
	"fmt"

	office "go.devnw.com/ooxml"
	"go.devnw.com/ooxml/schema/soo/ofc/extended_properties"
)

// AppProperties contains properties specific to the document and the
// application that created it.
type AppProperties struct {
	x *extended_properties.Properties
}

// NewAppProperties constructs a new AppProperties.
func NewAppProperties() AppProperties {
	p := AppProperties{x: extended_properties.NewProperties()}
	p.SetCompany("Developer Network GPL v3 OOXML Go Library")
	p.SetApplication("go.devnw.com/ooxml")
	p.SetDocSecurity(0)
	p.SetLinksUpToDate(false)

	var major, minor, patch int64
	fmt.Sscanf(Version, "%d.%d.%d", &major, &minor, &patch)
	f := float64(major) + float64(minor)/10000.0
	p.SetApplicationVersion(fmt.Sprintf("%07.4f", f))
	return p
}

// Application returns the name of the application that created the document.
// For gooxml created documents, it defaults to go.devnw.com/ooxml
func (a AppProperties) Application() string {
	if a.x.Application != nil {
		return *a.x.Application
	}
	return ""
}

// SetLinksUpToDate sets the links up to date flag.
func (a AppProperties) SetLinksUpToDate(v bool) {
	a.x.LinksUpToDate = office.Bool(v)
}

// SetDocSecurity sets the document security flag.
func (a AppProperties) SetDocSecurity(v int32) {
	a.x.DocSecurity = office.Int32(v)
}

// SetApplication sets the name of the application that created the document.
func (a AppProperties) SetApplication(s string) {
	a.x.Application = &s
}

// ApplicationVersion returns the version of the application that created the
// document.
func (a AppProperties) ApplicationVersion() string {
	if a.x.AppVersion != nil {
		return *a.x.AppVersion
	}
	return ""
}

// SetApplicationVersion sets the version of the application that created the
// document.  Per MS, the verison string mut be in the form 'XX.YYYY'.
func (a AppProperties) SetApplicationVersion(s string) {
	a.x.AppVersion = &s
}

// X returns the inner wrapped XML type.
func (a AppProperties) X() *extended_properties.Properties {
	return a.x
}

// Company returns the name of the company that created the document.
// For gooxml created documents, it defaults to go.devnw.com/ooxml
func (a AppProperties) Company() string {
	if a.x.Company != nil {
		return *a.x.Company
	}
	return ""
}

// SetCompany sets the name of the company that created the document.
func (a AppProperties) SetCompany(s string) {
	a.x.Company = &s
}
