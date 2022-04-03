install:
	virtualenv venv
	source ./venv/bin/activate
	pip install -r requirements.txt
.PHONY: install

dev:
	uvicorn main:app --reload
.PHONY: dev

prod:
	uvicorn main:app
.PHONY: prod