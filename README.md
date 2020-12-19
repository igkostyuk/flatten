# flatten
Package for flatten slice matrix to slice.

## Clockwise
Flatten slice square matrix to slice in clockwise order.

### Example
    package main

    import (
	    "fmt"

	    "github.com/igkostyuk/flatten"
    )

    func main() {
	    a := [][]int{
		    {1, 2, 3},
		    {8, 9, 4},
		    {7, 6, 5}}

	    b, _ := flatten.Clockwise(a)
	    fmt.Printf("%v is %v after flatten", b, a)
    }
