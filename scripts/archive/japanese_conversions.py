import csv
import io
import urllib.request
from pathlib import Path

# Credit: https://github.com/shuheilocale/japanese-personal-name-dataset

GITHUB_BASE_URL = "https://raw.githubusercontent.com/shuheilocale/japanese-personal-name-dataset/refs/heads/main/"

SURNAME_URL = (
    f"{GITHUB_BASE_URL}/japanese_personal_name_dataset/dataset/last_name_org.csv"
)
MALE_URL = (
    f"{GITHUB_BASE_URL}/japanese_personal_name_dataset/dataset/first_name_man_org.csv"
)
FEMALE_URL = (
    f"{GITHUB_BASE_URL}/japanese_personal_name_dataset/dataset/first_name_woman_org.csv"
)

OUT_PATH = Path(__file__).parent / "tmp_names"


def main():
    OUT_PATH.mkdir(parents=True, exist_ok=True)

    _process_surnames()
    _process_given("male", MALE_URL)
    _process_given("female", FEMALE_URL)


def _process_surnames():
    print(f"Downloading {SURNAME_URL} ...")
    lines = _fetch_lines(SURNAME_URL)

    seen: set[str] = set()
    unique: list[str] = []
    for line in lines:
        parts = line.split(",")
        if len(parts) < 4:
            continue
        romaji = parts[3].strip().title()
        if romaji and romaji not in seen:
            seen.add(romaji)
            unique.append(romaji)

    _write_list(OUT_PATH / "japanese_surnames.txt", unique)
    print(f"  -> {len(unique)} unique surnames")


def _process_given(gender: str, url: str):
    lines = _fetch_lines(url)

    seen: set[str] = set()
    unique: list[str] = []

    for line in lines:
        parts = line.split(",")

        if len(parts) < 2:
            continue

        romaji = parts[1].strip().title()
        if romaji and romaji not in seen:
            seen.add(romaji)
            unique.append(romaji)

    _write_list(OUT_PATH / f"japanese_given_{gender}.txt", unique)


def _fetch_lines(url: str) -> list[str]:
    with urllib.request.urlopen(url) as response:
        text = response.read().decode("utf-8-sig")

    return [line for line in text.splitlines() if line.strip()]


def _write_list(path: Path, items: list[str]):
    path.write_text("\n".join(items) + "\n", encoding="utf-8")


if __name__ == "__main__":
    main()
