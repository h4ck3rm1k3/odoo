# NOTE: please keep your version of sass up to date: sudo gem update
.PHONY: watch css
SASS_FILES=$(wildcard addons/*/static/src/css/*.sass openerp/addons/*/static/src/css/*.sass)
CSS_FILES=$(patsubst %.sass,%.css,${SASS_FILES})

test2: readmanifest
	./readmanifest

readmanifest : readmanifest.go encoding2/json.o
	gccgo-5 -L. -g -o readmanifest readmanifest.go encoding2/json.o

encoding2/json.o : encoding2/json/decode.go \
	encoding2/json/encode.go \
	encoding2/json/fold.go \
	encoding2/json/indent.go \
	encoding2/json/scanner.go \
	encoding2/json/stream.go \
	encoding2/json/tags.go
	gccgo-5 -g -c -o encoding2/json.o $^



css: ${CSS_FILES}
%.css: %.sass
	sass -t expanded --compass --unix-newlines --sourcemap=none $< $@
watch:
	sass -t expanded --compass --unix-newlines --sourcemap=none --watch .:.


test :
	python convert_go.py

