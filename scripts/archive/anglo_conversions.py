import csv
import io
import urllib.request
from pathlib import Path

# Credit (forenames): https://github.com/COMHIS/names_and_genders
# Credit (surnames): https://github.com/matt40k/Names

GITHUB_BASE_URL = "https://raw.githubusercontent.com"

FORENAME_URL = f"{GITHUB_BASE_URL}/COMHIS/names_and_genders/refs/heads/master/historical_forenames_genders.csv"
SURNAME_URL = f"{GITHUB_BASE_URL}/matt40k/Names/refs/heads/main/Surname_Autumn2014.csv"

OUT_PATH = Path(__file__).parent / "tmp_names"


def main():
    OUT_PATH.mkdir(parents=True, exist_ok=True)

    _process_surnames()
    _process_given_names()


def _process_surnames():
    rows = _fetch_csv(SURNAME_URL)

    seen: set[str] = set()
    unique: list[str] = []

    for row in rows:
        name = row["Surname"].strip().title()

        if name and name not in seen:
            seen.add(name)
            unique.append(name)

    _write_list(OUT_PATH / "anglo_surnames.txt", unique)


def _process_given_names():
    rows = _fetch_csv(FORENAME_URL)

    male: list[str] = []
    female: list[str] = []
    seen_male: set[str] = set()
    seen_female: set[str] = set()

    for row in rows:
        name = row["forename"].strip().title()
        gender = row["gender"].strip().upper()

        if gender == "M" and name not in seen_male:
            seen_male.add(name)
            male.append(name)
        elif gender == "F" and name not in seen_female:
            seen_female.add(name)
            female.append(name)

    _write_list(OUT_PATH / "anglo_given_male.txt", male)
    _write_list(OUT_PATH / "anglo_given_female.txt", female)


def _fetch_csv(url: str) -> list[dict[str, str]]:
    with urllib.request.urlopen(url) as response:
        text = response.read().decode("utf-8-sig")

    return list(csv.DictReader(io.StringIO(text)))


def _write_list(path: Path, items: list[str]):
    path.write_text("\n".join(items) + "\n", encoding="utf-8")


if __name__ == "__main__":
    main()
