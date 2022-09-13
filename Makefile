GO_META_MOD := github.com/avakarev/go-buildmeta

LDFLAGS += -s -w
LDFLAGS += -X ${GO_META_MOD}.Commit=${GITHUB_SHA}
LDFLAGS += -X ${GO_META_MOD}.Ref=${GITHUB_REF}
LDFLAGS += -X ${GO_META_MOD}.Date=$(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

lint:
	@echo ">> Running revive..."
	@revive -config .revive.toml -formatter friendly ./...
	@echo ">> Running staticcheck..."
	@staticcheck ./...

vet:
	@echo ">> Vetting..."
	@go vet ./...

sec:
	@echo ">> Auditing..."
	@gosec -quiet -tests ./...

test:
	@echo ">> Running tests..."
	@go test -v -race ./...
.PHONY: test

testcov:
	@echo ">> Generating coverage report..."
	@go test -v -race -coverprofile=coverage.txt ./...
	@go tool cover -html=coverage.txt

setup-ci:
	@go install github.com/mgechev/revive@latest
	@go install github.com/securego/gosec/v2/cmd/gosec@latest
	@go install honnef.co/go/tools/cmd/staticcheck@latest

ci: lint vet sec test

build:
	@echo ">> Building ./bin/server..."
	@CGO_ENABLED=0 go build -ldflags '${LDFLAGS}' -o ./bin/server ./cmd/server
