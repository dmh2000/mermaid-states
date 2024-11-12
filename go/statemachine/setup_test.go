package statemachine

import (
	"os"
	"testing"
)

const (
	State1 = 1
	State2 = 2
	State3 = 3
	State4 = 4
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
	if sm.String() != "StateMachine test , 0" {
		t.Errorf("unexpected state machine string: %s", sm.String())
	}
}

func TestAddState(t *testing.T) {
	// create the state machine
	sm := NewStateMachine[testModel, int](
		&testModel{0},
		"test",
	)
	if sm.String() != "StateMachine test , 0" {
		t.Errorf("unexpected state machine string: %s", sm.String())
	}

	// add a state
	state := NewState[testModel, int](
		State1,
		"state1",
		func(model *testModel, input int) (key int32, err error) {
			model.value = input
			return 2, nil
		},
	)

	if err := sm.AddState(state); err != nil {
		t.Errorf("unexpected error: %s", err)
	}
	if sm.String() != "StateMachine test , 1" {
		t.Errorf("unexpected state machine string: %s", sm.String())
	}
}

func TestMain(m *testing.M) {
	os.Exit(m.Run())
}
