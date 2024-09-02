# Pointer

A tool, that helps to create pointers on the fly.

For APIs, tests, etc

Usage:

```go

import "github.com/spirin/pointer"

type Bar{}

func main() {
    pObj := pointer.Of(Bar{})
    
    pString := pointer.Of("some string")

    pIntA := pointer.Of(1)
    pIntB := pointer.Of(2)
    
    if pointer.Equal(pIntA, pIntB){
        ...
    }

    if pointer.EqualValue(pIntA, 123){
        ...
    }
    
    fmt.Println(pointer.Value(pIntB, 123))
    
    // types that cannot be autoresolved
    
    pInt64 := pointer.Int64(1)
    
    pFloat16 := pointer.Float16(1)
    
    ...
}
```
