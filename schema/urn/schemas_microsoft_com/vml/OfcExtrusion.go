package vml

import (
	"encoding/xml"
	"fmt"
	"strconv"
)

type OfcExtrusion struct {
	OfcCT_Extrusion
}

func NewOfcExtrusion() *OfcExtrusion {
	ret := &OfcExtrusion{}
	ret.OfcCT_Extrusion = *NewOfcCT_Extrusion()
	return ret
}

func (m *OfcExtrusion) MarshalXML(e *xml.Encoder, start xml.StartElement) error {
	return m.OfcCT_Extrusion.MarshalXML(e, start)
}

func (m *OfcExtrusion) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	// initialize to default
	m.OfcCT_Extrusion = *NewOfcCT_Extrusion()
	for _, attr := range start.Attr {
		if attr.Name.Local == "colormode" {
			m.ColormodeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "color" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ColorAttr = &parsed
			continue
		}
		if attr.Name.Local == "type" {
			m.TypeAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "shininess" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			pt := float32(parsed)
			m.ShininessAttr = &pt
			continue
		}
		if attr.Name.Local == "viewpointorigin" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ViewpointoriginAttr = &parsed
			continue
		}
		if attr.Name.Local == "specularity" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SpecularityAttr = &parsed
			continue
		}
		if attr.Name.Local == "plane" {
			m.PlaneAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "diffusity" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.DiffusityAttr = &parsed
			continue
		}
		if attr.Name.Local == "skewamt" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.SkewamtAttr = &parsed
			continue
		}
		if attr.Name.Local == "metal" {
			m.MetalAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "backdepth" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.BackdepthAttr = &parsed
			continue
		}
		if attr.Name.Local == "edge" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.EdgeAttr = &parsed
			continue
		}
		if attr.Name.Local == "lightlevel2" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.Lightlevel2Attr = &parsed
			continue
		}
		if attr.Name.Local == "orientationangle" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			pt := float32(parsed)
			m.OrientationangleAttr = &pt
			continue
		}
		if attr.Name.Local == "on" {
			m.OnAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "lightharsh" {
			m.LightharshAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "lightface" {
			m.LightfaceAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "foredepth" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ForedepthAttr = &parsed
			continue
		}
		if attr.Name.Local == "ext" {
			m.ExtAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "autorotationcenter" {
			m.AutorotationcenterAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "facet" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.FacetAttr = &parsed
			continue
		}
		if attr.Name.Local == "render" {
			m.RenderAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "lightlevel" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LightlevelAttr = &parsed
			continue
		}
		if attr.Name.Local == "brightness" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.BrightnessAttr = &parsed
			continue
		}
		if attr.Name.Local == "skewangle" {
			parsed, err := strconv.ParseFloat(attr.Value, 64)
			if err != nil {
				return err
			}
			pt := float32(parsed)
			m.SkewangleAttr = &pt
			continue
		}
		if attr.Name.Local == "lightposition2" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.Lightposition2Attr = &parsed
			continue
		}
		if attr.Name.Local == "rotationangle" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RotationangleAttr = &parsed
			continue
		}
		if attr.Name.Local == "lightharsh2" {
			m.Lightharsh2Attr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "orientation" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.OrientationAttr = &parsed
			continue
		}
		if attr.Name.Local == "lockrotationcenter" {
			m.LockrotationcenterAttr.UnmarshalXMLAttr(attr)
			continue
		}
		if attr.Name.Local == "rotationcenter" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.RotationcenterAttr = &parsed
			continue
		}
		if attr.Name.Local == "viewpoint" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.ViewpointAttr = &parsed
			continue
		}
		if attr.Name.Local == "lightposition" {
			parsed, err := attr.Value, error(nil)
			if err != nil {
				return err
			}
			m.LightpositionAttr = &parsed
			continue
		}
	}
	// skip any extensions we may find, but don't support
	for {
		tok, err := d.Token()
		if err != nil {
			return fmt.Errorf("parsing OfcExtrusion: %s", err)
		}
		if el, ok := tok.(xml.EndElement); ok && el.Name == start.Name {
			break
		}
	}
	return nil
}

// Validate validates the OfcExtrusion and its children
func (m *OfcExtrusion) Validate() error {
	return m.ValidateWithPath("OfcExtrusion")
}

// ValidateWithPath validates the OfcExtrusion and its children, prefixing error messages with path
func (m *OfcExtrusion) ValidateWithPath(path string) error {
	if err := m.OfcCT_Extrusion.ValidateWithPath(path); err != nil {
		return err
	}
	return nil
}
