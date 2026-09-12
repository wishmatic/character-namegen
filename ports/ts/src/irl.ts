import { irlData, type CultureData } from "./data.js";
import { Female, Male, type Gender } from "./gender.js";
import type { IRLCulture } from "./culture.js";
import { pick } from "./random.js";
import type { Name } from "./types.js";

export function generateIRLName(culture: IRLCulture, gender: Gender): Name {
  const data = determineNameData(culture);

  const isMaleOrFemale = gender === Male || gender === Female;

  let givenPool: string[];
  if (!isMaleOrFemale) {
    givenPool = [...data.givenMale, ...data.givenFemale];
  } else if (gender === Female) {
    givenPool = data.givenFemale;
  } else {
    givenPool = data.givenMale;
  }

  const givenName = pick(givenPool);
  let surname = pick(data.surnames);

  switch (culture) {
    case "Russian":
      surname = genderizeRussianSurname(surname, gender);
      break;
    case "European (Central)":
      surname = pickCzechSlovakSurname(data.surnames, gender);
      break;
  }

  return { givenName, surname, gender };
}

function determineNameData(culture: IRLCulture): CultureData {
  return irlData[culture];
}

function genderizeRussianSurname(surname: string, gender: Gender): string {
  const isFemale = gender === Female;
  const isMale = gender === Male;
  const lower = surname.toLowerCase();

  if (isFemale) {
    if (lower.endsWith("skiy") || lower.endsWith("skij")) {
      return surname.slice(0, -4) + "skaya";
    }
    if (lower.endsWith("sky")) {
      return surname.slice(0, -3) + "skaya";
    }
    if (lower.endsWith("oy")) {
      return surname.slice(0, -2) + "aya";
    }
    if (lower.endsWith("ev") || lower.endsWith("ov") || lower.endsWith("in")) {
      return surname + "a";
    }
    return surname;
  }

  if (isMale) {
    if (lower.endsWith("skaya")) {
      return surname.slice(0, -5) + "sky";
    }
    if (lower.endsWith("aya")) {
      return surname.slice(0, -3) + "oy";
    }
    if (lower.endsWith("ova") || lower.endsWith("eva") || lower.endsWith("ina")) {
      return surname.slice(0, -1);
    }
    return surname;
  }

  return surname;
}

// Czech/Slovak surnames are stored in both forms, so picking (rather than converting) avoids the fleeting-vowel
// problem (e.g. Havlíček/Havlíčková).
function pickCzechSlovakSurname(surnames: string[], gender: Gender): string {
  for (let i = 0; i < 20; i++) {
    const candidate = pick(surnames);
    if (matchesCzechSlovakGender(candidate, gender)) {
      return candidate;
    }
  }
  return pick(surnames);
}

/**
 * Determine if a surname matches the gender of a Czech/Slovak name.
 *
 * Female forms end in -ová (or unaccented -ova) or -á; soft adjectival -í is invariant and valid for both genders.
 *
 * @param {string} surname the surname to check
 * @param {Gender} gender the gender to match
 * @returns {boolean} whether the surname matches the gender
 */
function matchesCzechSlovakGender(surname: string, gender: Gender): boolean {
  const lower = surname.toLowerCase();
  const isFemale =
    lower.endsWith("ová") || lower.endsWith("ova") || lower.endsWith("á");

  switch (gender) {
    case Male:
      return !isFemale;
    case Female:
      return isFemale || lower.endsWith("í");
    default:
      return true;
  }
}
