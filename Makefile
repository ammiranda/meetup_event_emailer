DATA_DIR := $(shell pwd)/data

build-image:
	docker build --no-cache -t amiranda/meetup-emailer .

push-image:
	docker push amiranda/meetup-emailer

run-docker:
	docker run --rm -v "$(DATA_DIR):/data" --env-file .env amiranda/meetup-emailer

create-env:
	cp .env.example .env