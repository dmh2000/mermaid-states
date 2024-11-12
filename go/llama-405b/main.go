package main

import (
	"fmt"
)

func main() {
	sm := NewStateMachine()

	sm.AddState("start")
	sm.AddState("running")
	sm.AddState("paused")
	sm.AddState("stopped")

	sm.AddTransition("start", "run", "running", func() error {
		fmt.Println("Starting the machine...")
		return nil
	})
	sm.AddTransition("running", "pause", "paused", func() error {
		fmt.Println("Pausing the machine...")
		return nil
	})
	sm.AddTransition("paused", "resume", "running", func() error {
		fmt.Println("Resuming the machine...")
		return nil
	})
	sm.AddTransition("running", "stop", "stopped", func() error {
		fmt.Println("Stopping the machine...")
		return nil
	})

	sm.SetInitialState("start")

	fmt.Println("Current State:", sm.CurrentState())

	sm.TriggerEvent("run")
	fmt.Println("Current State:", sm.CurrentState())

	sm.TriggerEvent("pause")
	fmt.Println("Current State:", sm.CurrentState())

	sm.TriggerEvent("resume")
	fmt.Println("Current State:", sm.CurrentState())

	sm.TriggerEvent("stop")
	fmt.Println("Current State:", sm.CurrentState())
}
