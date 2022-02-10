module philcrockett.com/hello

go 1.17

replace philcrockett.com/greetings => ../greetings

require philcrockett.com/greetings v0.0.0-00010101000000-000000000000

require (
	github.com/alecthomas/kong v0.4.0 // indirect
	github.com/pkg/errors v0.9.1 // indirect
)
