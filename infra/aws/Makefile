SHELL := /bin/sh
CLOUD_REGION_NAME := $(if $(CLOUD_REGION_NAME),$(CLOUD_REGION_NAME),us-east-2)
ENV := 	$(if $(ENV),$(ENV),staging)
STACK := $(if $(STACK),$(STACK),pululapp-core)

deploy:
	aws cloudformation deploy --region $(CLOUD_REGION_NAME) --template-file template.yml --stack-name $(STACK)-$(ENV) --parameter-overrides PululappEnv=$(ENV) --capabilities CAPABILITY_NAMED_IAM

deploy-cluster:
	aws cloudformation deploy --region $(CLOUD_REGION_NAME) --template-file template-cluster.yml --stack-name $(STACK)-cluster-$(ENV) --parameter-overrides PululappEnv=$(ENV) --capabilities CAPABILITY_NAMED_IAM

