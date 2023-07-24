package elements

import "go.devnw.com/ooxml"

// init registers constructor functions for dynamically creating elements based off the XML namespace and name
func init() {
	office.RegisterConstructor("http://purl.org/dc/elements/1.1/", "SimpleLiteral", NewSimpleLiteral)
	office.RegisterConstructor("http://purl.org/dc/elements/1.1/", "elementContainer", NewElementContainer)
	office.RegisterConstructor("http://purl.org/dc/elements/1.1/", "any", NewAny)
	office.RegisterConstructor("http://purl.org/dc/elements/1.1/", "elementsGroup", NewElementsGroup)
}
