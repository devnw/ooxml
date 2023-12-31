package dml

import (
	"encoding/xml"
	"fmt"
)

type CT_EmbeddedWAVAudioFile struct {
	EmbedAttr string
	NameAttr  *string
}

func NewCT_EmbeddedWAVAudioFile() *CT_EmbeddedWAVAudioFile {
	ret := &CT_EmbeddedWAVAudioFile{}
	return ret
}

func (m *CT_EmbeddedWAVAudioFile) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "r:embed"},
		Value: fmt.Sprintf("%v", m.EmbedAttr)})
	if m.NameAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "name"},
			Value: fmt.Sprintf("%v", *m.NameAttr)})
	}
	e.EncodeToken(start)
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_EmbeddedWAVAudioFile) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Space == "http://schemas.openxmlformats.org/officeDocument/2006/relationships" && attr.Name.Local == "embed" ||
			attr.Name.Space == "http://purl.oclc.org/ooxml/officeDocument/relationships" && attr.Name.Local == "embed" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.EmbedAttr = parsed
			continue
		}
		if attr.Name.Local == "name" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.NameAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing CT_EmbeddedWAVAudioFile: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the CT_EmbeddedWAVAudioFile and its children
func (m *CT_EmbeddedWAVAudioFile) Validate() error {
	return m.ValidateWithPath("CT_EmbeddedWAVAudioFile")
}

// ValidateWithPath validates the CT_EmbeddedWAVAudioFile and its children, prefixing error messages with path
func (m *CT_EmbeddedWAVAudioFile) ValidateWithPath(path string) error {
	return nil
}
