.PHONY: help
.DEFAULT_GOAL := help

include .env
-include .env.local

ENV ?= dev
TAG ?= ${ENV}
DOCKER_SHELL ?= sh

dc_files = docker-compose.yml
dc = docker compose $(foreach f,$(dc_files),-f $(f))

export ENV
export TAG

help h: ## Show help
	@awk 'BEGIN {FS = ":.*##"} \
	/^[a-zA-Z0-9_. -]+:.*##/ { \
		split($$1,a," "); \
		printf "%-20s %s\n", a[1], $$2 \
	} \
	/^##@/ {printf "\n%s\n", substr($$0,5)}' $(MAKEFILE_LIST)

##@ Docker
config c: ## Show config
	$(dc) config

build b: ## Build images
	$(dc) build

start s: ## Start containers
	$(dc) up -d

stop st: ## Stop containers
	$(dc) down -t0

restart r: stop start ## Restart

run: ## Run shell
	$(dc) run --rm --entrypoint $(DOCKER_SHELL) app

shell sh: ## Exec shell
	$(dc) exec app $(DOCKER_SHELL)

logs l: ## Follow logs
	$(dc) logs -f

ps: ## List containers
	$(dc) ps -a

clean.all ca: stop ## Cleanup all docker
	docker system prune -fa && docker volume prune -fa

##@ SLS
s.print s.p: ## Print
	$(dc) exec app sh -c "sls print"

s.package s.pa: ## Package
	$(dc) exec app sh -c "sls package --verbose"

s.deploy s.d: build ## Deploy
	$(dc) exec app sh -c "sls deploy --verbose"
	
s.remove s.r: ## Remove
	$(dc) exec app sh -c "sls remove --verbose"

s.invoke s.i: ## Invoke
	$(dc) exec app sh -c "sls invoke local -f hello"

##@ Project
project: build ## Create
	@$(dc) run --rm --entrypoint sh app -c "serverless"\
	&& $(dc) run --rm --entrypoint sh app -c "npm install -D serverless-offline"\
	&& $(dc) run --rm --entrypoint sh app -c "sed -i '/^service: app/a plugins:\n  - serverless-offline # serverless-offline needs to be last in the list' serverless.yml"\
	&& $(dc) run --rm --entrypoint sh app -c "sed -i '/^  - serverless-offline/a custom:\n  serverless-offline:\n    host: 0.0.0.0\n    httpPort: 3000\\n    lambdaPort: 3002\' serverless.yml"\
	&& make start\
	&& make logs

project.dynamodb: ## Create dynamoDB
	rm docker-compose.yml\
	&& mv docker-compose-dynamodb.yml docker-compose.yml\
	&& rm .env && mv .env.dynamodb .env\
	&& make project