import urllib.request
from pathlib import Path

# Credit: https://github.com/van-himmelheimer/German-Name-Generator

GITHUB_BASE_URL = "https://raw.githubusercontent.com"

SURNAME_URL = f"{GITHUB_BASE_URL}/van-himmelheimer/German-Name-Generator/refs/heads/master/nachnamen_deutsch"
MALE_URL = f"{GITHUB_BASE_URL}/van-himmelheimer/German-Name-Generator/refs/heads/master/vornamen_m_deutsch"
FEMALE_URL = f"{GITHUB_BASE_URL}/van-himmelheimer/German-Name-Generator/refs/heads/master/vornamen_w_deutsch"

OUT_PATH = Path(__file__).parent / "tmp_names"


def main():
    OUT_PATH.mkdir(parents=True, exist_ok=True)

    _process_surnames()
    _process_given("male", MALE_URL)
    _process_given("female", FEMALE_URL)


def _process_surnames():
    lines = _fetch_lines(SURNAME_URL)

    seen: set[str] = set()
    unique: list[str] = []

    for line in lines:
        name = line.split()[0].strip()

        if name and name not in seen:
            seen.add(name)
            unique.append(name)

    _write_list(OUT_PATH / "german_surnames.txt", unique)


def _process_given(gender: str, url: str):
    lines = _fetch_lines(url)

    seen: set[str] = set()
    unique: list[str] = []

    for line in lines:
        name = line.strip()

        if name and name not in seen:
            seen.add(name)
            unique.append(name)

    _write_list(OUT_PATH / f"german_given_{gender}.txt", unique)


def _fetch_lines(url: str) -> list[str]:
    with urllib.request.urlopen(url) as response:
        text = response.read().decode("utf-8")

    return [line for line in text.splitlines() if line.strip()]


def _write_list(path: Path, items: list[str]):
    path.write_text("\n".join(items) + "\n", encoding="utf-8")


if __name__ == "__main__":
    main()
