package common_test

import (
	"bytes"
	"encoding/xml"
	"os"
	"testing"

	"go.devnw.com/ooxml/common"
	"go.devnw.com/ooxml/testhelper"
	"go.devnw.com/ooxml/zippkg"
)

func TestThemeUnmarshal(t *testing.T) {
	f, err := os.Open("testdata/theme.xml")
	if err != nil {
		t.Fatalf("error reading theme file")
	}
	dec := xml.NewDecoder(f)
	ct := common.NewTheme()
	if err := dec.Decode(ct.X()); err != nil {
		t.Errorf("error decoding theme: %s", err)
	}

	got := &bytes.Buffer{}
	enc := xml.NewEncoder(zippkg.SelfClosingWriter{W: got})
	if err := enc.Encode(ct.X()); err != nil {
		t.Errorf("error encoding theme: %s", err)
	}

	testhelper.CompareGoldenXML(t, "theme.xml", got.Bytes())
}
