package cardanometadata

import (
	"strings"
	"testing"
)

func TestStringOrChunks_Short(t *testing.T) {
	v := StringOrChunks("hello", 64)
	s, ok := v.(string)
	if !ok {
		t.Fatalf("expected string, got %T", v)
	}
	if s != "hello" {
		t.Errorf("expected %q, got %q", "hello", s)
	}
}

func TestStringOrChunks_Long(t *testing.T) {
	long := strings.Repeat("a", 100)
	v := StringOrChunks(long, 64)
	chunks, ok := v.([]string)
	if !ok {
		t.Fatalf("expected []string, got %T", v)
	}
	if len(chunks) != 2 {
		t.Errorf("expected 2 chunks, got %d", len(chunks))
	}
	if len(chunks[0]) != 64 {
		t.Errorf("expected first chunk of 64, got %d", len(chunks[0]))
	}
	if len(chunks[1]) != 36 {
		t.Errorf("expected second chunk of 36, got %d", len(chunks[1]))
	}
}

func TestStringOrChunks_Exact64(t *testing.T) {
	s := strings.Repeat("b", 64)
	v := StringOrChunks(s, 64)
	if _, ok := v.(string); !ok {
		t.Errorf("expected string for exactly 64 bytes, got %T", v)
	}
}

func TestStringOrChunks_ZeroMaxLen(t *testing.T) {
	s := strings.Repeat("c", 10)
	v := StringOrChunks(s, 0)
	if _, ok := v.(string); !ok {
		t.Errorf("zero maxLen should default to 64; expected string, got %T", v)
	}
}

func TestValidationError(t *testing.T) {
	e := &ValidationError{Field: "name", Message: "empty"}
	if e.Error() == "" {
		t.Error("expected non-empty error string")
	}
}

func TestSchemaError(t *testing.T) {
	e := &SchemaError{Expected: "string", Got: "int"}
	if e.Error() == "" {
		t.Error("expected non-empty error string")
	}
}
