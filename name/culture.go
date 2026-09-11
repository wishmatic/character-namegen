package name

type Culture string

var SupportedCultures = func() []Culture {
	all := make([]Culture, 0, len(IRLCultures)+len(PhonemeCultures))

	for _, c := range IRLCultures {
		all = append(all, Culture(c))
	}

	for _, c := range PhonemeCultures {
		all = append(all, Culture(c))
	}

	return all
}()

type IRLCulture string

var IRLCultures = []IRLCulture{
	"Afghan",
	"African (Central)",
	"African (Southern)",
	"African (West)",
	"Anglophone",
	"Arabic (Gulf)",
	"Arabic (Levantine)",
	"Arabic (Maghrebi)",
	"Asian (South)",
	"Asian (Southeast)",
	"Chinese",
	"Ethiopian",
	"European (Central)",
	"European (Southern)",
	"Fijian",
	"French",
	"Georgian",
	"Germanic",
	"Greek",
	"Iberoamerican",
	"Iranian",
	"Irish",
	"Japanese",
	"Korean",
	"Nordic",
	"Russian",
	"Slavic (South)",
	"Turkic",
}

type PhonemeCulture string

var PhonemeCultures = []PhonemeCulture{
	"phonemes",
	"elvish",
	"khuzdul",
	"orkind",
	"fae",
}
