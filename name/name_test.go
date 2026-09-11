package name

import (
	"testing"

	"github.com/wishmatic/character-namegen/gender"
)

func TestGenerate(t *testing.T) {
	for _, g := range []gender.Gender{"", gender.Male, gender.Female, gender.NonBinary} {
		n := Generate(g, nil)
		if n.GivenName == "" {
			t.Errorf("Generate(%q) produced empty given name", g)
		}

		if n.Surname == "" {
			t.Errorf("Generate(%q) produced empty surname", g)
		}

		if g != "" && n.Gender != g {
			t.Errorf("Generate(%q) produced gender %q", g, n.Gender)
		}
	}
}

func TestGenerate_PhonemeCultures(t *testing.T) {
	for _, c := range PhonemeCultures {
		c := c
		culture := Culture(c)

		for range 20 {
			n := Generate("", &culture)
			if n.GivenName == "" || n.Surname == "" {
				t.Fatalf("phoneme culture %q produced empty name: %+v", c, n)
			}
		}
	}
}

func TestGenerate_IRLCultures(t *testing.T) {
	for _, c := range IRLCultures {
		c := c
		culture := Culture(c)

		for range 20 {
			n := Generate("", &culture)
			if n.GivenName == "" || n.Surname == "" {
				t.Fatalf("IRL culture %q produced empty name: %+v", c, n)
			}
		}
	}
}

func TestGeneratePhonemes_NoAwkwardClusters(t *testing.T) {
	for _, c := range PhonemeCultures {
		set := phonemeSets[c]
		for range 500 {
			n := generatePhonemes(3, "neutral", set)
			if hasAwkwardConsonantCluster(n) {
				t.Fatalf("culture %q produced awkward consonant cluster: %q", c, n)
			}
			if hasAwkwardVowelCluster(n) {
				t.Fatalf("culture %q produced awkward vowel cluster: %q", c, n)
			}
		}
	}
}

func TestMatchesCzechSlovakGender(t *testing.T) {
	cases := []struct {
		surname string
		male    bool
		female  bool
	}{
		{"Novák", true, false},
		{"Nováková", false, true},
		{"Černý", true, false},
		{"Černá", false, true},
		{"Svoboda", true, false},
		{"Krejčí", true, true}, // invariant
	}
	for _, c := range cases {
		if got := matchesCzechSlovakGender(c.surname, gender.Male); got != c.male {
			t.Errorf("matchesCzechSlovakGender(%q, Male) = %v, want %v", c.surname, got, c.male)
		}
		if got := matchesCzechSlovakGender(c.surname, gender.Female); got != c.female {
			t.Errorf("matchesCzechSlovakGender(%q, Female) = %v, want %v", c.surname, got, c.female)
		}
	}
}
