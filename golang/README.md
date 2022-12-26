
# Go lang modules


## Local modules
go mod init example.com/hello

> go.mod file
> Need set module name, go version and dependencies
```
module example.com/hello

go 1.19

replace (

)

require (

)
```


>Remote modules
go mod init github.com/username/repo