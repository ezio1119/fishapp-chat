DC = docker-compose
CURRENT_DIR = $(shell pwd)
API = chat

sql:
	docker run --rm -v $(CURRENT_DIR)/db/sql:/sql \
	migrate/migrate:latest create -ext sql -dir /sql ${ARG}

sql-doc:
	docker run --rm --net=fishapp-net -v $(CURRENT_DIR)/db:/work ezio1119/tbls \
	doc -f -t svg mysql://root:password@$(API)-db:3306/$(API)_DB ./

proto:
	docker run --rm -v $(CURRENT_DIR)/src/pb:/pb -v $(CURRENT_DIR)/schema:/proto ezio1119/protoc \
	-I/proto \
	-I/go/src/github.com/envoyproxy/protoc-gen-validate \
	--go_opt=paths=source_relative \
	--go_out=plugins=grpc:/pb \
	--validate_out="lang=go,paths=source_relative:/pb" \
	chat/chat.proto event/event.proto

migrate:
	docker run --rm -it --name migrate --net=fishapp-net \
	-v $(CURRENT_DIR)/db/sql:/sql migrate/migrate:latest \
	-path /sql/ -database "mysql://root:password@tcp($(API)-db:3306)/$(API)_DB" up

cli:
		docker run --entrypoint sh --rm -it --net=fishapp-net namely/grpc-cli
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
