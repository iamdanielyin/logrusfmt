# logrusfmt

Logs time and message on the same line.

Example:

```go
package main

import (
	log "github.com/sirupsen/logrus"
	fmt "github.com/yinfxs/logrusfmt"
)

func main() {
	log.SetFormatter(&fmt.SimpleTextFormatter{})
	log.WithFields(log.Fields{
		"animal": "walrus",
	}).Info("A walrus appears")
}
```
