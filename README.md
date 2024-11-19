## A simple golang list generator

```golang
import (
	"fmt"
	"github.com/AmirmahdiShahbazi/golanglist/list"
)

var list = list.NewList[int]()
list.Add(1).Add(2).Add(3)
fmt.Println(list.Get())
```