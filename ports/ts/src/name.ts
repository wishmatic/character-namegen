import { Female, Male, type Gender } from "./gender";
import { generateIRLName } from "./irl";
import { generateRandomPhonemeName } from "./phonemes";
import { pick } from "./random";
import {
  PHONEME_CULTURES,
  SUPPORTED_CULTURES,
  type Culture,
  type PhonemeCulture,
} from "./culture";
import type { Name } from "./types";

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
