.PHONY: test
test:
	@go test -count=1 ./...

.PHONY: tpl
tpl:
	@go run cmd/tpl/main.go --section=$(S) --exercise=$(E) 