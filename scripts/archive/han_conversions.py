import csv
import io
import urllib.request
from pathlib import Path

from pypinyin import lazy_pinyin, Style

# Credit: https://github.com/psychbruce/ChineseNames/tree/main

GITHUB_BASE_URL = "https://raw.githubusercontent.com"

SURNAME_URL = (
    f"{GITHUB_BASE_URL}/psychbruce/ChineseNames/refs/heads/main/data-csv/familyname.csv"
)
GIVEN_NAME_URL = f"{GITHUB_BASE_URL}/psychbruce/ChineseNames/refs/heads/main/data-csv/top1000name.prov.csv"

OUT_PATH = Path(__file__).parent / "tmp_names"

# Gender ratio thresholds are based on n.male / (n.male + n.female).
#
# - A name with ratio >= MALE_THRESH is considered male.
# - A name with ratio <= FEMALE_THRESH is considered female.
# - Names with ratios in between are considered unisex and are included in both lists.

MALE_THRESH = 0.65
FEMALE_THRESH = 0.35


def main():
    OUT_PATH.mkdir(parents=True, exist_ok=True)

    _process_surnames()
    _process_given_names()


def _process_surnames():
    rows = _fetch_csv(SURNAME_URL)

    seen: set[str] = set()
    unique: list[str] = []

    for row in rows:
        romanized = _romanize(row["surname"])

        if romanized not in seen:
            seen.add(romanized)
            unique.append(romanized)

    out = OUT_PATH / "han_surnames.txt"

    _write_list(out, unique)


def _process_given_names():
    rows = _fetch_csv(GIVEN_NAME_URL)

    male: list[str] = []
    female: list[str] = []

    seen_male: set[str] = set()
    seen_female: set[str] = set()

    for row in rows:
        romanized = _romanize(row["name"])

        n_male = int(row["n.male"])
        n_female = int(row["n.female"])

        total = n_male + n_female

        if total == 0:
            continue

        male_ratio = n_male / total

        if male_ratio >= MALE_THRESH and romanized not in seen_male:
            seen_male.add(romanized)
            male.append(romanized)

        elif male_ratio <= FEMALE_THRESH and romanized not in seen_female:
            seen_female.add(romanized)
            female.append(romanized)
        else:
            # Unisex names are added to both lists.

            if romanized not in seen_male:
                seen_male.add(romanized)
                male.append(romanized)
            if romanized not in seen_female:
                seen_female.add(romanized)
                female.append(romanized)

    _write_list(OUT_PATH / "han_given_male.txt", male)
    _write_list(OUT_PATH / "han_given_female.txt", female)


def _fetch_csv(url: str) -> list[dict[str, str]]:
    with urllib.request.urlopen(url) as response:
        text = response.read().decode("utf-8-sig")

    return list(csv.DictReader(io.StringIO(text)))


def _romanize(hanzi: str) -> str:
    parts = lazy_pinyin(hanzi, style=Style.NORMAL)

    return " ".join(text.capitalize() for text in parts)


def _write_list(path: Path, items: list[str]):
    path.write_text("\n".join(items) + "\n", encoding="utf-8")


if __name__ == "__main__":
    main()
