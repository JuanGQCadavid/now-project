SHELL := /bin/sh
CLOUD_REGION_NAME := $(if $(CLOUD_REGION_NAME),$(CLOUD_REGION_NAME),us-east-2)
ENV := 	$(if $(ENV),$(ENV),staging)
STACK := $(if $(STACK),$(STACK),spots-core-service)
PLATFORM := $(if $(PLATFORM),$(PLATFORM),arm64)
VERSION := $(if $(VERSION),$(VERSION),)
ECR_URI := $(shell aws cloudformation --region $(CLOUD_REGION_NAME) describe-stacks --stack-name pululapp-core-prod --query  "Stacks[0].Outputs[?ExportName=='spot-creator-ecr-uri'].OutputValue" --output text)

deploy:
	aws cloudformation deploy --region $(CLOUD_REGION_NAME) --template-file template.yml --stack-name $(STACK)-$(ENV) --parameter-overrides PululappEnv=$(ENV) ServiceImageVersion=$(VERSION) --capabilities CAPABILITY_NAMED_IAM

build:
	docker build -t $(STACK)-$(ENV)-$(VERSION) .

ecr-login:
	aws ecr get-login-password --region $(CLOUD_REGION_NAME) | docker login --username AWS --password-stdin 732596568988.dkr.ecr.$(CLOUD_REGION_NAME).amazonaws.com

ecr-deploy: ecr-login
	rm -rf tmp \
	&& mkdir tmp \
	&& cp -r ../pkgs tmp/pkgs \
	&& docker buildx build --platform=linux/$(PLATFORM) --build-arg CMD="lambda" --build-arg GOARCH="$(PLATFORM)" -t $(STACK):$(VERSION) --load . \
	&& rm -rf tmp \
	&& docker tag $(STACK):$(VERSION) $(ECR_URI):$(VERSION) \
	&& docker push $(ECR_URI):$(VERSION)

docker-buildx-driver:
	docker buildx create --use

publish: ecr-deploy deploy

run: 
	source .env && go run cmd/httpserver/main.go --debug
