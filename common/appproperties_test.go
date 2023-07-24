package common_test

import (
	"fmt"
	"testing"

	"go.devnw.com/ooxml/common"
)

func TestNewAppDefaultProperties(t *testing.T) {
	ap := common.NewAppProperties()
	if ap.X() == nil {
		t.Errorf("expected non-nil internal element")
	}
	if got := ap.Application(); got != "go.devnw.com/ooxml" {
		t.Errorf("unexpected application: %s", got)
	}

	ap.X().Application = nil
	if got := ap.Application(); got != "" {
		t.Errorf("unexpected application: %s", got)
	}

	var major, minor, patch int64
	fmt.Sscanf(common.Version, "%d.%d.%d", &major, &minor, &patch)
	fv := float64(major) + float64(minor)/10000.0
	if got := ap.ApplicationVersion(); got != fmt.Sprintf("%07.4f", fv) {
		t.Errorf("unexpected application version: %s", got)
	}
	ap.X().AppVersion = nil
	if got := ap.ApplicationVersion(); got != "" {
		t.Errorf("unexpected application version: %s", got)
	}
}

func TestAppPropertiesSetApplication(t *testing.T) {
	ap := common.NewAppProperties()
	if ap.X() == nil {
		t.Errorf("expected non-nil internal element")
	}
	ap.SetApplication("foo")
	if got := ap.Application(); got != "foo" {
		t.Errorf("unexpected application: %s", got)
	}
}

func TestAppPropertiesSetApplicationVersion(t *testing.T) {
	ap := common.NewAppProperties()
	if ap.X() == nil {
		t.Errorf("expected non-nil internal element")
	}
	ap.SetApplicationVersion("foo")
	if got := ap.ApplicationVersion(); got != "foo" {
		t.Errorf("unexpected application: %s", got)
	}
}

func TestAppPropertiesSetCompany(t *testing.T) {
	ap := common.NewAppProperties()
	if ap.X() == nil {
		t.Errorf("expected non-nil internal element")
	}
	ap.SetCompany("foo")
	if got := ap.Company(); got != "foo" {
		t.Errorf("unexpected company: %s", got)
	}
}
