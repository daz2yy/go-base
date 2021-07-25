
# Build all by default, even if it's not first
.DEFAULT_GOAL := all

.PHONY: all
all: gen format lint cover build
# all: build

# ==============================================================================
# Build options

ROOT_PACKAGE=github.com/daz2yy/go-base
VERSION_PACKAGE=github.com/daz2yy/go-base/pkg/version

# ==============================================================================
# Includes

include scripts/make-rules/common.mk # make sure include common.mk at the first include line
include scripts/make-rules/golang.mk
include scripts/make-rules/gen.mk
include scripts/make-rules/tools.mk
# include scripts/make-rules/image.mk
# include scripts/make-rules/deploy.mk
# include scripts/make-rules/copyright.mk
# include scripts/make-rules/ca.mk
# include scripts/make-rules/release.mk
# include scripts/make-rules/swagger.mk
# include scripts/make-rules/dependencies.mk

# ==============================================================================
# Usage

define USAGE_OPTIONS

Options:
  DEBUG        Whether to generate debug symbols. Default is 0.
  BINS         The binaries to build. Default is all of cmd.
               This option is available when using: make build/build.multiarch
               Example: make build BINS="iam-apiserver iam-authz-server"
  IMAGES       Backend images to make. Default is all of cmd starting with iam-.
               This option is available when using: make image/image.multiarch/push/push.multiarch
               Example: make image.multiarch IMAGES="iam-apiserver iam-authz-server"
  PLATFORMS    The multiple platforms to build. Default is linux_amd64 and linux_arm64.
               This option is available when using: make build.multiarch/image.multiarch/push.multiarch
               Example: make image.multiarch IMAGES="iam-apiserver iam-pump" PLATFORMS="linux_amd64 linux_arm64"
  VERSION      The version information compiled into binaries.
               The default is obtained from gsemver or git.
  V            Set to 1 enable verbose build. Default is 0.
endef
export USAGE_OPTIONS



# ==============================================================================
# Targets

## help: Show this help info
.PHONY: help
help: Makefile
	@echo -e "\nUsage: make <TARGETS> <OPTIONS> ...\n\nTargets:"
	@sed -n 's/^##//p' $< | column -t -s ':' | sed -e 's/^/ /'
	@echo "$$USAGE_OPTIONS"

## gen: Generate all necessary files, such as error code files.
.PHONY: gen
gen:
	@$(MAKE) gen.run

## format: Gofmt (reformat) package sources (exclude vendor dir if existed).
.PHONY: format
format: tools.verify.golines tools.verify.goimports
	@echo "==========> Formatting codes"
	@$(FIND) -type f -name '*.go' | $(XARGS) gofmt -s -w
	@$(FIND) -type f -name '*.go' | $(XARGS) goimports -w -local $(ROOT_PACKAGE)
	@$(FIND) -type f -name '*.go' | $(XARGS) golines -w --max-len=120 --reformat-tags --shorten-comments --ignore-generated
	@$(GO) mod edit -fmt

## lint: Check syntax and styling of go sources.
.PHONY: lint
lint:
	@$(MAKE) go.lint

## cover: Run unit test and get test coverage.
.PHONY: cover
cover:
	@$(MAKE) go.test.cover

## build: Build source code for host platform.
.PHONY: build
build: 
	@$(MAKE) go.build

## build.multiarch: Build source code for multiple platforms. See option PLATFORMS.
.PHONY: build.multiarch
build.multiarch:
	@$(MAKE) go.build.multiarch



