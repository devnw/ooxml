package terms

import "go.devnw.com/ooxml"

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	office.RegisterConstructor("http://purl.org/dc/terms/", "LCSH", NewLCSH)
	office.RegisterConstructor("http://purl.org/dc/terms/", "MESH", NewMESH)
	office.RegisterConstructor("http://purl.org/dc/terms/", "DDC", NewDDC)
	office.RegisterConstructor("http://purl.org/dc/terms/", "LCC", NewLCC)
	office.RegisterConstructor("http://purl.org/dc/terms/", "UDC", NewUDC)
	office.RegisterConstructor("http://purl.org/dc/terms/", "Period", NewPeriod)
	office.RegisterConstructor("http://purl.org/dc/terms/", "W3CDTF", NewW3CDTF)
	office.RegisterConstructor("http://purl.org/dc/terms/", "DCMIType", NewDCMIType)
	office.RegisterConstructor("http://purl.org/dc/terms/", "IMT", NewIMT)
	office.RegisterConstructor("http://purl.org/dc/terms/", "URI", NewURI)
	office.RegisterConstructor("http://purl.org/dc/terms/", "ISO639-2", NewISO639_2)
	office.RegisterConstructor("http://purl.org/dc/terms/", "RFC1766", NewRFC1766)
	office.RegisterConstructor("http://purl.org/dc/terms/", "RFC3066", NewRFC3066)
	office.RegisterConstructor("http://purl.org/dc/terms/", "Point", NewPoint)
	office.RegisterConstructor("http://purl.org/dc/terms/", "ISO3166", NewISO3166)
	office.RegisterConstructor("http://purl.org/dc/terms/", "Box", NewBox)
	office.RegisterConstructor("http://purl.org/dc/terms/", "TGN", NewTGN)
	office.RegisterConstructor("http://purl.org/dc/terms/", "elementOrRefinementContainer", NewElementOrRefinementContainer)
	office.RegisterConstructor("http://purl.org/dc/terms/", "elementsAndRefinementsGroup", NewElementsAndRefinementsGroup)
}
