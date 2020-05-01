g/%: templates/google-competitions/template.go
	[ -f $*/main.go ] || (mkdir -p $*; touch $*/main.go; cp $< $*/main.go)
