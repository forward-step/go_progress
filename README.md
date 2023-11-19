# simple

```go
import (
	"time"
	"github.com/forward-step/go_progress/progress"
)

func main() {
	f1 := progress.New()
	p := f1.Add(100)

	for p.Add(10) {
		time.Sleep(time.Millisecond * 100)
	}

	<-f1.Done
	defer close(f1.Done)
}
```

# mult progress

```go
import (
	"time"
	"github.com/forward-step/go_progress/progress"
)

func main() {
	factory := progress.New()

	p1 := factory.Add(100)
	go func() {
		for p1.Add(1) {
			time.Sleep(100 * time.Millisecond)
		}
	}()

	p2 := factory.Add(300)
	go func() {
		for p2.Add(2) {
			time.Sleep(100 * time.Millisecond)
		}
	}()

	p3 := factory.Add(200)
	go func() {
		for p3.Add(3) {
			time.Sleep(100 * time.Millisecond)
		}
	}()

	// wait
	<-factory.Done
}
```
