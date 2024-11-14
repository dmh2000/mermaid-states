package statemachine

import (
	"testing"
)

// TestStateMachineExecution tests the full execution of the state machine through all states.
func TestStateMachineExecution(t *testing.T) {
	// Create the state machine
	sm := NewStateMachine[testModel, int](
		&testModel{0},
		"test",
	)

	// Define states and their transitions
	state1 := NewState(
		STATE1,
		"state1",
		func(model *testModel, input int) (key StateKey, err error) {
			model.value = input
			return STATE2, nil
		},
	)

	state2 := NewState(
		STATE2,
		"state2",
		func(model *testModel, input int) (key StateKey, err error) {
			model.value = input
			return STATE3, nil
		},
	)

	state3 := NewState(
		STATE3,
		"state3",
		func(model *testModel, input int) (key StateKey, err error) {
			model.value = input
			return STATE4, nil
		},
	)

	state4 := NewState(
		STATE4,
		"state4",
		func(model *testModel, input int) (key StateKey, err error) {
			model.value = input
			return STATE4, nil
		},
	)

	// Add states to the state machine
	if err := sm.AddState(state1); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if err := sm.AddState(state2); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if err := sm.AddState(state3); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}
	if err := sm.AddState(state4); err != nil {
		t.Fatalf("unexpected error: %s", err)
	}

	// Execute the state machine and step through all states
	inputs := []int{10, 20, 30, 40}
	expectedKeys := []StateKey{STATE2, STATE3, STATE4, STATE4} // Correct expected keys
	for i, input := range inputs {
		key, err := sm.Execute(&testModel{0}, input)
		if err != nil {
			t.Fatalf("unexpected error during execution: %s", err)
		}
		if key != expectedKeys[i] {
			t.Errorf("unexpected state key: got %d, want %d", key, expectedKeys[i])
		}
	}
}

// TestNilAction verifies that NewState panics with nil action
func TestNilAction(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Error("expected panic for nil action")
		}
	}()
	NewState[testModel, int](STATE1, "state1", nil)
}

// TestExecuteWithoutInitialState verifies proper error when executing without initial state
func TestExecuteWithoutInitialState(t *testing.T) {
	sm := NewStateMachine[testModel, int](&testModel{0}, "test")
	_, err := sm.Execute(&testModel{0}, 1)
	if err == nil {
		t.Error("expected error when executing without initial state")
	}
}

// TestGetStates verifies the GetStates method
func TestGetStates(t *testing.T) {
	sm := NewStateMachine[testModel, int](&testModel{0}, "test")
	
	states := []struct {
		key  StateKey
		name string
	}{
		{STATE1, "state1"},
		{STATE2, "state2"},
	}

	for _, s := range states {
		state := NewState(s.key, s.name,
			func(model *testModel, input int) (key StateKey, err error) {
				return s.key, nil
			},
		)
		if err := sm.AddState(state); err != nil {
			t.Fatalf("unexpected error: %s", err)
		}
	}

	allStates := sm.GetStates()
	if len(allStates) != len(states) {
		t.Errorf("expected %d states, got %d", len(states), len(allStates))
	}

	for _, s := range states {
		if state, exists := allStates[s.key]; !exists {
			t.Errorf("state %d not found", s.key)
		} else if state.GetName() != s.name {
			t.Errorf("expected name %s, got %s", s.name, state.GetName())
		}
	}
}
