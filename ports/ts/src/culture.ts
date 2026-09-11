export type IRLCulture =
  | "Afghan"
  | "African (Central)"
  | "African (Southern)"
  | "African (West)"
  | "Anglophone"
  | "Arabic (Gulf)"
  | "Arabic (Levantine)"
  | "Arabic (Maghrebi)"
  | "Asian (South)"
  | "Asian (Southeast)"
  | "Chinese"
  | "Ethiopian"
  | "European (Central)"
  | "European (Southern)"
  | "Fijian"
  | "French"
  | "Georgian"
  | "Germanic"
  | "Greek"
  | "Iberoamerican"
  | "Iranian"
  | "Irish"
  | "Japanese"
  | "Korean"
  | "Nordic"
  | "Russian"
  | "Slavic (South)"
  | "Turkic";

export type PhonemeCulture = "phonemes" | "elvish" | "fae" | "khuzdul" | "orkind";

export type Culture = IRLCulture | PhonemeCulture;

export const IRL_CULTURES: IRLCulture[] = [
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
];

export const PHONEME_CULTURES: PhonemeCulture[] = [
  "phonemes",
  "elvish",
  "fae",
  "khuzdul",
  "orkind",
];

export const SUPPORTED_CULTURES: Culture[] = [...IRL_CULTURES, ...PHONEME_CULTURES];
