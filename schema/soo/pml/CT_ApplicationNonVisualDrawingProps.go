package pml

import (
	"encoding/xml"
	"fmt"
	"strconv"

	"go.devnw.com/ooxml"
	"go.devnw.com/ooxml/schema/soo/dml"
)

type CT_ApplicationNonVisualDrawingProps struct {
	// Is a Photo Album
	IsPhotoAttr *bool
	// Is User Drawn
	UserDrawnAttr *bool
	// Placeholder Shape
	Ph            *CT_Placeholder
	AudioCd       *dml.CT_AudioCD
	WavAudioFile  *dml.CT_EmbeddedWAVAudioFile
	AudioFile     *dml.CT_AudioFile
	VideoFile     *dml.CT_VideoFile
	QuickTimeFile *dml.CT_QuickTimeFile
	// Customer Data List
	CustDataLst *CT_CustomerDataList
	ExtLst      *CT_ExtensionList
}

func NewCT_ApplicationNonVisualDrawingProps() *CT_ApplicationNonVisualDrawingProps {
	ret := &CT_ApplicationNonVisualDrawingProps{}
	return ret
}

func (m *CT_ApplicationNonVisualDrawingProps) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	if m.IsPhotoAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "isPhoto"},
			Value: fmt.Sprintf("%d", b2i(*m.IsPhotoAttr))})
	}
	if m.UserDrawnAttr != nil {
		start.Attr = append(start.Attr, xml.Attr{Name: xml.Name{Local: "userDrawn"},
			Value: fmt.Sprintf("%d", b2i(*m.UserDrawnAttr))})
	}
	e.EncodeToken(start)
	if m.Ph != nil {
		seph := xml.StartElement{Name: xml.Name{Local: "p:ph"}}
		e.EncodeElement(m.Ph, seph)
	}
	if m.AudioCd != nil {
		seaudioCd := xml.StartElement{Name: xml.Name{Local: "p:audioCd"}}
		e.EncodeElement(m.AudioCd, seaudioCd)
	}
	if m.WavAudioFile != nil {
		sewavAudioFile := xml.StartElement{Name: xml.Name{Local: "p:wavAudioFile"}}
		e.EncodeElement(m.WavAudioFile, sewavAudioFile)
	}
	if m.AudioFile != nil {
		seaudioFile := xml.StartElement{Name: xml.Name{Local: "p:audioFile"}}
		e.EncodeElement(m.AudioFile, seaudioFile)
	}
	if m.VideoFile != nil {
		sevideoFile := xml.StartElement{Name: xml.Name{Local: "p:videoFile"}}
		e.EncodeElement(m.VideoFile, sevideoFile)
	}
	if m.QuickTimeFile != nil {
		sequickTimeFile := xml.StartElement{Name: xml.Name{Local: "p:quickTimeFile"}}
		e.EncodeElement(m.QuickTimeFile, sequickTimeFile)
	}
	if m.CustDataLst != nil {
		secustDataLst := xml.StartElement{Name: xml.Name{Local: "p:custDataLst"}}
		e.EncodeElement(m.CustDataLst, secustDataLst)
	}
	if m.ExtLst != nil {
		seextLst := xml.StartElement{Name: xml.Name{Local: "p:extLst"}}
		e.EncodeElement(m.ExtLst, seextLst)
	}
	e.EncodeToken(xml.EndElement{Name: start.Name})
	return nil
}

func (m *CT_ApplicationNonVisualDrawingProps) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	for _, attr := range start.Attr {
		if attr.Name.Local == "isPhoto" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.IsPhotoAttr = &parsed
			continue
		}
		if attr.Name.Local == "userDrawn" {
			parsed, err := strconv.ParseBool(attr.Value)
			if err != nil {
				return err
			}
			m.UserDrawnAttr = &parsed
			continue
		}
	}
lCT_ApplicationNonVisualDrawingProps:
	for {
		tok, err := d.Token()
		if err != nil {
			return err
		}
		switch el := tok.(type) {
		case xml.StartElement:
			switch el.Name {
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "ph"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "ph"}:
				m.Ph = NewCT_Placeholder()
				if err := d.DecodeElement(m.Ph, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "audioCd"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "audioCd"}:
				m.AudioCd = dml.NewCT_AudioCD()
				if err := d.DecodeElement(m.AudioCd, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "wavAudioFile"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "wavAudioFile"}:
				m.WavAudioFile = dml.NewCT_EmbeddedWAVAudioFile()
				if err := d.DecodeElement(m.WavAudioFile, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "audioFile"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "audioFile"}:
				m.AudioFile = dml.NewCT_AudioFile()
				if err := d.DecodeElement(m.AudioFile, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "videoFile"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "videoFile"}:
				m.VideoFile = dml.NewCT_VideoFile()
				if err := d.DecodeElement(m.VideoFile, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/drawingml/2006/main", Local: "quickTimeFile"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/drawingml/main", Local: "quickTimeFile"}:
				m.QuickTimeFile = dml.NewCT_QuickTimeFile()
				if err := d.DecodeElement(m.QuickTimeFile, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "custDataLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "custDataLst"}:
				m.CustDataLst = NewCT_CustomerDataList()
				if err := d.DecodeElement(m.CustDataLst, &el); err != nil {
					return err
				}
			case xml.Name{Space: "http://schemas.openxmlformats.org/presentationml/2006/main", Local: "extLst"},
				xml.Name{Space: "http://purl.oclc.org/ooxml/presentationml/main", Local: "extLst"}:
				m.ExtLst = NewCT_ExtensionList()
				if err := d.DecodeElement(m.ExtLst, &el); err != nil {
					return err
				}
			default:
				office.Log("skipping unsupported element on CT_ApplicationNonVisualDrawingProps %v", el.Name)
				if err := d.Skip(); err != nil {
					return err
				}
			}
		case xml.EndElement:
			break lCT_ApplicationNonVisualDrawingProps
		case xml.CharData:
		}
	}
	return nil
}

// Validate validates the CT_ApplicationNonVisualDrawingProps and its children
func (m *CT_ApplicationNonVisualDrawingProps) Validate() error {
	return m.ValidateWithPath("CT_ApplicationNonVisualDrawingProps")
}

// ValidateWithPath validates the CT_ApplicationNonVisualDrawingProps and its children, prefixing error messages with path
func (m *CT_ApplicationNonVisualDrawingProps) ValidateWithPath(path string) error {
	if m.Ph != nil {
		if err := m.Ph.ValidateWithPath(path + "/Ph"); err != nil {
			return err
		}
	}
	if m.AudioCd != nil {
		if err := m.AudioCd.ValidateWithPath(path + "/AudioCd"); err != nil {
			return err
		}
	}
	if m.WavAudioFile != nil {
		if err := m.WavAudioFile.ValidateWithPath(path + "/WavAudioFile"); err != nil {
			return err
		}
	}
	if m.AudioFile != nil {
		if err := m.AudioFile.ValidateWithPath(path + "/AudioFile"); err != nil {
			return err
		}
	}
	if m.VideoFile != nil {
		if err := m.VideoFile.ValidateWithPath(path + "/VideoFile"); err != nil {
			return err
		}
	}
	if m.QuickTimeFile != nil {
		if err := m.QuickTimeFile.ValidateWithPath(path + "/QuickTimeFile"); err != nil {
			return err
		}
	}
	if m.CustDataLst != nil {
		if err := m.CustDataLst.ValidateWithPath(path + "/CustDataLst"); err != nil {
			return err
		}
	}
	if m.ExtLst != nil {
		if err := m.ExtLst.ValidateWithPath(path + "/ExtLst"); err != nil {
			return err
		}
	}
	return nil
}
