Package unziper-go
===================

Package unziper-go provides a very simple library to extract zip archive

## Installation
```shell
go get -u github.com/jcarriere/unziper-go
```

## Examples

```go
package main

import (
    "github.com/jcarriere/unziper-gop"
	"fmt"
)

func main() {
	uz := unziper.New("file.zip", "directory/")
	err := uz.Extract()
	if err != nil {
		fmt.Println(err)
	}
}
```

## Contributing

Pull requests, bug fixes and issue reports are welcome.

Before proposing a change, please discuss your change by raising an issue.
