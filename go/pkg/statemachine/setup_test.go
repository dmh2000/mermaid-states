package statemachine

import (
	"os"
	"testing"
)

type testModel struct {
	value int
}

func TestEmpty(t *testing.T) {
	// create the state machine
	sm := NewStateMachine[testModel, int](
		&testModel{0},
		"test",
	)
	if sm.String() != "name: test , states: 0" {
		t.Errorf("unexpected state machine string: %s", sm.String())
	}
}

func TestAddState(t *testing.T) {
	// create the state machine
	sm := NewStateMachine[testModel, int](
		&testModel{0},
		"test",
	)
	if sm.String() != "name: test , states: 0" {
		t.Errorf("unexpected state machine string: %s", sm.String())
	}

	state := NewState(
		"state1",
		func(state *State[testModel, int], model *testModel, input int) (key StateKey, err error) {
			model.value = input
			return "", nil
		},
		nil,
	)

	if err := sm.AddState(state); err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if sm.String() != "name: test , states: 1" {
		t.Errorf("unexpected state machine string: %s", sm.String())
	}

	a := sm.currentState.Key
	if a != s1 {
		t.Errorf("unexpected state string: %s", a)
	}

}

func TestSetInitialState(t *testing.T) {
	// create the state machine
	sm := NewStateMachine[testModel, int](
		&testModel{0},
		"test",
	)
	if sm.String() != "name: test , states: 0" {
		t.Errorf("unexpected state machine string: %s", sm.String())
	}

	// state function
	fstate1 := func(state *State[testModel, int], model *testModel, input int) (key StateKey, err error) {
		model.value = input
		return "", nil
	}

	// add  state1
	state1 := NewState(
		"state1",
		fstate1,
		nil,
	)

	if err := sm.AddState(state1); err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if sm.String() != "name: test , states: 1" {
		t.Errorf("unexpected state machine string: %s", sm.String())
	}

	s := sm.GetCurrentState().String()
	if s != "key: state1" {
		t.Errorf("unexpected state string: %s", s)
	}

	fstate2 := func(state *State[testModel, int], model *testModel, input int) (key StateKey, err error) {
		model.value = input
		return s2, nil
	}
	// add a state
	state2 := NewState(
		"state2",
		fstate2,
		nil,
	)

	if err := sm.AddState(state2); err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if sm.String() != "name: test , states: 2" {
		t.Errorf("unexpected state machine string: %s", sm.String())
	}

	if err := sm.SetInitialState(s2); err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	s = sm.GetCurrentState().String()
	if s != "key: state2" {
		t.Errorf("unexpected state string: %s", s)
	}

}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
