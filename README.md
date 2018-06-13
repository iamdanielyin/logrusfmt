# logrusfmt

Logs time and message on the same line.

Example:

```go
package main

import (
	log "github.com/sirupsen/logrus"
	"github.com/yinfxs/logrusfmt"
)

func main() {
	log.SetFormatter(&SimpleTextFormatter{})
    log.WithFields(log.Fields{
        "animal": "walrus",
    }).Info("A walrus appears")
```
