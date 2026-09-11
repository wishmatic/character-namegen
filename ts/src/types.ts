import type { Gender } from "./gender";

export interface Name {
  givenName: string;
  surname: string;
  gender: Gender;
}
