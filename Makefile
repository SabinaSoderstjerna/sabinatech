# Run a complete build
.PHONY: all
all: \
	markdown-lint \
	json-lint \
	html-lint \
	css-lint \
	yaml-lint \
	go-lint \
	go-test \
	go-mod-tidy \
	go-build

export GO111MODULE := on

golangci_lint ?= files/golangci-lint/

# Clean after a build
.PHONY: clean
clean:
	# Removing files directory
	rm -rf files

# installing npm to be able to run markdown, json, html, css and yaml linters
.PHONY: npm-install
npm-install:
	npm install --no-save --prefix ./files \
	markdownlint-cli \
	jsonlint-cli \
	htmlhint \
	stylelint stylelint-config-standard \
	yaml-validator

# markdown-lint: lint checking markdown files
.PHONY: markdown-lint
markdown-lint: npm-install
	markdownlint **/*.md --ignore files

# json-lint: lint checking json files with jsonlint-cli
.PHONY: json-lint
json-lint: npm-install
	jsonlint-cli **/*.json --ignore files

# html-lint: lint checking json files with htmlhint
.PHONY: html-lint
html-lint: npm-install
	# custom rules: id-class-value=dash
	htmlhint --rules id-class-value=dash --ignore files

# css-lint: lint checking css files with stylelint
.PHONY: css-lint
css-lint: npm-install
	@[ -f ./files/.stylelintrc.json ] || echo "{\n\t\"extends\": \"stylelint-config-standard\"\n}" > ./files/.stylelintrc.json
	stylelint static/stylesheets/*.css --allow-empty-input --config ./files/.stylelintrc.json --cache --cache-location ./files/.stylelintcache

# yaml-lint: lint checking yaml files with yaml-validator
.PHONY: yaml-lint
yaml-lint: npm-install
	yaml-validator *.yaml

# go-lint: lint checking go files with GolangCI-Lint
.PHONY: go-lint
go-lint: $(golangci_lint)
	# Disabled
	# wsl: too strict
	# interfacer: deprecated
	$<golangci-lint run --enable-all --disable wsl,interfacer

$(golangci_lint): go.mod
	@[ -d $@ ] || mkdir -p $@ | go build -o $@ github.com/golangci/golangci-lint/cmd/golangci-lint

# go-test: run all go tests
.PHONY: go-test
go-test:
	go test -race ./...

# go-mod-tidy: cleaning up mod file
.PHONY: go-mod-tidy
go-mod-tidy:
	go mod tidy -v

# go-build: ensuring the library can be cross-compiled to supported OSes
.PHONY: go-build
go-build:
    GOOS=darwin GOARCH=amd64 go build ./...
    GOOS=windows GOARCH=amd64 go build ./...
    GOOS=linux GOARCH=amd64 go build ./...
