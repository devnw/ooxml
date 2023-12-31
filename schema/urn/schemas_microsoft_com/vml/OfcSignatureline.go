package vml

import (
	"encoding/xml"
	"fmt"
)

type OfcSignatureline struct {
	OfcCT_SignatureLine
}

func NewOfcSignatureline() *OfcSignatureline {
	ret := &OfcSignatureline{}
	ret.OfcCT_SignatureLine = *NewOfcCT_SignatureLine()
	return ret
}

func (m *OfcSignatureline) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns"}, Value: "urn:schemas-microsoft-com:office:office"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:o"}, Value: "urn:schemas-microsoft-com:office:office"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:r"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/relationships"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:s"}, Value: "http://schemas.openxmlformats.org/officeDocument/2006/sharedTypes"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:v"}, Value: "urn:schemas-microsoft-com:vml"})
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "xmlns:xml"}, Value: "http://www.w3.org/XML/1998/namespace"})
	start.Name.Local = "o:signatureline"
	return m.OfcCT_SignatureLine.MarshalXML(e, start)
}

func (m *OfcSignatureline) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.OfcCT_SignatureLine = *NewOfcCT_SignatureLine()
	for _, attr := range start.Attr {
		if attr.Name.Local == "suggestedsigner" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SuggestedsignerAttr = &parsed
			continue
		}
		if attr.Name.Local == "id" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.IdAttr = &parsed
			continue
		}
		if attr.Name.Local == "provid" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ProvidAttr = &parsed
			continue
		}
		if attr.Name.Local == "signinginstructionsset" {
			m.SigninginstructionssetAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "allowcomments" {
			m.AllowcommentsAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "showsigndate" {
			m.ShowsigndateAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "issignatureline" {
			m.IssignaturelineAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "suggestedsigner2" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.Suggestedsigner2Attr = &parsed
			continue
		}
		if attr.Name.Local == "suggestedsigneremail" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SuggestedsigneremailAttr = &parsed
			continue
		}
		if attr.Name.Local == "signinginstructions" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SigninginstructionsAttr = &parsed
			continue
		}
		if attr.Name.Local == "addlxml" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.AddlxmlAttr = &parsed
			continue
		}
		if attr.Name.Local == "sigprovurl" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SigprovurlAttr = &parsed
			continue
		}
		if attr.Name.Local == "ext" {
			m.ExtAttr.UnmarshalXMLAttr(attr)
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing OfcSignatureline: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the OfcSignatureline and its children
func (m *OfcSignatureline) Validate() error {
	return m.ValidateWithPath("OfcSignatureline")
}

// ValidateWithPath validates the OfcSignatureline and its children, prefixing error messages with path
func (m *OfcSignatureline) ValidateWithPath(path string) error {
	if err := m.OfcCT_SignatureLine.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
