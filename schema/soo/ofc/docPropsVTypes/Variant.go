package docPropsVTypes

import (
	"encoding/xml"
	"time"

	"go.devnw.com/ooxml"
)

type Variant struct {
	CT_Variant
}

func NewVariant() *Variant {
	ret := &Variant{}
	ret.CT_Variant = *NewCT_Variant()
	return ret
}

func (m *Variant) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return m.CT_Variant.MarshalXML(e, start)
}

func (m *Variant) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.CT_Variant = *NewCT_Variant()
lVariant:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "variant"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "variant"}:
				m.Variant = NewVariant()
				if err := d.DecodeElement(m.Variant, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "vector"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "vector"}:
				m.Vector = NewVector()
				if err := d.DecodeElement(m.Vector, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "array"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "array"}:
				m.Array = NewArray()
				if err := d.DecodeElement(m.Array, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "blob"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "blob"}:
				m.Blob = new(string)
				if err := d.DecodeElement(m.Blob, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "oblob"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "oblob"}:
				m.Oblob = new(string)
				if err := d.DecodeElement(m.Oblob, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "empty"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "empty"}:
				m.Empty = NewEmpty()
				if err := d.DecodeElement(m.Empty, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "null"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "null"}:
				m.Null = NewNull()
				if err := d.DecodeElement(m.Null, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "i1"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "i1"}:
				m.I1 = new(int8)
				if err := d.DecodeElement(m.I1, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "i2"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "i2"}:
				m.I2 = new(int16)
				if err := d.DecodeElement(m.I2, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "i4"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "i4"}:
				m.I4 = new(int32)
				if err := d.DecodeElement(m.I4, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "i8"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "i8"}:
				m.I8 = new(int64)
				if err := d.DecodeElement(m.I8, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "int"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "int"}:
				m.Int = new(int32)
				if err := d.DecodeElement(m.Int, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "ui1"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "ui1"}:
				m.Ui1 = new(uint8)
				if err := d.DecodeElement(m.Ui1, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "ui2"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "ui2"}:
				m.Ui2 = new(uint16)
				if err := d.DecodeElement(m.Ui2, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "ui4"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "ui4"}:
				m.Ui4 = new(uint32)
				if err := d.DecodeElement(m.Ui4, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "ui8"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "ui8"}:
				m.Ui8 = new(uint64)
				if err := d.DecodeElement(m.Ui8, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "uint"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "uint"}:
				m.Uint = new(uint32)
				if err := d.DecodeElement(m.Uint, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "r4"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "r4"}:
				m.R4 = new(float32)
				if err := d.DecodeElement(m.R4, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "r8"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "r8"}:
				m.R8 = new(float64)
				if err := d.DecodeElement(m.R8, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "decimal"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "decimal"}:
				m.Decimal = new(float64)
				if err := d.DecodeElement(m.Decimal, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "lpstr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "lpstr"}:
				m.Lpstr = new(string)
				if err := d.DecodeElement(m.Lpstr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "lpwstr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "lpwstr"}:
				m.Lpwstr = new(string)
				if err := d.DecodeElement(m.Lpwstr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "bstr"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "bstr"}:
				m.Bstr = new(string)
				if err := d.DecodeElement(m.Bstr, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "date"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "date"}:
				m.Date = new(time.Time)
				if err := d.DecodeElement(m.Date, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "filetime"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "filetime"}:
				m.Filetime = new(time.Time)
				if err := d.DecodeElement(m.Filetime, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "bool"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "bool"}:
				m.Bool = new(bool)
				if err := d.DecodeElement(m.Bool, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "cy"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "cy"}:
				m.Cy = new(string)
				if err := d.DecodeElement(m.Cy, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "error"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "error"}:
				m.Error = new(string)
				if err := d.DecodeElement(m.Error, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "stream"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "stream"}:
				m.Stream = new(string)
				if err := d.DecodeElement(m.Stream, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "ostream"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "ostream"}:
				m.Ostream = new(string)
				if err := d.DecodeElement(m.Ostream, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "storage"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "storage"}:
				m.Storage = new(string)
				if err := d.DecodeElement(m.Storage, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "ostorage"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "ostorage"}:
				m.Ostorage = new(string)
				if err := d.DecodeElement(m.Ostorage, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "vstream"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "vstream"}:
				m.Vstream = NewVstream()
				if err := d.DecodeElement(m.Vstream, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/officeDocument/2006/docPropsVTypes", Local: "clsid"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/officeDocument/docPropsVTypes", Local: "clsid"}:
				m.Clsid = new(string)
				if err := d.DecodeElement(m.Clsid, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on Variant %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lVariant
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the Variant and its children
func (m *Variant) Validate() error {
	return m.ValidateWithPath("Variant")
}

// ValidateWithPath validates the Variant and its children, prefixing error messages with path
func (m *Variant) ValidateWithPath(path string) error {
	if err := m.CT_Variant.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
