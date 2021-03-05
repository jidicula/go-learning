module hello

go 1.16

replace example.com/greetings => ../greetings

require (
	example.com/greetings v0.0.0-00010101000000-000000000000 // indirect
	rsc.io/quote v1.5.2
)
