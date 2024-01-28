# intermittent progress 

```go
package main 
import (
	"time"
	"github.com/intdxdt/interprog"
)

func main(){
    var prog = interprog.NewInterProg("progress")
    for i := 0; i < 10000; i++ {
        time.Sleep(1 * time.Millisecond) //simulate some work
        prog.update() 
    }
    prog.done()
	
}
```