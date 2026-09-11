# `character-namegen` NPM Package

RPG name generator with a decently good IRL name generator. This is the TypeScript package — the Go library, MCP
server, and data-generation scripts live in the
[main repository](https://github.com/wishmatic/character-namegen).

## Install

```sh
npm install character-namegen
# ...or
pnpm add character-namegen
# ...or
yarn add character-namegen
# ...or
bun add character-namegen
```

## Usage

```ts
import { generate, Male } from "character-namegen";

const generated = generate(Male);
// => { givenName, surname, gender }
```

`generate(gender?, culture?)`: an empty gender or a `null` culture is resolved at random.

### Types

- `Gender` (`"male" | "female" | "non-binary"`), with `Male`, `Female`, `NonBinary` constants
- `Name{ givenName, surname: string; gender: Gender }`
- `Culture`, `IRLCulture`, `PhonemeCulture`; `SUPPORTED_CULTURES`, `IRL_CULTURES`, `PHONEME_CULTURES`

## License

Apache-2.0. The name lists are derived from public sources and transformed and cleaned by us; please see the
[credits](https://github.com/wishmatic/character-namegen#credits) and
[LICENSE](https://github.com/wishmatic/character-namegen/blob/main/LICENSE).
