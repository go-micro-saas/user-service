.PHONY: store-configuration
# run :-->: run database-migration
store-configuration:
	go run ./app/user-service/cmd/store-configuration/... -conf=./app/user-service/configs

.PHONY: run-database-migration
# run :-->: run database-migration
run-database-migration:
	go run ./app/user-service/cmd/database-migration/... -conf=./app/user-service/configs

.PHONY: run-user-service
# run service :-->: run user-service
run-user-service:
	go run ./app/user-service/cmd/user-service/... -conf=./app/user-service/configs

.PHONY: testing-user-service
# testing service :-->: testing user-service
testing-user-service:
	go run testdata/get-node-id/main.go

.PHONY: run-service
# run service :-->: run user-service
run-service:
	#@$(MAKE) run-user-service
	go run ./app/user-service/cmd/user-service/... -conf=./app/user-service/configs

.PHONY: testing-service
# testing service :-->: testing user-service
testing-service:
	go run testdata/get-node-id/main.go
