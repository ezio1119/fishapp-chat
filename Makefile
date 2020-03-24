DC = docker-compose
CURRENT_DIR = $(shell pwd)
API = chat

sql:
	docker run --rm -v $(CURRENT_DIR)/migrate/sql:/sql \
	migrate/migrate:latest create -ext sql -dir /sql ${ARG}

sql-doc:
	docker run --rm --net=api-gateway_default -v $(CURRENT_DIR)/db:/work ezio1119/tbls \
	doc -f -t svg mysql://root:password@$(API)-db:3306/$(API)_DB ./

proto:
	docker run --rm -w $(CURRENT_DIR) \
	-v $(CURRENT_DIR)/schema/$(API):/schema \
	-v $(CURRENT_DIR)/src/interfaces/controllers/$(API)_grpc:$(CURRENT_DIR) \
	thethingsindustries/protoc \
	-I/schema \
	-I/usr/include/github.com/envoyproxy/protoc-gen-validate \
	--go_out=plugins=grpc:. \
	--validate_out="lang=go:." \
	--doc_out=markdown,README.md:/schema \
	$(API).proto

migrate:
	docker run --rm -it --name migrate --net=api-gateway_default \
	-v $(CURRENT_DIR)/db/sql:/sql migrate/migrate:latest \
	-path /sql/ -database "mysql://root:password@tcp($(API)-db:3306)/$(API)_DB" up

cli:
		docker run --entrypoint sh --rm -it --net=api-gateway_default namely/grpc-cli
		# call $(API):50051 $(API)_grpc.ChatService.$(m) "$(q)" $(o)

test:
	$(DC) exec $(API) sh -c "go test -v -coverprofile=cover.out -coverpkg=$(a) $(a) && \
									go tool cover -html=cover.out -o ./cover.html" && \
	open ./src/cover.html

up:
	$(DC) up -d

ps:
	$(DC) ps

build:
	$(DC) build

down:
	$(DC) down

exec:
	$(DC) exec $(API) sh

logs:
	docker logs -f --tail 100 $(API)_$(API)_1

redis:
	$(DC) exec $(API)-kvs sh
