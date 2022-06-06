package main

import (
	aw "github.com/deanishe/awgo"
)

func main() {
	wf := aw.New()
	wf.Run(func() {
		runWithAlfred(wf)
	})

}
