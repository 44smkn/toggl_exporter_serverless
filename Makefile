.PHONY: build
build:
	sam build

.PHONY: deploy
deploy:
	sam deploy --stack-name toggl-exporter --parameter-overrides TogglApiKey=$(TOGGL_API_KEY) --capabilities CAPABILITY_IAM --resolve-s3 --role-arn arn:aws:iam::171457761414:role/CloudFormationFullAccessForSAM --no-confirm-changeset --no-fail-on-empty-changeset