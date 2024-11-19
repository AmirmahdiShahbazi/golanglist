## A simple golang list generator

```golang
import (
	"fmt"
	"github.com/AmirmahdiShahbazi/golanglist/list"
)

exampleList := list.NewList[int]()
exampleList.Add(1).Add(2).Add(3)
fmt.Println(exampleList.Get())
```
