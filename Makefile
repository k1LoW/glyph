PKG = github.com/k1LoW/glyph
COMMIT = $$(git describe --tags --always)
OSNAME=${shell uname -s}
ifeq ($(OSNAME),Darwin)
	DATE = $$(gdate --utc '+%Y-%m-%d_%H:%M:%S')
else
	DATE = $$(date --utc '+%Y-%m-%d_%H:%M:%S')
endif

BUILD_LDFLAGS = -X $(PKG).commit=$(COMMIT) -X $(PKG).date=$(DATE)

default: test

ci: test

test:
	go test ./... -coverprofile=coverage.out -covermode=count

build:
	go build -ldflags="$(BUILD_LDFLAGS)"

doc:
	go run ./misc/logo/main.go > img/logo.svg
	go run ./misc/coordinates/main.go > img/coordinates.svg
	go run ./misc/database_with_c/main.go > img/database_with_c.svg
	go run ./misc/included/main.go

lint:
	golangci-lint run ./...

depsdev:
	go install github.com/Songmu/ghch/cmd/ghch@latest
	go install github.com/Songmu/gocredits/cmd/gocredits@latest

prerelease:
	git pull origin main --tag
	go mod tidy
	ghch -w -N ${VER}
	gocredits . -w
	git add CHANGELOG.md CREDITS go.mod go.sum
	git commit -m'Bump up version number'
	git tag ${VER}

prerelease_for_tagpr:
	gocredits . -w
	git add CHANGELOG.md CREDITS go.mod go.sum

.PHONY: default test
