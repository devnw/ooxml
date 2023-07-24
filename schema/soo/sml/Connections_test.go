package sml_test

import (
	"encoding/xml"
	"testing"

	"go.devnw.com/ooxml/schema/soo/sml"
)

func TestConnectionsConstructor(t *testing.T) {
	v := sml.NewConnections()
	if v == nil {
		t.Errorf("sml.NewConnections must return a non-nil value")
	}
	if err := v.Validate(); err != nil {
		t.Errorf("newly constructed sml.Connections should validate: %s", err)
	}
}

func TestConnectionsMarshalUnmarshal(t *testing.T) {
	v := sml.NewConnections()
	buf, _ := xml.Marshal(v)
	v2 := sml.NewConnections()
	xml.Unmarshal(buf, v2)
}
