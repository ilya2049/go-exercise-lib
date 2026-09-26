.PHONY: test
test:
	@go test -count=1 ./internal/lib/...

.PHONY: tpl
tpl:
	@go run cmd/tpl/main.go --section=$(S) --exercise=$(E) 