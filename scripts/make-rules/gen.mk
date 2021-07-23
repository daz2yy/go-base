
# ==============================================================================
# Makefile helper functions for generate necessary files
#

.PHONY: gen.run
#gen.run: gen.errcode gen.docgo
gen.run: gen.clean gen.errcode gen.docgo.doc

.PHONY: gen.errcode
gen.errcode



.PHONY: gen.ca.%
gen.ca.%:
	$(eval CA := $(word 1,$(subst ., ,$*)))
	@echo "==========> Generating CA files for $(CA)"
	@${ROOT_DIR}/scripts/gencerts.sh generate-cert $(OUTPUT_DIR)/cert $(CA)