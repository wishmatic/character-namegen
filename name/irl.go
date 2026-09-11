package name

import (
	_ "embed"
	"strings"

	"github.com/wishmatic/character-namegen/gender"
	"github.com/wishmatic/character-namegen/utils"
)

// lines splits raw newline-separated text into trimmed, non-empty lines.
func lines(raw string) []string {
	var out []string
	for _, line := range strings.Split(raw, "\n") {
		line = strings.TrimSpace(line)
		if line != "" {
			out = append(out, line)
		}
	}

	return out
}

type irlNameData struct {
	givenMale   []string
	givenFemale []string
	surnames    []string
}

var irlCultureToName = map[IRLCulture]irlNameData{
	"Afghan": {
		givenMale:   lines(afghanGivenMaleRaw),
		givenFemale: lines(afghanGivenFemaleRaw),
		surnames:    lines(afghanSurnamesRaw),
	},
	"African (Central)": {
		givenMale:   lines(africanCentralGivenMaleRaw),
		givenFemale: lines(africanCentralGivenFemaleRaw),
		surnames:    lines(africanCentralSurnamesRaw),
	},
	"African (Southern)": {
		givenMale:   lines(africanSouthernGivenMaleRaw),
		givenFemale: lines(africanSouthernGivenFemaleRaw),
		surnames:    lines(africanSouthernSurnamesRaw),
	},
	"African (West)": {
		givenMale:   lines(africanWestGivenMaleRaw),
		givenFemale: lines(africanWestGivenFemaleRaw),
		surnames:    lines(africanWestSurnamesRaw),
	},
	"Anglophone": {
		givenMale:   lines(anglophoneGivenMaleRaw),
		givenFemale: lines(anglophoneGivenFemaleRaw),
		surnames:    lines(anglophoneSurnamesRaw),
	},
	"Arabic (Gulf)": {
		givenMale:   lines(gulfGivenMaleRaw),
		givenFemale: lines(gulfGivenFemaleRaw),
		surnames:    lines(gulfSurnamesRaw),
	},
	"Arabic (Levantine)": {
		givenMale:   lines(levantineGivenMaleRaw),
		givenFemale: lines(levantineGivenFemaleRaw),
		surnames:    lines(levantineSurnamesRaw),
	},
	"Arabic (Maghrebi)": {
		givenMale:   lines(maghrebiGivenMaleRaw),
		givenFemale: lines(maghrebiGivenFemaleRaw),
		surnames:    lines(maghrebiSurnamesRaw),
	},
	"Asian (South)": {
		givenMale:   lines(asianSouthGivenMaleRaw),
		givenFemale: lines(asianSouthGivenFemaleRaw),
		surnames:    lines(asianSouthSurnamesRaw),
	},
	"Asian (Southeast)": {
		givenMale:   lines(asianSoutheastGivenMaleRaw),
		givenFemale: lines(asianSoutheastGivenFemaleRaw),
		surnames:    lines(asianSoutheastSurnamesRaw),
	},
	"Chinese": {
		givenMale:   lines(chineseGivenMaleRaw),
		givenFemale: lines(chineseGivenFemaleRaw),
		surnames:    lines(chineseSurnamesRaw),
	},
	"Ethiopian": {
		givenMale:   lines(ethiopianGivenMaleRaw),
		givenFemale: lines(ethiopianGivenFemaleRaw),
		surnames:    lines(ethiopianSurnamesRaw),
	},
	"European (Central)": {
		givenMale:   lines(europeanCentralGivenMaleRaw),
		givenFemale: lines(europeanCentralGivenFemaleRaw),
		surnames:    lines(europeanCentralSurnamesRaw),
	},
	"European (Southern)": {
		givenMale:   lines(europeanSouthernGivenMaleRaw),
		givenFemale: lines(europeanSouthernGivenFemaleRaw),
		surnames:    lines(europeanSouthernSurnamesRaw),
	},
	"Fijian": {
		givenMale:   lines(fijianGivenMaleRaw),
		givenFemale: lines(fijianGivenFemaleRaw),
		surnames:    lines(fijianSurnamesRaw),
	},
	"French": {
		givenMale:   lines(frenchGivenMaleRaw),
		givenFemale: lines(frenchGivenFemaleRaw),
		surnames:    lines(frenchSurnamesRaw),
	},
	"Georgian": {
		givenMale:   lines(georgianGivenMaleRaw),
		givenFemale: lines(georgianGivenFemaleRaw),
		surnames:    lines(georgianSurnamesRaw),
	},
	"Germanic": {
		givenMale:   lines(germanicGivenMaleRaw),
		givenFemale: lines(germanicGivenFemaleRaw),
		surnames:    lines(germanicSurnamesRaw),
	},
	"Greek": {
		givenMale:   lines(greekGivenMaleRaw),
		givenFemale: lines(greekGivenFemaleRaw),
		surnames:    lines(greekSurnamesRaw),
	},
	"Iberoamerican": {
		givenMale:   lines(americanIberoGivenMaleRaw),
		givenFemale: lines(americanIberoGivenFemaleRaw),
		surnames:    lines(americanIberoSurnamesRaw),
	},
	"Iranian": {
		givenMale:   lines(iranianGivenMaleRaw),
		givenFemale: lines(iranianGivenFemaleRaw),
		surnames:    lines(iranianSurnamesRaw),
	},
	"Irish": {
		givenMale:   lines(irishGivenMaleRaw),
		givenFemale: lines(irishGivenFemaleRaw),
		surnames:    lines(irishSurnamesRaw),
	},
	"Japanese": {
		givenMale:   lines(japaneseGivenMaleRaw),
		givenFemale: lines(japaneseGivenFemaleRaw),
		surnames:    lines(japaneseSurnamesRaw),
	},
	"Korean": {
		givenMale:   lines(koreanGivenMaleRaw),
		givenFemale: lines(koreanGivenFemaleRaw),
		surnames:    lines(koreanSurnamesRaw),
	},
	"Nordic": {
		givenMale:   lines(nordicGivenMaleRaw),
		givenFemale: lines(nordicGivenFemaleRaw),
		surnames:    lines(nordicSurnamesRaw),
	},
	"Russian": {
		givenMale:   lines(russianGivenMaleRaw),
		givenFemale: lines(russianGivenFemaleRaw),
		surnames:    lines(russianSurnamesRaw),
	},
	"Slavic (South)": {
		givenMale:   lines(slavicSouthGivenMaleRaw),
		givenFemale: lines(slavicSouthGivenFemaleRaw),
		surnames:    lines(slavicSouthSurnamesRaw),
	},
	"Turkic": {
		givenMale:   lines(turkicGivenMaleRaw),
		givenFemale: lines(turkicGivenFemaleRaw),
		surnames:    lines(turkicSurnamesRaw),
	},
}

// generateIRLName generates a real-world Name for the given culture and gender.
func generateIRLName(culture *Culture, g gender.Gender) Name {
	nameData := determineNameData(culture)

	isMaleOrFemale := g == gender.Male || g == gender.Female

	var givenPool []string
	switch {
	case !isMaleOrFemale:
		givenPool = append(append([]string{}, nameData.givenMale...), nameData.givenFemale...)
	case g == gender.Female:
		givenPool = nameData.givenFemale
	default:
		givenPool = nameData.givenMale
	}

	givenName := utils.Pick(givenPool)
	surname := utils.Pick(nameData.surnames)

	if culture != nil {
		switch IRLCulture(*culture) {
		case "Russian":
			surname = genderizeRussianSurname(surname, g)
		case "European (Central)":
			surname = pickCzechSlovakSurname(nameData.surnames, g)
		}
	}

	return Name{
		GivenName: givenName,
		Surname:   surname,
		Gender:    g,
	}
}

func determineNameData(culture *Culture) irlNameData {
	if culture != nil {
		return irlCultureToName[IRLCulture(*culture)]
	}

	// Allow culture mixing in the pool when culture is explicitly nil.

	var pooled irlNameData
	for _, datum := range irlCultureToName {
		pooled.givenMale = append(pooled.givenMale, datum.givenMale...)
		pooled.givenFemale = append(pooled.givenFemale, datum.givenFemale...)
		pooled.surnames = append(pooled.surnames, datum.surnames...)
	}
	return pooled
}

func genderizeRussianSurname(surname string, g gender.Gender) string {
	isFemale := g == gender.Female
	isMale := g == gender.Male

	surnameLower := strings.ToLower(surname)

	if isFemale {
		switch {
		case strings.HasSuffix(surnameLower, "skiy") || strings.HasSuffix(surnameLower, "skij"):
			return surname[:len(surname)-4] + "skaya"
		case strings.HasSuffix(surnameLower, "sky"):
			return surname[:len(surname)-3] + "skaya"
		case strings.HasSuffix(surnameLower, "oy"):
			return surname[:len(surname)-2] + "aya"
		case strings.HasSuffix(surnameLower, "ev") ||
			strings.HasSuffix(surnameLower, "ov") ||
			strings.HasSuffix(surnameLower, "in"):

			return surname + "a"
		default:
			return surname
		}
	}

	if isMale {
		switch {
		case strings.HasSuffix(surnameLower, "skaya"):
			return surname[:len(surname)-5] + "sky"
		case strings.HasSuffix(surnameLower, "aya"):
			return surname[:len(surname)-3] + "oy"
		case strings.HasSuffix(surnameLower, "ova") ||
			strings.HasSuffix(surnameLower, "eva") ||
			strings.HasSuffix(surnameLower, "ina"):
			return surname[:len(surname)-1]
		default:
			return surname
		}
	}

	return surname
}

// pickCzechSlovakSurname returns a surname from the mixed male/female pool that matches the given gender.
// Czech/Slovak surnames are stored in both forms, so picking (rather than converting) avoids the fleeting-vowel
// problem (e.g. Havlíček/Havlíčková, Svoboda/Svobodová).
func pickCzechSlovakSurname(surnames []string, g gender.Gender) string {
	for range 20 {
		s := utils.Pick(surnames)
		if matchesCzechSlovakGender(s, g) {
			return s
		}
	}
	return utils.Pick(surnames)
}

// matchesCzechSlovakGender reports whether a Czech/Slovak surname form is valid for the given gender. Female forms
// end in -ová (or unaccented -ova) or -á; soft adjectival -í is invariant and valid for both genders.
func matchesCzechSlovakGender(surname string, g gender.Gender) bool {
	lower := strings.ToLower(surname)
	isFemale := strings.HasSuffix(lower, "ová") ||
		strings.HasSuffix(lower, "ova") ||
		strings.HasSuffix(lower, "á")

	switch g {
	case gender.Male:
		return !isFemale
	case gender.Female:
		return isFemale || strings.HasSuffix(lower, "í")
	default:
		return true
	}
}
