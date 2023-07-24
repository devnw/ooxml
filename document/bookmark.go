package document

import "go.devnw.com/ooxml/schema/soo/wml"

// Bookmark is a bookmarked location within a document that can be referenced
// with a hyperlink.
type Bookmark struct {
	x *wml.CT_Bookmark
}

// X returns the inner wrapped XML type.
func (b Bookmark) X() *wml.CT_Bookmark {
	return b.x
}

// SetName sets the name of the bookmark. This is the name that is used to
// reference the bookmark from hyperlinks.
func (b Bookmark) SetName(name string) {
	b.x.NameAttr = name
}

// Name returns the name of the bookmark whcih is the document unique ID that
// identifies the bookmark.
func (b Bookmark) Name() string {
	return b.x.NameAttr
}
