# goimgtype

[![Godoc](http://img.shields.io/badge/godoc-reference-blue.svg?style=flat)](https://godoc.org/github.com/shamsher31/goimgtype)

Return type of image based on mime type

### How to install
```go
go get github.com/shamsher31/goimgtype
```

### How to use
```go
package main

import (
	"fmt"
	"github.com/shamsher31/goimgtype"
)

func main() {
  fmt.Println(imgtype.Get("./golang.jpg"))
  // image/jpeg
}
```

### Related
[goimgext](https://github.com/shamsher31/goimgext)<br>
[govdoext](https://github.com/shamsher31/govdoext)<br>
[gobinext](https://github.com/shamsher31/gobinext)<br>

# License
MIT © [Shamsher Ansari](https://github.com/shamsher31)