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

## Day01

- The executable must be named ft_otp
- Your program must take arguments.
  1. -g: The program receives as argument a hexadecimal key of at least 64 char-
acters. The program stores this key safely in a file called ft_otp.key, which
is encrypted.
    2. -k: The program generates a new temporary password based on the key given
as argument and prints it on the standard output.
- Your program must use the HOTP algorithm (RFC 4226).
- The generated one-time password must be random and must always contain the
same format, i.e. 6 digits.

## Day02

Setup a `.onion website` using the Tor network.

```bash
# access ssh on the server
ssh -i ~/.ssh/ft-onion root@172.17.0.2 -p 4242
```

[Host your own Tor darkweb .onion site for free with NGINX](https://www.youtube.com/watch?v=6BV-3yWzWcI&t=10s)