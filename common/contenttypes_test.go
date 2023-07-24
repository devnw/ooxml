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

func TestContentTypesUnmarshal(t *testing.T) {
	f, err := os.Open("testdata/contenttypes.xml")
	if err != nil {
		t.Fatalf("error reading content types file")
	}
	dec := xml.NewDecoder(f)
	ct := common.NewContentTypes()
	if err := dec.Decode(ct.X()); err != nil {
		t.Errorf("error decoding content types: %s", err)
	}

	got := &bytes.Buffer{}
	enc := xml.NewEncoder(zippkg.SelfClosingWriter{W: got})
	if err := enc.Encode(ct.X()); err != nil {
		t.Errorf("error encoding content types: %s", err)
	}

	testhelper.CompareGoldenXML(t, "contenttypes.xml", got.Bytes())
}

func TestCopyOverride(t *testing.T) {
	ct := common.NewContentTypes()
	ct.AddOverride("/foo/bar.xml", "application/xml")

	lenBefore := len(ct.X().Override)

	ct.CopyOverride("foo/bar.xml", "foo/bar2.xml")

	if len(ct.X().Override) != (lenBefore + 1) {
		t.Errorf("expected override len %d, got %d", lenBefore+1, len(ct.X().Override))
	}

	copyIdx := len(ct.X().Override) - 1

	if ct.X().Override[copyIdx].PartNameAttr != "/foo/bar2.xml" {
		t.Errorf("expected \"/foo/bar2.xml\" PartNameAttr, go %s", ct.X().Override[copyIdx].PartNameAttr)
	}
}
