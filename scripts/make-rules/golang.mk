# ==============================================================================
# Makefile helper functions for golang
#

GO := go
GO_SUPPORT_VERSIONS ?= 1.13|1.14|1.15|1.16|1.17
GO_LDFLAGS += -X $(VERSION_PACKAGE).GitVersion=$(VERSION) \
	-X $(VERSION_PACKAGE).GitCommit=$(GIT_COMMIT) \
	-X $(VERSION_PACKAGE).GitTreeState=$(GIT_TREE_STATE) \
	-X $(VERSION_PACKAGE).BuildDate=$(shell date -u + '%Y-%m-%dT%H:%M:%SZ')
