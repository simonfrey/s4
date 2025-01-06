# Base24 encoder/decoder for Go

This library is an encoder/decoder for the [Base24 encoding](https://www.kuon.ch/post/2020-02-27-base24/) for the go programming language.

## Usage

Install

```bash
go get gitlab.com/phil9909/base24
```

Encode

```go
import "gitlab.com/phil9909/base24"

func main() {
  data, err := base24.StdEncoding.EncodeToString(/* some []byte array */)
  if err != nil {
    // ...
  }
  // ... 
}
```

Decode

```go
import "gitlab.com/phil9909/base24"

func main() {
  data, err := base24.StdEncoding.DecodeString("ZZZZZZZ")
  if err != nil {
    // ...
  }
  // ... 
}
```

## License

Licensed under either of

 * Apache License, Version 2.0
   ([LICENSE-APACHE](LICENSE-APACHE) or http://www.apache.org/licenses/LICENSE-2.0)
 * MIT license
   ([LICENSE-MIT](LICENSE-MIT) or http://opensource.org/licenses/MIT)

at your option.

## Acknowledgment

The content of this repos is heavily inspiered by:

* The Kotlin implementation of base24: https://github.com/kuon/java-base24
  
  Especially the:
  * Tests cases
  * README

* The go standardlib implmentation of base64: https://github.com/golang/go/blob/master/src/encoding/base64/base64.go

  Especially the:
  * Interface
  * The constructor (mapping the alphabet to the decoding map)
