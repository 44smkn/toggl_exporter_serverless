.PHONY: build
build:
	sam build

.PHONY: deploy
deploy:
	sam deploy --no-confirm-changeset --no-fail-on-empty-changeset --stack-name toggl-exporter --parameter-overrides TogglApiKey=$(TOGGL_API_KEY) --capabilities CAPABILITY_IAM