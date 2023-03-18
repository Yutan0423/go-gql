dev:
	@cd docker && docker-compose up -d --build

db-exec:
	@cd docker && docker-compose exec mysql /bin/bash
	# mysql -u user -ppassword

down:
	@cd docker && docker-compose down --remove-orphans

gql_generate:
	@cd app && gqlgen

sqlc_generate:
	@cd app && sqlc generate

tidy:
	@cd app && go mod tidy