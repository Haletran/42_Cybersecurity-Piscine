# 42_Cybersecurity-Piscine


## GO language

```bash
go mod init <module_name>
go run .
```

```go
// basic hello world in go
package main

import "fmt"

func main() {
    fmt.Println("Hello, World!")
}
```
## Day00

The `spider` program allow you to extract all the images from a website, recursively, by
providing a url as a parameter.

Useful packages for `spider` : 
- `net/http`
- `colly`
- `goquery`


```./spider.go:39:14: more than one character in rune literal```
-> use double quotes instead of single quotes for strings

### TODO 

- [] Need to add more tags check
- [] Add flags for depth and recursive
