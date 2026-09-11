import csv
import io
import urllib.request
from pathlib import Path

# Credit: https://github.com/sctg-development/french-names-extractor

GITHUB_BASE_URL = "https://raw.githubusercontent.com"

FIRSTNAME_URL = f"{GITHUB_BASE_URL}/sctg-development/french-names-extractor/refs/heads/main/firstnames.csv"
LASTNAME_URL = f"{GITHUB_BASE_URL}/sctg-development/french-names-extractor/refs/heads/main/lastnames.csv"

OUT_PATH = Path(__file__).parent / "tmp_names"

MALE_INDEX = "1"
FEMALE_INDEX = "2"


def main():
    OUT_PATH.mkdir(parents=True, exist_ok=True)

    _process_surnames()
    _process_given_names()


def _process_surnames():
    rows = _fetch_csv(LASTNAME_URL)

    seen: set[str] = set()
    unique: list[str] = []

    for row in rows:
        name = row["lastname"].strip().title()

        if name and name not in seen:
            seen.add(name)
            unique.append(name)

    _write_list(OUT_PATH / "french_surnames.txt", unique)


def _process_given_names():
    rows = _fetch_csv(FIRSTNAME_URL)

    male: list[str] = []
    female: list[str] = []
    seen_male: set[str] = set()
    seen_female: set[str] = set()

    for row in rows:
        name = row["firstname"].strip().title()
        gender = row["gender"].strip()

        if gender == MALE_INDEX and name not in seen_male:
            seen_male.add(name)
            male.append(name)
        elif gender == FEMALE_INDEX and name not in seen_female:
            seen_female.add(name)
            female.append(name)

    _write_list(OUT_PATH / "french_given_male.txt", male)
    _write_list(OUT_PATH / "french_given_female.txt", female)


def _fetch_csv(url: str) -> list[dict[str, str]]:
    with urllib.request.urlopen(url) as response:
        text = response.read().decode("utf-8-sig")

    return list(csv.DictReader(io.StringIO(text)))


def _write_list(path: Path, items: list[str]) -> None:
    path.write_text("\n".join(items) + "\n", encoding="utf-8")


if __name__ == "__main__":
    main()
