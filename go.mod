module github.com/cdvelop/compiler

go 1.20

require (
	github.com/cdvelop/model v0.0.76
	github.com/cdvelop/strings v0.0.7
	github.com/tdewolff/minify v2.3.6+incompatible
)

require (
	github.com/cdvelop/filehandler v0.0.12 // indirect
	github.com/cdvelop/git v0.0.1 // indirect
	github.com/cdvelop/maps v0.0.7 // indirect
	github.com/cdvelop/object v0.0.40 // indirect
	github.com/cdvelop/timetools v0.0.25 // indirect
	github.com/cdvelop/unixid v0.0.25 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/tdewolff/test v1.0.10 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	golang.org/x/net v0.19.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)

require (
	github.com/cdvelop/fileserver v0.0.30
	github.com/cdvelop/gomod v0.0.44
	github.com/cdvelop/gotools v0.0.64
	github.com/cdvelop/input v0.0.59
	github.com/cdvelop/ldflags v0.0.2
	github.com/cdvelop/output v0.0.16
	github.com/cdvelop/platform v0.0.47
	github.com/tdewolff/parse v2.3.4+incompatible // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/ldflags => ../ldflags

replace github.com/cdvelop/js => ../js

replace github.com/cdvelop/gotools => ../gotools

replace github.com/cdvelop/input => ../input

replace github.com/cdvelop/gomod => ../gomod

replace github.com/cdvelop/platform => ../platform

replace github.com/cdvelop/output => ../output

replace github.com/cdvelop/fileserver => ../fileserver

replace github.com/cdvelop/filehandler => ../filehandler

replace github.com/cdvelop/strings => ../strings
