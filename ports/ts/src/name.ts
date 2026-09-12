import { Female, Male, type Gender } from "./gender.js";
import { generateIRLName } from "./irl.js";
import { generateRandomPhonemeName } from "./phonemes.js";
import { pick } from "./random.js";
import {
  PHONEME_CULTURES,
  SUPPORTED_CULTURES,
  type Culture,
  type PhonemeCulture,
} from "./culture.js";
import type { Name } from "./types.js";

const PHONEME_CULTURE_SET: ReadonlySet<string> = new Set(PHONEME_CULTURES);

function isPhonemeCulture(culture: Culture): culture is PhonemeCulture {
  return PHONEME_CULTURE_SET.has(culture);
}

export function generate(gender: Gender | "" = "", culture: Culture | null = null): Name {
  const resolvedCulture: Culture = culture ?? pick(SUPPORTED_CULTURES);
  const resolvedGender: Gender = gender === "" ? pick([Female, Male]) : gender;

  if (isPhonemeCulture(resolvedCulture)) {
    return generateRandomPhonemeName(resolvedGender, resolvedCulture);
  }

  return generateIRLName(resolvedCulture, resolvedGender);
}
