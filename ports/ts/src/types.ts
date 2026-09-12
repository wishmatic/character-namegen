import type { Gender } from "./gender.js";

export interface Name {
  givenName: string;
  surname: string;
  gender: Gender;
}
