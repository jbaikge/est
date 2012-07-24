package main

import (
	"flag"
	"fmt"
	"time"
)

func init() {
	cmd := &command{
		short: "starts working on a task",
		long:  "foob",
		usage: "start <task name>",

		needsBackend: true,

		flags: flag.NewFlagSet("start", flag.ExitOnError),
		run:   start,
	}

	commands["start"] = cmd
}

func start(c *command) {
	args := c.flags.Args()
	if len(args) != 1 {
		c.Usage(1)
	}

	log, err := defaultBackend.Status()
	if err != nil {
		c.Error(err)
	}
	if log != nil {
		fmt.Println("already working on", log.Name)
		if err := defaultBackend.Stop(); err != nil {
			c.Error(err)
		}
		dur := time.Since(log.When)
		fmt.Println("adding", dur, "to", log.Name)

		ann := Annotation{
			When:        time.Now(),
			ActualDelta: dur,
		}
		if err := defaultBackend.AddAnnotation(log.Name, ann); err != nil {
			c.Error(err)
		}
	}

	task, err := defaultBackend.Load(args[0])
	if err != nil {
		c.Error(err)
	}
	if err := defaultBackend.Start(task.Name); err != nil {
		c.Error(err)
	}

	fmt.Println("started working on", task.Name)
	fmt.Println(task)
}
