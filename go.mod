module github.com/cdvelop/compiler

go 1.20

require (
	github.com/cdvelop/model v0.0.39
	github.com/tdewolff/minify v2.3.6+incompatible
)

require (
	github.com/tdewolff/test v1.0.9 // indirect
	golang.org/x/text v0.9.0 // indirect
)

require (
	github.com/cdvelop/gomod v0.0.6
	github.com/cdvelop/gotools v0.0.26
	github.com/cdvelop/input v0.0.21
	github.com/cdvelop/output v0.0.2
	github.com/cdvelop/platform v0.0.4
	github.com/tdewolff/parse v2.3.4+incompatible // indirect
)

replace github.com/cdvelop/model => ../model

replace github.com/cdvelop/js => ../js

replace github.com/cdvelop/gotools => ../gotools

replace github.com/cdvelop/input => ../input

replace github.com/cdvelop/gomod => ../gomod

replace github.com/cdvelop/platform => ../platform

replace github.com/cdvelop/output => ../output
