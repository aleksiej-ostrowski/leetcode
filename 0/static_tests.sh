# python3 -m pip install --upgrade pip
# python3 -m pip install git+https://github.com/psf/black
# python3 -m pip install flake8
# python3 -m pip install mypy
# python3 -m pip install pylint

# python3 -m pip install types-requests

echo "" > report.txt
python3 -m black *.py >> report.txt # --fast
python3 -m pylint *.py >> report.txt
python3 -m mypy --strict *.py >> report.txt
python3 -m flake8 *.py >> report.txt

# https://vald-phoenix.github.io/pylint-errors/plerr/errors/basic/C0116.html
