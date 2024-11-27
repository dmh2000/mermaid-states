package statemachine

import (
	"testing"
)

const (
	s1 StateKey = "state1"
	s2 StateKey = "state2"
	s3 StateKey = "state3"
	s4 StateKey = "state4"
)

func fstate1(s *State[testModel, int], model *testModel, input int) (key StateKey, err error) {
	model.value = input
	return s2, nil
}

func fstate2(s *State[testModel, int], model *testModel, input int) (key StateKey, err error) {
	model.value = input
	return s3, nil
}

func fstate3(s *State[testModel, int], model *testModel, input int) (key StateKey, err error) {
	model.value = input
	return s4, nil
}

func fstate4(s *State[testModel, int], model *testModel, input int) (key StateKey, err error) {
	model.value = input
	return s1, nil
}

// TestStateMachineExecution tests the full execution of the state machine through all states.
func TestStateMachineExecution(t *testing.T) {
	// Create the state machine
	sm := NewStateMachine[testModel, int](
		&testModel{0},
		"test",
	)

	// Define states and their transitions
	state1 := NewState(
		s1,
		fstate1,
		nil,
	)

	state2 := NewState(
		s2,
		fstate2,
		nil,
	)

	state3 := NewState(
		s3,
		fstate3,
		nil,
	)

	state4 := NewState(
		s4,
		fstate4,
		nil,
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
	expectedKeys := []StateKey{s2, s3, s4, s1}
	for i, input := range inputs {
		key, err := sm.Execute(&testModel{0}, input)
		if err != nil {
			t.Fatalf("unexpected error during execution: %s", err)
		}
		if key != expectedKeys[i] {
			t.Errorf("unexpected state key: got %v, want %v", key, expectedKeys[i])
		}
	}
}

// TestNilAction verifies that NewState returns nil if the key is an empty stringor action is nil
func TestNilAction(t *testing.T) {
	s := NewState[any, any]("", nil, nil)
	if s != nil {
		t.Errorf("expected nil state, got %v", s)
	}
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
		key StateKey
	}{
		{s1},
		{s2},
	}

	for _, s := range states {
		state := NewState(s.key,
			func(x *State[testModel, int], model *testModel, input int) (key StateKey, err error) {
				return x.Key, nil
			},
			nil,
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
			t.Errorf("state %v not found", s.key)
		} else if state.GetKey() != s.key {
			t.Errorf("expected key %s, got %s", s.key, state.GetKey())
		}
	}
}
