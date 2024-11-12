// Package statemachine provides a simple state machine implementation
package statemachine

import (
	"fmt"
)

type StateKey int32

func (k StateKey) String() string {
	return fmt.Sprintf("%d", k)
}

type State[Model any, Input any] struct {
	key    StateKey
	name   string
	action func(*Model, Input) (key StateKey, err error)
}

func (s *State[Model, Input]) String() string {
	return fmt.Sprintf("key: %d , name: %s", s.key, s.name)
}

func (s *State[Model, Input]) GetKey() StateKey {
	return s.key
}

func (s *State[Model, Input]) GetName() string {
	return s.name
}

// Execute :  the state action and return a key to the next state (could be the same or different)
func (s *State[Model, Input]) Execute(model *Model, input Input) (key StateKey, err error) {
	return s.action(model, input)
}

func NewState[Model any, Input any](
	key StateKey,
	name string,
	act func(*Model, Input) (key StateKey, err error),
) *State[Model, Input] {
	return &State[Model, Input]{
		key:    key,
		name:   name,
		action: act,
	}
}

// ============================================================================

type StateMachine[Model any, Input any] struct {
	currentState *State[Model, Input]
	states       map[StateKey]*State[Model, Input]
	name         string
}

func NewStateMachine[Model any, Input any](model *Model, name string) *StateMachine[Model, Input] {
	return &StateMachine[Model, Input]{
		currentState: nil,
		states:       make(map[StateKey]*State[Model, Input]),
		name:         name,
	}
}

func (sm *StateMachine[Model, Input]) String() string {
	return fmt.Sprintf("name: %s , states: %d", sm.name, len(sm.states))
}

func (sm *StateMachine[Model, Input]) GetCurrentState() *State[Model, Input] {
	return sm.currentState
}

func (sm *StateMachine[Model, Input]) AddState(state *State[Model, Input]) error {
	// check if state already exists
	if _, exists := sm.states[state.key]; exists {
		return fmt.Errorf("state %d already exists", state.key)
	}
	// ok, add it to the map
	sm.states[state.key] = state

	// if there is no current state, set it as the initial state
	if sm.currentState == nil {
		if err := sm.SetInitialState(state.key); err != nil {
			return err
		}
	}
	return nil
}

func (sm *StateMachine[Model, Input]) SetInitialState(key StateKey) error {

	state, exists := sm.states[key]
	if !exists {
		return fmt.Errorf("state %d does not exist", key)
	}

	sm.currentState = state

	return nil
}

func (sm *StateMachine[Model, Input]) Execute(model *Model, input Input) (key StateKey, err error) {
	key, err = sm.currentState.Execute(model, input)
	if err != nil {
		return
	}

	// same state, no change
	if key == sm.currentState.GetKey() {
		return key, nil
	}

	// new state
	newState, exists := sm.states[key]
	if !exists {
		return 0, fmt.Errorf("state %d does not exist", key)
	}

	// set next state
	sm.currentState = newState

	return newState.GetKey(), nil
}
