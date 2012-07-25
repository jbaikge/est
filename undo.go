package main

import (
	"flag"
	"fmt"
)

func init() {
	cmd := &command{
		short: "removes the last annotation from a task",
		long:  "gsafdg",
		usage: "undo <name>",

		needsBackend: true,

		flags: flag.NewFlagSet("undo", flag.ExitOnError),
		run:   undo,
	}

	commands["undo"] = cmd
}

func undo(c *command) {
	args := c.flags.Args()
	if len(args) != 1 {
		c.Usage(1)
	}
	name := args[0]

	task, err := defaultBackend.Load(name)
	if err != nil {
		c.Error(err)
	}

	if err := defaultBackend.PopAnnotation(task); err != nil {
		c.Error(err)
	}

	//get the last annotation and slice it off, and apply its negation.
	anno := task.Annotations[len(task.Annotations)-1]
	task.Annotations = task.Annotations[:len(task.Annotations)-1]
	task.Apply(anno.Negate())

	//print the new data and the removed annotation
	fmt.Println("removed:", anno)
	fmt.Println(task)
}
