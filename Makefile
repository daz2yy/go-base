
# Build all by default, even if it's not first
.DEFAULT_GOAL := all

.PHONY: all
all: gen format lint cover build

# ==============================================================================
# Build options

ROOT_PACKAGE=github.com/daz2yy/ctl
VERSION_PACKAGE=github.com/daz2yy/go-base/pkg/version

# ==============================================================================
# Includes


# ==============================================================================
# Usage



# ==============================================================================
# Targets
