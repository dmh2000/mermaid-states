package build

import (
	"testing"

	sm "sqirvy.xyz/state-gen/pkg/statemachine"
)

var statekeys = []sm.StateKey{
	"1",
	"2",
	"3",
	"END",
}

type XModel struct {
	value int
}
type XInput int

var statex = []*sm.State[XModel, XInput]{
	sm.NewState(
		statekeys[0],
		func(current *sm.State[XModel, XInput], model *XModel, input XInput) (key sm.StateKey, err error) {
			// COMMENT
			model.value++
			return statekeys[2], nil
		},
		nil,
	),
	sm.NewState(
		statekeys[1],
		func(currentState *sm.State[XModel, XInput], model *XModel, input XInput) (key sm.StateKey, err error) {
			// COMMENT
			model.value++
			return statekeys[3], nil
		},
		nil,
	),
	sm.NewState(
		statekeys[2],
		func(currentState *sm.State[XModel, XInput], model *XModel, input XInput) (key sm.StateKey, err error) {
			// COMMENT
			model.value++
			return statekeys[3], nil
		},
		nil,
	),
	sm.NewState(
		statekeys[3],
		func(currentState *sm.State[XModel, XInput], model *XModel, input XInput) (key sm.StateKey, err error) {
			return statekeys[3], nil
		},
		nil,
	),
}

func TestStateMachine(t *testing.T) {

	sm := sm.NewStateMachine[XModel, XInput](&XModel{}, "test")
	for _, s := range statex {
		if err := sm.AddState(s); err != nil {
			t.Errorf("unexpected error: %s", err)
		}
	}

	for {
		key, err := sm.Execute(&XModel{}, XInput(1))
		if err != nil {
			t.Fatalf("unexpected error during execution: %s", err)
		}
		if key == "END" {
			break
		}
	}

}
