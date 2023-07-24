package office_test

import (
	"testing"

	"go.devnw.com/ooxml"
)

func TestFloat32(t *testing.T) {
	exp := float32(1.234)
	got := office.Float32(exp)
	if *got != exp {
		t.Errorf("expected %f, got %f", exp, *got)
	}
}

func TestFloat64(t *testing.T) {
	exp := 1.234
	got := office.Float64(exp)
	if *got != exp {
		t.Errorf("expected %f, got %f", exp, *got)
	}
}

func TestUint64(t *testing.T) {
	exp := uint64(123)
	got := office.Uint64(exp)
	if *got != exp {
		t.Errorf("expected %d, got %d", exp, *got)
	}
}

func TestUint32(t *testing.T) {
	exp := uint32(123)
	got := office.Uint32(exp)
	if *got != exp {
		t.Errorf("expected %d, got %d", exp, *got)
	}
}

func TestInt64(t *testing.T) {
	exp := int64(123)
	got := office.Int64(exp)
	if *got != exp {
		t.Errorf("expected %d, got %d", exp, *got)
	}
}

func TestInt32(t *testing.T) {
	exp := int32(123)
	got := office.Int32(exp)
	if *got != exp {
		t.Errorf("expected %d, got %d", exp, *got)
	}
}

func TestInt8(t *testing.T) {
	exp := int8(123)
	got := office.Int8(exp)
	if *got != exp {
		t.Errorf("expected %d, got %d", exp, *got)
	}
}

func TestBool(t *testing.T) {
	exp := bool(true)
	got := office.Bool(exp)
	if *got != exp {
		t.Errorf("expected %v, got %v", exp, *got)
	}
}

func TestString(t *testing.T) {
	exp := "foo"
	got := office.String(exp)
	if *got != exp {
		t.Errorf("expected %s, got %s", exp, *got)
	}
}

func TestUint8(t *testing.T) {
	exp := uint8(123)
	got := office.Uint8(exp)
	if *got != exp {
		t.Errorf("expected %d, got %d", exp, *got)
	}
}

func TestUint16(t *testing.T) {
	exp := uint16(123)
	got := office.Uint16(exp)
	if *got != exp {
		t.Errorf("expected %d, got %d", exp, *got)
	}
}

func TestStringf(t *testing.T) {
	exp := "foobar123"
	got := office.Stringf("foo%s%d", "bar", 123)
	if *got != exp {
		t.Errorf("expected %s, got %s", exp, *got)
	}
}
