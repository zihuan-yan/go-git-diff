local_package=github.com/lower-coder/go-git-diff

lint: deps
	@$(set_env) go fmt ./...
	@$(set_env) goimports -local $(local_package) -w .
	@$(set_env) go mod tidy
.PHONY: vet
vet:
	@$(set_env) go vet ./...
.PHONY: deps
deps:
	@$(set_env) hash goimports > /dev/null 2>&1; if [ $$? -ne 0 ]; then \
		go get golang.org/x/tools/cmd/goimports; \
	fi