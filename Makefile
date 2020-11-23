PKG = github.com/k1LoW/glyph
COMMIT = $$(git describe --tags --always)
OSNAME=${shell uname -s}
ifeq ($(OSNAME),Darwin)
	DATE = $$(gdate --utc '+%Y-%m-%d_%H:%M:%S')
else
	DATE = $$(date --utc '+%Y-%m-%d_%H:%M:%S')
endif

export GO111MODULE=on

BUILD_LDFLAGS = -X $(PKG).commit=$(COMMIT) -X $(PKG).date=$(DATE)

default: test

ci: depsdev test sec

test:
	go test ./... -coverprofile=coverage.txt -covermode=count

sec:
	gosec ./...

build:
	go build -ldflags="$(BUILD_LDFLAGS)"

doc:
	go run ./misc/coordinates/main.go > img/coordinates.svg
	go run ./misc/database_with_c/main.go > img/database_with_c.svg
	go run ./misc/included/main.go

ci_doc: doc
	$(eval DIFF_EXIST := $(shell git checkout go.* && git diff --exit-code --quiet || echo "exist"))
	test -z "$(DIFF_EXIST)" || (git add -A ./img && git add -A *.md && git commit -m "Update images by GitHub Action (${GITHUB_SHA})" && git push -v origin ${GITHUB_BRANCH})

depsdev:
	go get github.com/Songmu/ghch/cmd/ghch
	go get github.com/Songmu/gocredits/cmd/gocredits
	go get github.com/securego/gosec/cmd/gosec

prerelease:
	git pull origin --tag
	ghch -w -N ${VER}
	gocredits . > CREDITS
	git add CHANGELOG.md CREDITS
	git commit -m'Bump up version number'
	git tag ${VER}

.PHONY: default test
