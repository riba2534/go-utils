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
