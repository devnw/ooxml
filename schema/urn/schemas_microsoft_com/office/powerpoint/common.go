package powerpoint

import "go.devnw.com/ooxml"

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	office.RegisterConstructor("urn:schemas-microsoft-com:office:powerpoint", "CT_Empty", NewCT_Empty)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:powerpoint", "CT_Rel", NewCT_Rel)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:powerpoint", "iscomment", NewIscomment)
	office.RegisterConstructor("urn:schemas-microsoft-com:office:powerpoint", "textdata", NewTextdata)
}
