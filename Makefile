.PHONY: g-contest $(PROBS)

PROBS = a b c d e f

$(PROBS):
	@make g/contest/$@

g/%: templates/google-competitions/template.go
	@[ -f $*/main.go ] || (mkdir -p $*; touch $*/main.go; cp $< $*/main.go)

g-contest: $(PROBS)
	@echo "Google contest created! Good luck señor."
