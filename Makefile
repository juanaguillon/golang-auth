.PHONY: runnodemon dev test

runnodemon:
	nodemon --watch './**/*.go' --signal SIGTERM --exec 'go' run cmd/main.go
dev:
	APP_ENV=development $(MAKE) runnodemon
test:
	APP_ENV=test $(MAKE) runnodemon
	