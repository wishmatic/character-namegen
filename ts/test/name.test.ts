import { describe, expect, it } from "vitest";

import {
  Female,
  IRL_CULTURES,
  Male,
  NonBinary,
  PHONEME_CULTURES,
  generate,
  type Gender,
} from "../src/index";

const VOWELS = new Set(["a", "e", "i", "o", "u"]);

function hasAwkwardConsonantCluster(text: string): boolean {
  let streak = 0;
  let seenVowel = false;

  for (const ch of text) {
    if (VOWELS.has(ch.toLowerCase())) {
      streak = 0;
      seenVowel = true;
      continue;
    }

    streak++;
    if (streak >= 3 && seenVowel) return true;
  }

  return false;
}

function hasAwkwardVowelCluster(text: string): boolean {
  let streak = 0;

  for (const ch of text) {
    if (VOWELS.has(ch.toLowerCase())) {
      streak++;
      if (streak >= 3) return true;
    } else {
      streak = 0;
    }
  }

  return false;
}

describe("generate", () => {
  it("produces non-empty names and respects an explicit gender", () => {
    const genders: (Gender | "")[] = ["", Male, Female, NonBinary];

    for (const gender of genders) {
      const name = generate(gender);
      expect(name.givenName).not.toBe("");
      expect(name.surname).not.toBe("");
      if (gender !== "") {
        expect(name.gender).toBe(gender);
      }
    }
  });

  it("supports every phoneme culture", () => {
    for (const culture of PHONEME_CULTURES) {
      for (let i = 0; i < 20; i++) {
        const name = generate("", culture);
        expect(name.givenName).not.toBe("");
        expect(name.surname).not.toBe("");
      }
    }
  });

  it("supports every IRL culture", () => {
    for (const culture of IRL_CULTURES) {
      for (let i = 0; i < 20; i++) {
        const name = generate("", culture);
        expect(name.givenName).not.toBe("");
        expect(name.surname).not.toBe("");
      }
    }
  });

  it("avoids awkward consonant and vowel clusters in phoneme names", () => {
    for (const culture of PHONEME_CULTURES) {
      for (let i = 0; i < 500; i++) {
        const name = generate("", culture);
        expect(hasAwkwardConsonantCluster(name.givenName)).toBe(false);
        expect(hasAwkwardVowelCluster(name.givenName)).toBe(false);
        expect(hasAwkwardConsonantCluster(name.surname)).toBe(false);
        expect(hasAwkwardVowelCluster(name.surname)).toBe(false);
      }
    }
  });
});
