# office
--
    import "."


## Usage

```go
const (
	ContentTypesFilename = "[Content_Types].xml"
	BaseRelsFilename     = "_rels/.rels"
)
```
Common filenames used in zip packages.

```go
const (
	// Common strict
	OfficeDocumentTypeStrict     = "http://purl.oclc.org/ooxml/officeDocument/relationships/officeDocument"
	StylesTypeStrict             = "http://purl.oclc.org/ooxml/officeDocument/relationships/styles"
	ThemeTypeStrict              = "http://purl.oclc.org/ooxml/officeDocument/relationships/theme"
	SettingsTypeStrict           = "http://purl.oclc.org/ooxml/officeDocument/relationships/settings"
	ImageTypeStrict              = "http://purl.oclc.org/ooxml/officeDocument/relationships/image"
	CommentsTypeStrict           = "http://purl.oclc.org/ooxml/officeDocument/relationships/comments"
	ThumbnailTypeStrict          = "http://purl.oclc.org/ooxml/officeDocument/relationships/metadata/thumbnail"
	DrawingTypeStrict            = "http://purl.oclc.org/ooxml/officeDocument/relationships/drawing"
	ChartTypeStrict              = "http://purl.oclc.org/ooxml/officeDocument/relationships/chart"
	ExtendedPropertiesTypeStrict = "http://purl.oclc.org/ooxml/officeDocument/relationships/extendedProperties"
	CustomXMLTypeStrict          = "http://purl.oclc.org/ooxml/officeDocument/relationships/customXml"

	// SML strict
	WorksheetTypeStrict     = "http://purl.oclc.org/ooxml/officeDocument/relationships/worksheet"
	SharedStringsTypeStrict = "http://purl.oclc.org/ooxml/officeDocument/relationships/sharedStrings"
	// Deprecated: Renamed to SharedStringsTypeStrict, will be removed in next major version.
	SharedStingsTypeStrict = SharedStringsTypeStrict
	TableTypeStrict        = "http://purl.oclc.org/ooxml/officeDocument/relationships/table"

	// WML strict
	HeaderTypeStrict      = "http://purl.oclc.org/ooxml/officeDocument/relationships/header"
	FooterTypeStrict      = "http://purl.oclc.org/ooxml/officeDocument/relationships/footer"
	NumberingTypeStrict   = "http://purl.oclc.org/ooxml/officeDocument/relationships/numbering"
	FontTableTypeStrict   = "http://purl.oclc.org/ooxml/officeDocument/relationships/fontTable"
	WebSettingsTypeStrict = "http://purl.oclc.org/ooxml/officeDocument/relationships/webSettings"
	FootNotesTypeStrict   = "http://purl.oclc.org/ooxml/officeDocument/relationships/footnotes"
	EndNotesTypeStrict    = "http://purl.oclc.org/ooxml/officeDocument/relationships/endnotes"

	// PML strict
	SlideTypeStrict = "http://purl.oclc.org/ooxml/officeDocument/relationships/slide"

	// VML strict
	VMLDrawingTypeStrict = "http://purl.oclc.org/ooxml/officeDocument/relationships/vmlDrawing"

	// Common
	OfficeDocumentType     = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/officeDocument"
	StylesType             = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/styles"
	ThemeType              = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/theme"
	ThemeContentType       = "application/vnd.openxmlformats-officedocument.theme+xml"
	SettingsType           = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/settings"
	ImageType              = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/image"
	CommentsType           = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/comments"
	CommentsContentType    = "application/vnd.openxmlformats-officedocument.spreadsheetml.comments+xml"
	ThumbnailType          = "http://schemas.openxmlformats.org/package/2006/relationships/metadata/thumbnail"
	DrawingType            = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/drawing"
	DrawingContentType     = "application/vnd.openxmlformats-officedocument.drawing+xml"
	ChartType              = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/chart"
	ChartContentType       = "application/vnd.openxmlformats-officedocument.drawingml.chart+xml"
	HyperLinkType          = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/hyperlink"
	ExtendedPropertiesType = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/extended-properties"
	CorePropertiesType     = "http://schemas.openxmlformats.org/package/2006/relationships/metadata/core-properties"
	CustomPropertiesType   = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/custom-properties"
	CustomXMLType          = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/customXml"
	TableStylesType        = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/tableStyles"
	ViewPropertiesType     = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/viewProps"

	// SML
	WorksheetType        = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/worksheet"
	WorksheetContentType = "application/vnd.openxmlformats-officedocument.spreadsheetml.worksheet+xml"
	SharedStringsType    = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/sharedStrings"
	// Deprecated: Renamed to SharedStringsType, will be removed in next major version.
	SharedStingsType         = SharedStringsType
	SharedStringsContentType = "application/vnd.openxmlformats-officedocument.spreadsheetml.sharedStrings+xml"
	SMLStyleSheetContentType = "application/vnd.openxmlformats-officedocument.spreadsheetml.styles+xml"
	TableType                = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/table"
	TableContentType         = "application/vnd.openxmlformats-officedocument.spreadsheetml.table+xml"

	// WML
	HeaderType      = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/header"
	FooterType      = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/footer"
	NumberingType   = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/numbering"
	FontTableType   = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/fontTable"
	WebSettingsType = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/webSettings"
	FootNotesType   = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/footnotes"
	EndNotesType    = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/endnotes"

	// PML
	SlideType                  = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/slide"
	SlideContentType           = "application/vnd.openxmlformats-officedocument.presentationml.slide+xml"
	SlideMasterType            = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/slideMaster"
	SlideMasterContentType     = "application/vnd.openxmlformats-officedocument.presentationml.slideMaster+xml"
	SlideLayoutType            = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/slideLayout"
	SlideLayoutContentType     = "application/vnd.openxmlformats-officedocument.presentationml.slideLayout+xml"
	PresentationPropertiesType = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/presProps"
	HandoutMasterType          = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/handoutMaster"
	NotesMasterType            = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/notesMaster"

	// VML
	VMLDrawingType        = "http://schemas.openxmlformats.org/officeDocument/2006/relationships/vmlDrawing"
	VMLDrawingContentType = "application/vnd.openxmlformats-officedocument.vmlDrawing"
)
```
Consts for content types used throughout the package

```go
const MinGoVersion = requires_go_18
```
MinGoVersion is used to cause a compile time error if office is compiled with an
older version of go. Specifically it requires a feature in go 1.8 regarding
collecting all attributes from arbitrary xml used in decode office.XSDAny.

```go
var Log = log.Printf
```
Log is used to log content from within the library. The intent is to use logging
sparingly, preferring to return an error. At the very least this allows
redirecting logs to somewhere more appropriate than stdout.

#### func  AbsoluteFilename

```go
func AbsoluteFilename(dt DocType, typ string, index int) string
```
AbsoluteFilename returns the full path to a file from the root of the zip
container. Index is used in some cases for files which there may be more than
one of (e.g. worksheets/drawings/charts)

#### func  AbsoluteImageFilename

```go
func AbsoluteImageFilename(dt DocType, index int, fileExtension string) string
```
AbsoluteImageFilename returns the full path to an image from the root of the zip
container.

#### func  AddPreserveSpaceAttr

```go
func AddPreserveSpaceAttr(se *xml.StartElement, s string)
```
AddPreserveSpaceAttr adds an xml:space="preserve" attribute to a start element
if it is required for the string s.

#### func  Bool

```go
func Bool(v bool) *bool
```
Bool returns a copy of v as a pointer.

#### func  DisableLogging

```go
func DisableLogging()
```
DisableLogging sets the Log function to a no-op so that any log messages are
silently discarded.

#### func  Float32

```go
func Float32(v float32) *float32
```
Float32 returns a copy of v as a pointer.

#### func  Float64

```go
func Float64(v float64) *float64
```
Float64 returns a copy of v as a pointer.

#### func  Int32

```go
func Int32(v int32) *int32
```
Int32 returns a copy of v as a pointer.

#### func  Int64

```go
func Int64(v int64) *int64
```
Int64 returns a copy of v as a pointer.

#### func  Int8

```go
func Int8(v int8) *int8
```
Int8 returns a copy of v as a pointer.

#### func  NeedsSpacePreserve

```go
func NeedsSpacePreserve(s string) bool
```
NeedsSpacePreserve returns true if the string has leading or trailing space.

#### func  RegisterConstructor

```go
func RegisterConstructor(ns, name string, fn interface{})
```
RegisterConstructor registers a constructor function used for unmarshaling
xsd:any elements.

#### func  RelativeFilename

```go
func RelativeFilename(dt DocType, relToTyp, typ string, index int) string
```
RelativeFilename returns a filename relative to the source file referenced from
a relationships file. Index is used in some cases for files which there may be
more than one of (e.g. worksheets/drawings/charts)

#### func  RelativeImageFilename

```go
func RelativeImageFilename(dt DocType, relToTyp, typ string, index int, fileExtension string) string
```
RelativeImageFilename returns an image filename relative to the source file
referenced from a relationships file. It is identical to RelativeFilename but is
used particularly for images in order to handle different image formats.

#### func  String

```go
func String(v string) *string
```
String returns a copy of v as a pointer.

#### func  Stringf

```go
func Stringf(f string, args ...interface{}) *string
```
Stringf formats according to a format specifier and returns a pointer to the
resulting string.

#### func  Uint16

```go
func Uint16(v uint16) *uint16
```
Uint16 returns a copy of v as a pointer.

#### func  Uint32

```go
func Uint32(v uint32) *uint32
```
Uint32 returns a copy of v as a pointer.

#### func  Uint64

```go
func Uint64(v uint64) *uint64
```
Uint64 returns a copy of v as a pointer.

#### func  Uint8

```go
func Uint8(v uint8) *uint8
```
Uint8 returns a copy of v as a pointer.

#### type Any

```go
type Any interface {
	MarshalXML(e *xml.Encoder, start xml.StartElement) error
	UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
}
```

Any is the interface used for marshaling/unmarshaling xsd:any

#### func  CreateElement

```go
func CreateElement(start xml.StartElement) (Any, error)
```
CreateElement creates an element with the given namespace and name. It is used
to unmarshal some xsd:any elements to the appropriate concrete type.

#### type DocType

```go
type DocType byte
```

DocType represents one of the three document types supported (docx/xlsx/pptx)

```go
const (
	Unknown DocType = iota
	DocTypeSpreadsheet
	DocTypeDocument
	DocTypePresentation
)
```
Document Type constants

#### type XSDAny

```go
type XSDAny struct {
	XMLName xml.Name
	Attrs   []xml.Attr
	Data    []byte
	Nodes   []*XSDAny
}
```

XSDAny is used to marshal/unmarshal xsd:any types in the OOXML schema.

#### func (*XSDAny) MarshalXML

```go
func (x *XSDAny) MarshalXML(e *xml.Encoder, start xml.StartElement) error
```
MarshalXML implements the xml.Marshaler interface.

#### func (*XSDAny) UnmarshalXML

```go
func (x *XSDAny) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error
```
UnmarshalXML implements the xml.Unmarshaler interface.
