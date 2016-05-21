# goji-gzip

Package `goji-gzip` allows to compress all responses to a certain goji instance. The allowed compression is `gzip`. It is recommended to put this middleware before any other middleware that might show content to the end-user so all responses are gzipped.

### Example usage

```go
package main

import (
    "github.com/patrickdappollonio/goji-gzip"
    "github.com/zenazn/goji"
)

func main() {
    // Enables gzip compression to all routes and responses 
    // under the gzip middleware. 
    goji.Use(gzip.Compress)
    goji.Get("/", myHandler)
    goji.Serve()
}
```
