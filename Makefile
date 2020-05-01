g/%: templates/google-competitions/template.go
	mkdir -p $*
	touch -p $*/main.go
	cp $< $*/main.go
