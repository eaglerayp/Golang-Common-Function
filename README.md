# IsISODate #

IsISODate is a simple function for checking whether the given string is legal ISO 8601 string.

```go
package main

import (
	"fmt"
	"time"
	"github.com/eaglerayp/Golang-Common-Function"
)

func main() {
	fmt.Println(IsISODate("2012-11-01T22:08:41+00:00"))
	fmt.Println(IsISODate("2016-04-14T01:17:25.980+0000"))
	fmt.Println(IsISODate("2016-04-14T01:17:25.980Z"))
	fmt.Println(IsISODate("2012-11-01T22:08:41+00:00"))
}

```