package main

import (
	"fmt"
	"log"
)

func main() {
	// Create a new state machine
	sm := NewStateMachine()

	// Add states
	sm.AddState("idle")
	sm.AddState("running")
	sm.AddState("paused")
	sm.AddState("completed")

	// Add transitions with actions
	sm.AddTransition("idle", "running", "start", func() error {
		fmt.Println("Starting the process...")
		return nil
	})

	sm.AddTransition("running", "paused", "pause", func() error {
		fmt.Println("Pausing the process...")
		return nil
	})

	sm.AddTransition("paused", "running", "resume", func() error {
		fmt.Println("Resuming the process...")
		return nil
	})

	sm.AddTransition("running", "completed", "finish", func() error {
		fmt.Println("Finishing the process...")
		return nil
	})

	// Set initial state
	if err := sm.SetInitialState("idle"); err != nil {
		log.Fatal(err)
	}

	// Example usage
	fmt.Printf("Current state: %s\n", sm.CurrentState())

	transitions := []string{"start", "pause", "resume", "finish"}
	
	for _, event := range transitions {
		fmt.Printf("\nTriggering event: %s\n", event)
		if err := sm.Trigger(event); err != nil {
			log.Printf("Error: %v\n", err)
			continue
		}
		fmt.Printf("Current state: %s\n", sm.CurrentState())
	}
}
