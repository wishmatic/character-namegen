package mcp

import (
	"context"
	"fmt"
	"strings"

	"github.com/google/jsonschema-go/jsonschema"
	"github.com/modelcontextprotocol/go-sdk/mcp"
	"github.com/wishmatic/character-namegen/gender"
	"github.com/wishmatic/character-namegen/name"
	"go.uber.org/zap"
)

var genders = []string{
	string(gender.Male),
	string(gender.Female),
	string(gender.NonBinary),
}

func cultures() []string {
	out := make([]string, len(name.SupportedCultures))
	for i, c := range name.SupportedCultures {
		out[i] = string(c)
	}

	return out
}

type generateNameInput struct {
	Gender  string `json:"gender,omitempty" jsonschema:"character gender; leave empty to pick randomly"`
	Culture string `json:"culture,omitempty" jsonschema:"naming culture; leave empty to pick randomly. See the tool description for allowed values"`
}

type generatedName struct {
	GivenName string `json:"given_name" jsonschema:"the generated given name"`
	Surname   string `json:"surname" jsonschema:"the generated surname"`
	Gender    string `json:"gender" jsonschema:"the gender the name was generated for"`
}

type generateNameOutput struct {
	Name generatedName `json:"name"`
}

func registerGenerateName(srv *mcp.Server, log *zap.Logger) {
	mcp.AddTool(srv, &mcp.Tool{
		Name:        "generate_name",
		Description: generateNameDescription(),
		InputSchema: generateNameSchema(),
	}, func(
		ctx context.Context,
		_ *mcp.CallToolRequest,
		in generateNameInput,
	) (*mcp.CallToolResult, generateNameOutput, error) {
		log.Debug("tool called",
			zap.String("tool", "generate_name"),
			zap.String("gender", in.Gender),
			zap.String("culture", in.Culture),
		)

		g, culture, err := resolveGenderAndCulture(in.Gender, in.Culture)
		if err != nil {
			return nil, generateNameOutput{}, fmt.Errorf("generate_name: %w", err)
		}

		n := name.Generate(g, culture)

		return &mcp.CallToolResult{}, generateNameOutput{Name: generatedName{
			GivenName: n.GivenName,
			Surname:   n.Surname,
			Gender:    string(n.Gender),
		}}, nil
	})
}

func generateNameDescription() string {
	return "Generate a random fantasy/RPG name. Both gender and culture are optional: " +
		"leave them empty to pick randomly. Acceptable genders: " +
		strings.Join(genders, ", ") + ". Acceptable cultures: " +
		strings.Join(cultures(), ", ") + "."
}

func resolveGenderAndCulture(genderStr, cultureStr string) (gender.Gender, *name.Culture, error) {
	g, err := resolveGender(genderStr)
	if err != nil {
		return "", nil, err
	}

	culture, err := resolveCulture(cultureStr)
	if err != nil {
		return "", nil, err
	}

	return g, culture, nil
}

func resolveGender(s string) (gender.Gender, error) {
	if s == "" {
		return "", nil
	}

	for _, g := range genders {
		if strings.EqualFold(g, s) {
			return gender.Gender(g), nil
		}
	}

	return "", fmt.Errorf(
		"unknown gender %q; acceptable values: %s (or empty for random)",
		s, strings.Join(genders, ", "),
	)
}

func resolveCulture(s string) (*name.Culture, error) {
	if s == "" {
		return nil, nil
	}

	for _, c := range name.SupportedCultures {
		if strings.EqualFold(string(c), s) {
			culture := c

			return &culture, nil
		}
	}

	return nil, fmt.Errorf(
		"unknown culture %q; acceptable values: %s (or empty for random)",
		s, strings.Join(cultures(), ", "),
	)
}

func generateNameSchema() *jsonschema.Schema {
	s, err := jsonschema.For[generateNameInput](nil)
	if err != nil {
		panic(fmt.Sprintf("generate_name: infer input schema: %v", err))
	}

	s.Properties["gender"].Enum = append([]any{""}, toEnum(genders)...)
	setDefault(s.Properties, "gender", "")

	s.Properties["culture"].Enum = append([]any{""}, toEnum(cultures())...)
	setDefault(s.Properties, "culture", "")

	return s
}
