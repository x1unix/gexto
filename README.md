# gexto

## EXT2/EXT3/EXT4 Filesystem library for Golang

This library is a maintained fork of [github.com/nerd2/gexto](https://github.com/nerd2/gexto) with additional improvements.

#### Introduction

Gexto is a Go library to allow read / write access to EXT2/3/4 filesystems.

#### Minimal Example

Error checking omitted for brevity

```
import (
  "log"
  "github.com/x1unix/gexto"
)

func main() {
  fs, _ := gexto.NewFileSystem("file.ext4")
  
  f, _ := fs.Create("/test")
  f.Write([]byte("hello world")
  f.Close()
  
  g, _ := fs.Open("/another/file")
  log.Println(ioutil.ReadAll(file))
}
```  

#### Testing

Note that testing requires (passwordless) sudo, in order that the test filesystems can be mounted.
  



