module github.com/rebbyv/go-module-tutorial.git/hello

go 1.20

replace github.com/rebbyv/go-module-tutorial.git/greetings => ../greetings

require github.com/rebbyv/go-module-tutorial.git/greetings v0.0.0-00010101000000-000000000000
