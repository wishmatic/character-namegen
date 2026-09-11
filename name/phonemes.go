package name

import (
	"strings"
	"unicode"

	"github.com/wishmatic/character-namegen/gender"
	"github.com/wishmatic/character-namegen/utils"
)

// phonemeVowels is the set of vowel letters used by the phoneme sets.
var phonemeVowels = map[rune]bool{
	'a': true, 'e': true, 'i': true, 'o': true, 'u': true,
}

// leadingConsonants returns the number of consecutive consonant runes at the start of s.
func leadingConsonants(s string) int {
	n := 0

	for _, r := range s {
		if phonemeVowels[unicode.ToLower(r)] {
			break
		}

		n++
	}

	return n
}

// trailingConsonants returns the number of consecutive consonant runes at the end of s.
func trailingConsonants(s string) int {
	rs := []rune(s)
	n := 0

	for i := len(rs) - 1; i >= 0; i-- {
		if phonemeVowels[unicode.ToLower(rs[i])] {
			break
		}

		n++
	}

	return n
}

// leadingVowels returns the number of consecutive vowel runes at the start of s.
func leadingVowels(s string) int {
	n := 0

	for _, r := range s {
		if !phonemeVowels[unicode.ToLower(r)] {
			break
		}

		n++
	}

	return n
}

// trailingVowels returns the number of consecutive vowel runes at the end of s.
func trailingVowels(s string) int {
	rs := []rune(s)
	n := 0

	for i := len(rs) - 1; i >= 0; i-- {
		if !phonemeVowels[unicode.ToLower(rs[i])] {
			break
		}

		n++
	}

	return n
}

// hasAwkwardConsonantCluster reports whether text contains three or more consecutive consonants anywhere after the
// first vowel. A leading cluster such as "thr" in "thrash" is considered acceptable.
func hasAwkwardConsonantCluster(text string) bool {
	streak := 0
	seenVowel := false

	for _, r := range text {
		if phonemeVowels[unicode.ToLower(r)] {
			streak = 0
			seenVowel = true

			continue
		}

		streak++
		if streak >= 3 && seenVowel {
			return true
		}
	}

	return false
}

// hasAwkwardVowelCluster reports whether text contains three or more consecutive vowels.
func hasAwkwardVowelCluster(text string) bool {
	streak := 0

	for _, r := range text {
		if phonemeVowels[unicode.ToLower(r)] {
			streak++
			if streak >= 3 {
				return true
			}
		} else {
			streak = 0
		}
	}

	return false
}

// pickComponent picks a random candidate that, when appended to built, keeps both consonant and vowel clusters below
// three. The word-initial onset may instead carry a leading consonant cluster such as "thr".
func pickComponent(candidates []string, built string, allowLeadingCluster bool) string {
	maxLeadingCons := 2 - trailingConsonants(built)
	if allowLeadingCluster {
		maxLeadingCons = 3
	}
	if maxLeadingCons < 0 {
		maxLeadingCons = 0
	}

	maxLeadingVow := 2 - trailingVowels(built)
	if maxLeadingVow < 0 {
		maxLeadingVow = 0
	}

	var allowed []string
	for _, c := range candidates {
		if leadingConsonants(c) <= maxLeadingCons && leadingVowels(c) <= maxLeadingVow {
			allowed = append(allowed, c)
		}
	}
	return utils.Pick(allowed)
}

// withoutEmpty returns a copy of s with empty strings removed.
func withoutEmpty(s []string) []string {
	var out []string
	for _, x := range s {
		if x != "" {
			out = append(out, x)
		}
	}
	return out
}

// pickWithMaxTrailing picks a random candidate whose trailing consonant count is at most maxTrailing.
func pickWithMaxTrailing(candidates []string, maxTrailing int) string {
	var allowed []string
	for _, c := range candidates {
		if trailingConsonants(c) <= maxTrailing {
			allowed = append(allowed, c)
		}
	}
	return utils.Pick(allowed)
}

type phonemeSet struct {
	onsets         []string
	nuclei         []string
	medialCodas    []string
	maleEndings    []string
	femaleEndings  []string
	neutralEndings []string
}

var phonemeSets = map[PhonemeCulture]phonemeSet{
	"phonemes": {
		onsets: []string{
			"", "b", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p",
			"r", "s", "sh", "t", "th", "v", "w", "br", "ch", "cr", "dr", "fl",
			"fr", "gr", "pr", "st", "tr",
		},
		nuclei:      []string{"a", "e", "i", "o", "u", "ai", "ei", "ou"},
		medialCodas: []string{"", "", "", "", "l", "n", "r", "s"},
		maleEndings: []string{
			"d", "k", "l", "m", "n", "nd", "r", "rd", "rn", "s", "st", "t",
			"th", "x",
		},
		femaleEndings: []string{
			"a", "e", "i", "ia", "ra", "na", "la", "ne", "le", "sa", "ya",
		},
		neutralEndings: []string{"l", "n", "r", "a", "e", "i", "o"},
	},
	"elvish": {
		onsets: []string{
			"", "l", "n", "r", "s", "th", "f", "v", "m", "gl", "w", "y", "sh",
			"br", "dr",
		},
		nuclei:      []string{"a", "e", "i", "o", "ae", "ie", "ia"},
		medialCodas: []string{"", "", "l", "n", "r", "nd", "th"},
		maleEndings: []string{"l", "n", "r", "s", "nd", "las", "ren", "mir", "ion", "or"},
		femaleEndings: []string{
			"a", "e", "ia", "iel", "wen", "riel", "na", "ya",
		},
		neutralEndings: []string{"el", "en", "il", "al", "ar", "an", "iel", "ion", "s", "n"},
	},
	"fae": {
		onsets: []string{
			"", "f", "l", "n", "p", "t", "w", "y", "fl", "tw", "sp", "br",
		},
		nuclei:         []string{"a", "e", "i", "o", "u", "ee", "ie"},
		medialCodas:    []string{"", "", "", "l", "n"},
		maleEndings:    []string{"n", "l", "x", "p", "s", "ck", "ll"},
		femaleEndings:  []string{"a", "e", "i", "li", "ny"},
		neutralEndings: []string{"l", "n", "a", "i", "e"},
	},
	"khuzdul": {
		onsets: []string{
			"", "b", "d", "g", "k", "t", "z", "th", "kh", "gr", "dr", "br",
			"kr", "dz", "gl", "gm", "n",
		},
		nuclei:      []string{"a", "u", "o", "i", "e", "ur", "ul", "az", "un"},
		medialCodas: []string{"", "r", "l", "z", "n", "m", "rk", "lk"},
		maleEndings: []string{
			"k", "r", "d", "m", "n", "rk", "rd", "grim", "dur", "rik", "mund",
		},
		femaleEndings:  []string{"a", "i", "ra", "da", "ka", "ri", "di", "na"},
		neutralEndings: []string{"r", "k", "n", "m", "ul", "az"},
	},
	"orkind": {
		onsets: []string{
			"", "g", "gr", "kr", "k", "z", "zh", "b", "d", "dr", "n", "r",
			"sk", "sh", "t", "thr", "m", "gn",
		},
		nuclei:      []string{"a", "u", "o", "ug", "ur", "ag", "og", "uk", "ash"},
		medialCodas: []string{"", "g", "k", "r", "z", "rg", "rk"},
		maleEndings: []string{
			"g", "k", "rg", "zg", "th", "rk", "gul", "bur", "nash", "rok",
		},
		femaleEndings:  []string{"a", "ga", "ra", "sha", "ka", "za", "gra"},
		neutralEndings: []string{"g", "k", "r", "z", "uk", "og"},
	},
}

var PhonemeSets = phonemeSets

// generatePhonemes builds a name of the given syllable length and ending, constraining each boundary so no internal
// consonant cluster exceeds two.
func generatePhonemes(length int, ending string, set phonemeSet) string {
	var b strings.Builder

	for i := 0; i < length; i++ {
		b.WriteString(pickComponent(set.onsets, b.String(), i == 0))

		if i == length-1 {
			b.WriteString(pickWithMaxTrailing(set.nuclei, 1))
		} else {
			b.WriteString(utils.Pick(set.nuclei))
		}

		if i < length-1 {
			codas := set.medialCodas
			if trailingVowels(b.String()) > 0 {
				codas = withoutEmpty(codas)
			}
			b.WriteString(pickComponent(codas, b.String(), false))
		}
	}

	var endings []string
	switch ending {
	case "female":
		endings = set.femaleEndings
	case "male":
		endings = set.maleEndings
	default:
		endings = set.neutralEndings
	}

	b.WriteString(pickComponent(endings, b.String(), false))

	name := b.String()

	return strings.ToUpper(name[:1]) + name[1:]
}

func generateRandomPhonemeName(g gender.Gender, culture PhonemeCulture) Name {
	set := phonemeSets[culture]

	givenLength := utils.Pick([]int{1, 1, 2, 2, 2})
	surnameLength := utils.Pick([]int{1, 1, 2, 2, 2})

	givenEnding := "neutral"
	if g == gender.Female {
		givenEnding = "female"
	} else if g == gender.Male {
		givenEnding = "male"
	}

	return Name{
		GivenName: generatePhonemes(givenLength, givenEnding, set),
		Surname:   generatePhonemes(surnameLength, "neutral", set),
		Gender:    g,
	}
}
