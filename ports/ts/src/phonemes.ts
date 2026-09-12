import { Female, Male, type Gender } from "./gender.js";
import type { PhonemeCulture } from "./culture.js";
import { pick } from "./random.js";
import type { Name } from "./types.js";

interface PhonemeSet {
  onsets: string[];
  nuclei: string[];
  medialCodas: string[];
  maleEndings: string[];
  femaleEndings: string[];
  neutralEndings: string[];
}

const phonemeSets: Record<PhonemeCulture, PhonemeSet> = {
  phonemes: {
    onsets: [
      "", "b", "d", "f", "g", "h", "j", "k", "l", "m", "n", "p",
      "r", "s", "sh", "t", "th", "v", "w", "br", "ch", "cr", "dr", "fl",
      "fr", "gr", "pr", "st", "tr",
    ],
    nuclei: ["a", "e", "i", "o", "u", "ai", "ei", "ou"],
    medialCodas: ["", "", "", "", "l", "n", "r", "s"],
    maleEndings: [
      "d", "k", "l", "m", "n", "nd", "r", "rd", "rn", "s", "st", "t",
      "th", "x",
    ],
    femaleEndings: [
      "a", "e", "i", "ia", "ra", "na", "la", "ne", "le", "sa", "ya",
    ],
    neutralEndings: ["l", "n", "r", "a", "e", "i", "o"],
  },
  elvish: {
    onsets: [
      "", "l", "n", "r", "s", "th", "f", "v", "m", "gl", "w", "y", "sh",
      "br", "dr",
    ],
    nuclei: ["a", "e", "i", "o", "ae", "ie", "ia"],
    medialCodas: ["", "", "l", "n", "r", "nd", "th"],
    maleEndings: ["l", "n", "r", "s", "nd", "las", "ren", "mir", "ion", "or"],
    femaleEndings: ["a", "e", "ia", "iel", "wen", "riel", "na", "ya"],
    neutralEndings: ["el", "en", "il", "al", "ar", "an", "iel", "ion", "s", "n"],
  },
  fae: {
    onsets: ["", "f", "l", "n", "p", "t", "w", "y", "fl", "tw", "sp", "br"],
    nuclei: ["a", "e", "i", "o", "u", "ee", "ie"],
    medialCodas: ["", "", "", "l", "n"],
    maleEndings: ["n", "l", "x", "p", "s", "ck", "ll"],
    femaleEndings: ["a", "e", "i", "li", "ny"],
    neutralEndings: ["l", "n", "a", "i", "e"],
  },
  khuzdul: {
    onsets: [
      "", "b", "d", "g", "k", "t", "z", "th", "kh", "gr", "dr", "br",
      "kr", "dz", "gl", "gm", "n",
    ],
    nuclei: ["a", "u", "o", "i", "e", "ur", "ul", "az", "un"],
    medialCodas: ["", "r", "l", "z", "n", "m", "rk", "lk"],
    maleEndings: [
      "k", "r", "d", "m", "n", "rk", "rd", "grim", "dur", "rik", "mund",
    ],
    femaleEndings: ["a", "i", "ra", "da", "ka", "ri", "di", "na"],
    neutralEndings: ["r", "k", "n", "m", "ul", "az"],
  },
  orkind: {
    onsets: [
      "", "g", "gr", "kr", "k", "z", "zh", "b", "d", "dr", "n", "r",
      "sk", "sh", "t", "thr", "m", "gn",
    ],
    nuclei: ["a", "u", "o", "ug", "ur", "ag", "og", "uk", "ash"],
    medialCodas: ["", "g", "k", "r", "z", "rg", "rk"],
    maleEndings: [
      "g", "k", "rg", "zg", "th", "rk", "gul", "bur", "nash", "rok",
    ],
    femaleEndings: ["a", "ga", "ra", "sha", "ka", "za", "gra"],
    neutralEndings: ["g", "k", "r", "z", "uk", "og"],
  },
};

const PHONEME_VOWELS = new Set(["a", "e", "i", "o", "u"]);

type Ending = "neutral" | "female" | "male";

function leadingConsonants(s: string): number {
  let n = 0;
  for (const ch of s) {
    if (PHONEME_VOWELS.has(ch.toLowerCase())) break;
    n++;
  }
  return n;
}

function trailingConsonants(s: string): number {
  const chars = Array.from(s);
  let n = 0;
  for (let i = chars.length - 1; i >= 0; i--) {
    if (PHONEME_VOWELS.has(chars[i].toLowerCase())) break;
    n++;
  }
  return n;
}

function leadingVowels(s: string): number {
  let n = 0;
  for (const ch of s) {
    if (!PHONEME_VOWELS.has(ch.toLowerCase())) break;
    n++;
  }
  return n;
}

function trailingVowels(s: string): number {
  const chars = Array.from(s);
  let n = 0;
  for (let i = chars.length - 1; i >= 0; i--) {
    if (!PHONEME_VOWELS.has(chars[i].toLowerCase())) break;
    n++;
  }
  return n;
}

/**
 * When appended to `built`, the result must keep consonant and vowel clusters below three; the word-initial onset
 * may instead carry a leading cluster such as "thr".
 *
 * @param {string[]} candidates the list of candidate phonemes from which to pick
 * @param {string} built the string being built so far
 * @param {boolean} allowLeadingCluster whether to allow a leading cluster
 * @returns {string} the selected phoneme
 */
function pickComponent(candidates: string[], built: string, allowLeadingCluster: boolean): string {
  let maxLeadingCons = 2 - trailingConsonants(built);
  if (allowLeadingCluster) maxLeadingCons = 3;
  if (maxLeadingCons < 0) maxLeadingCons = 0;

  let maxLeadingVow = 2 - trailingVowels(built);
  if (maxLeadingVow < 0) maxLeadingVow = 0;

  const allowed = candidates.filter(
    (c) => leadingConsonants(c) <= maxLeadingCons && leadingVowels(c) <= maxLeadingVow,
  );
  return pick(allowed);
}

function withoutEmpty(items: string[]): string[] {
  return items.filter((x) => x !== "");
}

function pickWithMaxTrailing(candidates: string[], maxTrailing: number): string {
  const allowed = candidates.filter((c) => trailingConsonants(c) <= maxTrailing);
  return pick(allowed);
}

function generatePhonemes(length: number, ending: Ending, set: PhonemeSet): string {
  let built = "";

  for (let i = 0; i < length; i++) {
    built += pickComponent(set.onsets, built, i === 0);

    if (i === length - 1) {
      built += pickWithMaxTrailing(set.nuclei, 1);
    } else {
      built += pick(set.nuclei);
    }

    if (i < length - 1) {
      let codas = set.medialCodas;
      if (trailingVowels(built) > 0) {
        codas = withoutEmpty(codas);
      }
      built += pickComponent(codas, built, false);
    }
  }

  let endings: string[];
  switch (ending) {
    case "female":
      endings = set.femaleEndings;
      break;
    case "male":
      endings = set.maleEndings;
      break;
    default:
      endings = set.neutralEndings;
  }

  built += pickComponent(endings, built, false);

  const name = built;
  return name.charAt(0).toUpperCase() + name.slice(1);
}

export function generateRandomPhonemeName(gender: Gender, culture: PhonemeCulture): Name {
  const set = phonemeSets[culture];

  const givenLength = pick([1, 1, 2, 2, 2]);
  const surnameLength = pick([1, 1, 2, 2, 2]);

  let givenEnding: Ending = "neutral";
  if (gender === Female) givenEnding = "female";
  else if (gender === Male) givenEnding = "male";

  return {
    givenName: generatePhonemes(givenLength, givenEnding, set),
    surname: generatePhonemes(surnameLength, "neutral", set),
    gender,
  };
}
