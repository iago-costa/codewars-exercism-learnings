module example.com/golang

go 1.19


replace (
    example.com/golang => ./
    example.com/reverse_words => ./strings
)

require (
    example.com/golang v0.0.0-00010101000000-000000000000 // indirect
    example.com/reverse_words v0.0.0-00010101000000-000000000000 // indirect
)

