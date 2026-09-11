import type { IRLCulture } from "./culture";

export interface CultureData {
  givenMale: string[];
  givenFemale: string[];
  surnames: string[];
}

// Maps the key used in data file names to the public culture name. The keys don't always match the display names
// (e.g. "american-ibero" -> "Iberoamerican", "gulf" -> "Arabic (Gulf)").
const CULTURE_BY_KEY: Record<string, IRLCulture> = {
  afghan: "Afghan",
  "african-central": "African (Central)",
  "african-southern": "African (Southern)",
  "african-west": "African (West)",
  "american-ibero": "Iberoamerican",
  anglophone: "Anglophone",
  "asian-south": "Asian (South)",
  "asian-southeast": "Asian (Southeast)",
  chinese: "Chinese",
  ethiopian: "Ethiopian",
  "european-central": "European (Central)",
  "european-southern": "European (Southern)",
  fijian: "Fijian",
  french: "French",
  georgian: "Georgian",
  germanic: "Germanic",
  greek: "Greek",
  gulf: "Arabic (Gulf)",
  iranian: "Iranian",
  irish: "Irish",
  japanese: "Japanese",
  korean: "Korean",
  levantine: "Arabic (Levantine)",
  maghrebi: "Arabic (Maghrebi)",
  nordic: "Nordic",
  russian: "Russian",
  "slavic-south": "Slavic (South)",
  turkic: "Turkic",
};

function lines(raw: string): string[] {
  return raw
    .split("\n")
    .map((line) => line.trim())
    .filter((line) => line !== "");
}

const raw = import.meta.glob("../data/*.txt", {
  eager: true,
  query: "?raw",
  import: "default",
}) as Record<string, string>;

export const irlData: Record<IRLCulture, CultureData> = (() => {
  const data = {} as Record<IRLCulture, CultureData>;

  for (const [path, content] of Object.entries(raw)) {
    const file = path.slice(path.lastIndexOf("/") + 1);

    let key: string;
    let field: keyof CultureData;
    if (file.endsWith("_given_female.txt")) {
      key = file.slice(0, -"_given_female.txt".length);
      field = "givenFemale";
    } else if (file.endsWith("_given_male.txt")) {
      key = file.slice(0, -"_given_male.txt".length);
      field = "givenMale";
    } else if (file.endsWith("_surnames.txt")) {
      key = file.slice(0, -"_surnames.txt".length);
      field = "surnames";
    } else {
      continue;
    }

    const culture = CULTURE_BY_KEY[key];
    if (!culture) continue;

    data[culture] ??= { givenMale: [], givenFemale: [], surnames: [] };
    data[culture][field] = lines(content);
  }

  return data;
})();
