package main

import (

	"github.com/deanishe/awgo"
)

var wf *aw.Workflow

func init() {
	wf = aw.New()
}

func run() {
	wf.NewItem("fafasdfsfsg")
	wf.SendFeedback()
}

func main() {
	wf.Run(run)
}
