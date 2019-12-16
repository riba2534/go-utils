# go-utils

My golang tools.

导入：

```bash
go get -u github.com/riba2534/go-utils/utils
```


## 简单的并发程序

```go
package main

import (
	"github.com/riba2534/go-utils/utils"
)

func main() {
	var wg utils.WaitGroupWrapper
	wg.Wrap(func() {
		// do something..
	})
	wg.Wrap(func() {
		// do somerhing..
	})
	wg.Wait()
}

```