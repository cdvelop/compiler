module github.com/cdvelop/compiler

go 1.20

require (
	github.com/cdvelop/model v0.0.73
	github.com/cdvelop/strings v0.0.7
	github.com/tdewolff/minify v2.3.6+incompatible
)

require (
	github.com/cdvelop/filehandler v0.0.8 // indirect
	github.com/cdvelop/maps v0.0.7 // indirect
	github.com/cdvelop/object v0.0.35 // indirect
	github.com/cdvelop/timetools v0.0.21 // indirect
	github.com/cdvelop/unixid v0.0.21 // indirect
	github.com/gabriel-vasile/mimetype v1.4.3 // indirect
	github.com/tdewolff/test v1.0.10 // indirect
	github.com/x448/float16 v0.8.4 // indirect
	golang.org/x/net v0.18.0 // indirect
	golang.org/x/text v0.14.0 // indirect
)

require (
	github.com/cdvelop/fileserver v0.0.27
	github.com/cdvelop/gomod v0.0.39
	github.com/cdvelop/gotools v0.0.60
	github.com/cdvelop/input v0.0.55
	github.com/cdvelop/output v0.0.16
	github.com/cdvelop/platform v0.0.43
	github.com/tdewolff/parse v2.3.4+incompatible // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/js => ../js

replace github.com/cdvelop/gotools => ../gotools

replace github.com/cdvelop/input => ../input

replace github.com/cdvelop/gomod => ../gomod

replace github.com/cdvelop/platform => ../platform

replace github.com/cdvelop/output => ../output

replace github.com/cdvelop/fileserver => ../fileserver

replace github.com/cdvelop/filehandler => ../filehandler

replace github.com/cdvelop/strings => ../strings
