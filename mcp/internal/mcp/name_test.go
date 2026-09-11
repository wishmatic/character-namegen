package mcp

import (
	"testing"

	"github.com/wishmatic/character-namegen/gender"
	"github.com/wishmatic/character-namegen/name"
	"go.uber.org/zap"
)

func TestNewRegistersTools(t *testing.T) {
	srv, err := New(zap.NewNop())
	if err != nil {
		t.Fatalf("New() unexpected error: %v", err)
	}

	if srv == nil {
		t.Fatal("New() returned nil server")
	}
}

func TestResolveGender(t *testing.T) {
	tests := []struct {
		in      string
		want    gender.Gender
		wantErr bool
	}{
		{in: "", want: ""},
		{in: "male", want: gender.Male},
		{in: "FEMALE", want: gender.Female},
		{in: "non-binary", want: gender.NonBinary},
		{in: "elf", wantErr: true},
	}

	for _, tt := range tests {
		t.Run(tt.in, func(t *testing.T) {
			got, err := resolveGender(tt.in)
			if tt.wantErr {
				if err == nil {
					t.Fatalf("resolveGender(%q) expected error", tt.in)
				}

				return
			}

			if err != nil {
				t.Fatalf("resolveGender(%q) unexpected error: %v", tt.in, err)
			}

			if got != tt.want {
				t.Errorf("resolveGender(%q) = %q, want %q", tt.in, got, tt.want)
			}
		})
	}
}

func TestResolveCulture(t *testing.T) {
	culture, err := resolveCulture("nordic")
	if err != nil {
		t.Fatalf("resolveCulture(nordic) unexpected error: %v", err)
	}

	if culture == nil || *culture != name.Culture("Nordic") {
		t.Fatalf("resolveCulture(nordic) = %v, want Nordic", culture)
	}

	if _, err := resolveCulture("atlantean"); err == nil {
		t.Error("resolveCulture(atlantean) expected error")
	}

	empty, err := resolveCulture("")
	if err != nil {
		t.Fatalf("resolveCulture(\"\") unexpected error: %v", err)
	}

	if empty != nil {
		t.Errorf("resolveCulture(\"\") = %v, want nil", empty)
	}
}

func TestGenerateNameSchemaHasEnums(t *testing.T) {
	s := generateNameSchema()

	for _, prop := range []string{"gender", "culture"} {
		field, ok := s.Properties[prop]
		if !ok {
			t.Fatalf("schema missing property %q", prop)
		}

		if len(field.Enum) == 0 {
			t.Errorf("schema property %q has no enum values", prop)
		}
	}
}
