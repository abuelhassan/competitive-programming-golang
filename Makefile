g/%: templates/google-competitions/template.go
	@[ -f $*/main.go ] || (mkdir -p $*; touch $*/main.go; cp $< $*/main.go)

PROBS = a b c d e f
.PHONY: g-contest $(PROBS)
$(PROBS):
	@make g/contest/$@
g-contest: $(PROBS)
	@echo "Google contest created! Good luck señor."
