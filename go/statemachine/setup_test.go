package statemachine

import (
	"os"
	"testing"
)

const (
	INVALID = iota
	STATE1
	STATE2
	STATE3
	STATE4
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

	// add a state
	state := NewState(
		STATE1,
		"state1",
		func(model *testModel, input int) (key StateKey, err error) {
			model.value = input
			return 2, nil
		},
	)

	if err := sm.AddState(state); err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if sm.String() != "name: test , states: 1" {
		t.Errorf("unexpected state machine string: %s", sm.String())
	}

	s := sm.currentState.String()
	if s != "key: 1 , name: state1" {
		t.Errorf("unexpected state string: %s", s)
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

	// add a state1
	state1 := NewState(
		STATE1,
		"state1",
		func(model *testModel, input int) (key StateKey, err error) {
			model.value = input
			return 2, nil
		},
	)

	if err := sm.AddState(state1); err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if sm.String() != "name: test , states: 1" {
		t.Errorf("unexpected state machine string: %s", sm.String())
	}

	s := sm.GetCurrentState().String()
	if s != "key: 1 , name: state1" {
		t.Errorf("unexpected state string: %s", s)
	}

	// add a state
	state2 := NewState(
		STATE2,
		"state2",
		func(model *testModel, input int) (key StateKey, err error) {
			model.value = input
			return 2, nil
		},
	)

	if err := sm.AddState(state2); err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	if sm.String() != "name: test , states: 2" {
		t.Errorf("unexpected state machine string: %s", sm.String())
	}

	if err := sm.SetInitialState(state2.GetKey()); err != nil {
		t.Errorf("unexpected error: %s", err)
	}

	s = sm.GetCurrentState().String()
	if s != "key: 2 , name: state2" {
		t.Errorf("unexpected state string: %s", s)
	}

}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
