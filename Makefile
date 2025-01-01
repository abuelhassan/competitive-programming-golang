PROBS = a b c d e f

$(PROBS):
	@make c/contest/$@

s/%: templates/single/template.go
	@[ -f $*/main.go ] || (mkdir -p $*; touch $*/main.go; cp $< $*/main.go)

c/%: templates/single/template.go
	@[ -f $*/main.go ] || (mkdir -p $*; touch $*/main.go; cp $< $*/main.go)

contest: $(PROBS)
	@echo "Contest created!"
