from __future__ import annotations

import re
from pathlib import Path

from names_dataset import NameDataset
from unidecode import unidecode

OUT_PATH = Path(__file__).parent / "names"
TARGET_NAMES_PER_FILE = 200

GROUPS: dict[str, list[str]] = {
    "afghan": ["AF"],
    "african-central": ["CM", "BI"],
    "african-southern": ["ZA", "BW", "NA", "AO"],
    "african-west": ["GH", "NG", "BF"],
    "american-ibero": [
        "MX",
        "GT",
        "CU",
        "DO",
        "AR",
        "CO",
        "CL",
        "PE",
        "BR",
    ],
    "anglophone": ["GB", "US", "JM"],
    "asian-south": ["IN", "BD", "MV"],
    "asian-southeast": ["ID", "MY", "BN", "PH", "KH"],
    "chinese": ["CN"],
    "ethiopian": ["ET"],
    "european-central": ["CZ", "PL", "HU", "HR", "SI"],
    "european-southern": ["IT", "ES", "PT"],
    "fijian": ["FJ"],
    "french": ["FR"],
    "georgian": ["GE"],
    "germanic": ["DE", "NL"],
    "greek": ["GR"],
    "gulf": ["SA", "AE"],
    "iranian": ["IR"],
    "irish": ["IE"],
    "japanese": ["JP"],
    "korean": ["KR"],
    "levantine": ["SY", "LB", "JO", "PS"],
    "maghrebi": ["DZ", "MA"],
    "nordic": ["DK", "NO", "SE", "IS", "FI"],
    "russian": ["RU"],
    "slavic-south": ["RS"],
    "turkic": ["TR", "AZ", "TM"],
}

ROMANIZE_GROUPS: set[str] = {
    "afghan",
    "asian-south",
    "asian-southeast",
    "chinese",
    "ethiopian",
    "georgian",
    "greek",
    "gulf",
    "iranian",
    "japanese",
    "korean",
    "levantine",
    "maghrebi",
    "russian",
    "slavic-south",
    "turkic",
}

_HAS_LATIN = re.compile(r"[A-Za-z]")
_LATIN_ONLY = re.compile(r"^[A-Za-z \'\-]+$")
_HAS_VOWEL = re.compile(r"[AEIOUaeiou]")


def _is_latin(name: str) -> bool:
    return bool(_LATIN_ONLY.match(name))


def _romanize(name: str) -> str | None:
    if _is_latin(name):
        if len(name) < 2 or not _HAS_VOWEL.search(name):
            return None

        return name

    transliterated = unidecode(name).strip()

    # unidecode can produce empty or unusable names for some chars

    if not transliterated or not _HAS_LATIN.search(transliterated):
        return None

    transliterated = re.sub(r"\s+", " ", transliterated).strip(" -'")

    # Must be at least 2 chars and contain a vowel (filters abjad consonant-skeletons)

    if len(transliterated) < 2 or not _HAS_VOWEL.search(transliterated):
        return None

    return transliterated.title()


def _collect_first_names(
    dataset: NameDataset,
    countries: list[str],
    gender: str,
    n: int,
    romanize: bool = False,
) -> list[str]:
    seen: set[str] = set()
    result: list[str] = []

    for cc in countries:
        data = dataset.get_top_names(
            n=n, gender=gender, country_alpha2=cc, use_first_names=True
        )
        names = data.get(cc, {}).get("M" if gender == "Male" else "F", [])

        for name in names:
            if romanize:
                name = _romanize(name)

                if name is None:
                    continue
            key = name.lower()

            if key not in seen:
                seen.add(key)
                result.append(name)

    return result[:n]


def _collect_last_names(
    dataset: NameDataset,
    countries: list[str],
    n: int,
    romanize: bool = False,
) -> list[str]:
    seen: set[str] = set()
    result: list[str] = []

    for cc in countries:
        data = dataset.get_top_names(n=n, use_first_names=False, country_alpha2=cc)
        country_data = data.get(cc, {})

        if isinstance(country_data, list):
            names = country_data
        else:
            names = []

            for v in country_data.values():
                if isinstance(v, list):
                    names.extend(v)

        for name in names:
            if romanize:
                name = _romanize(name)

                if name is None:
                    continue

            key = name.lower()

            if key not in seen:
                seen.add(key)
                result.append(name)

    return result[:n]


def _write(path: Path, items: list[str]):
    path.write_text("\n".join(items) + "\n", encoding="utf-8")


def main():
    OUT_PATH.mkdir(parents=True, exist_ok=True)

    nd = NameDataset()

    for group, countries in GROUPS.items():
        romanize = group in ROMANIZE_GROUPS

        males = _collect_first_names(
            nd, countries, "Male", TARGET_NAMES_PER_FILE, romanize
        )
        females = _collect_first_names(
            nd, countries, "Female", TARGET_NAMES_PER_FILE, romanize
        )
        surnames = _collect_last_names(nd, countries, TARGET_NAMES_PER_FILE, romanize)

        _write(OUT_PATH / f"{group}_given_male.txt", males)
        _write(OUT_PATH / f"{group}_given_female.txt", females)
        _write(OUT_PATH / f"{group}_surnames.txt", surnames)


if __name__ == "__main__":
    main()
