package name

import (
	"github.com/wishmatic/namegen/gender"
	"github.com/wishmatic/namegen/utils"
)

type Name struct {
	GivenName string
	Surname   string
	Gender    gender.Gender
}

func Generate(g gender.Gender, culture *Culture) Name {
	var resolvedCulture Culture
	if culture == nil {
		resolvedCulture = utils.Pick(SupportedCultures)
	} else {
		resolvedCulture = *culture
	}

	resolvedGender := g
	if resolvedGender == "" {
		resolvedGender = utils.Pick([]gender.Gender{gender.Female, gender.Male})
	}

	if _, ok := phonemeSets[PhonemeCulture(resolvedCulture)]; ok {
		return generateRandomPhonemeName(resolvedGender, PhonemeCulture(resolvedCulture))
	}

	return generateIRLName(&resolvedCulture, resolvedGender)
}
