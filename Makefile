# NOTE: please keep your version of sass up to date: sudo gem update
.PHONY: watch css
SASS_FILES=$(wildcard addons/*/static/src/css/*.sass openerp/addons/*/static/src/css/*.sass)
CSS_FILES=$(patsubst %.sass,%.css,${SASS_FILES})

test2: readmanifest
	./readmanifest

xsd.o : ./docs.oasis-open.org/election/external/xAL.xsd_go/xAL.xsd.go ./kbcafe.com/rss/atom.xsd.xml_go/atom.xsd.xml.go ./schemas.opengis.net/kml/2.2.0/ogckml22.xsd_go/ogckml22.xsd.go ./schemas.opengis.net/kml/2.2.0/atom-author-link.xsd_go/atom-author-link.xsd.go ./docbook.org/xml/5.0/xsd/xlink.xsd_go/xlink.xsd.go  ./docbook.org/xml/5.0/xsd/xml.xsd_go/xml.xsd.go ./docbook.org/xml/5.0/xsd/docbook.xsd_go/docbook.xsd.go  ./khronos.org/files/collada_schema_1_5_go/collada_schema_1_5.go  ./khronos.org/files/collada_schema_1_4_go/collada_schema_1_4.go ./thearchitect.co.uk/schemas/rss-2_0.xsd_go/rss-2_0.xsd.go ./www.w3.org/2001/03/xml.xsd_go/xml.xsd.go ./www.w3.org/2001/xml.xsd_go/xml.xsd.go ./www.w3.org/2007/schema-for-xslt20.xsd_go/schema-for-xslt20.xsd.go  ./www.w3.org/TR/2002/WD-SVG11-20020108/xlink.xsd_go/xlink.xsd.go ./www.w3.org/TR/2002/WD-SVG11-20020108/xml.xsd_go/xml.xsd.go  ./www.w3.org/TR/2002/WD-SVG11-20020108/SVG.xsd_go/SVG.xsd.go  ./www.w3.org/Math/XMLSchema/mathml2/common/xlink-href.xsd_go/xlink-href.xsd.go ./www.w3.org/Math/XMLSchema/mathml2/mathml2.xsd_go/mathml2.xsd.go
	gccgo-5 -L. -c -g -o xsd.o $^

readmanifest : readmanifest.go encoding2/json.o readdatafiles.go encoding2/xml.o xsd.o
	gccgo-5 -L. -g -o readmanifest readmanifest.go encoding2/json.o readdatafiles.go encoding2/xml.o xsd.o

encoding2/json.o : encoding2/json/decode.go \
	encoding2/json/encode.go \
	encoding2/json/fold.go \
	encoding2/json/indent.go \
	encoding2/json/scanner.go \
	encoding2/json/stream.go \
	encoding2/json/tags.go
	gccgo-5 -g -c -o encoding2/json.o $^

encoding2/xml.o : encoding2/xml/marshal.go encoding2/xml/read.go encoding2/xml/typeinfo.go encoding2/xml/xml.go
	gccgo-5 -g -c -o encoding2/xml.o $^



css: ${CSS_FILES}
%.css: %.sass
	sass -t expanded --compass --unix-newlines --sourcemap=none $< $@
watch:
	sass -t expanded --compass --unix-newlines --sourcemap=none --watch .:.


test :
	python convert_go.py

