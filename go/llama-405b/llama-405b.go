package main

import (
	"fmt"
)

// State represents a state in the state machine.
type State string

// Event represents an event that triggers a transition.
type Event string

// Action is a function that performs an action on state entry, exit, or transition.
type Action func() error

// Transition represents a transition between two states.
type Transition struct {
	From  State
	Event Event
	To    State
	Action
}

// StateMachine is a simple finite state machine.
type StateMachine struct {
	currentState State
	states       map[State]struct{}
	transitions  map[State]map[Event]Transition
}

// NewStateMachine returns a new state machine.
func NewStateMachine() *StateMachine {
	return &StateMachine{
		states:      make(map[State]struct{}),
		transitions: make(map[State]map[Event]Transition),
	}
}

// AddState adds a new state to the state machine.
func (sm *StateMachine) AddState(state State) {
	sm.states[state] = struct{}{}
	if _, ok := sm.transitions[state]; !ok {
		sm.transitions[state] = make(map[Event]Transition)
	}
}

// AddTransition adds a new transition to the state machine.
func (sm *StateMachine) AddTransition(from State, event Event, to State, action Action) error {
	if _, ok := sm.states[from]; !ok {
		return fmt.Errorf("state %s does not exist", from)
	}
	if _, ok := sm.states[to]; !ok {
		return fmt.Errorf("state %s does not exist", to)
	}
	sm.transitions[from][event] = Transition{
		From:   from,
		Event:  event,
		To:     to,
		Action: action,
	}
	return nil
}

// TriggerEvent triggers an event in the current state.
func (sm *StateMachine) TriggerEvent(event Event) error {
	transitions, ok := sm.transitions[sm.currentState]
	if !ok {
		return fmt.Errorf("no transitions defined for state %s", sm.currentState)
	}
	transition, ok := transitions[event]
	if !ok {
		return fmt.Errorf("no transition defined for event %s in state %s", event, sm.currentState)
	}
	if transition.Action != nil {
		if err := transition.Action(); err != nil {
			return err
		}
	}
	sm.currentState = transition.To
	return nil
}

// SetInitialState sets the initial state of the state machine.
func (sm *StateMachine) SetInitialState(state State) error {
	if _, ok := sm.states[state]; !ok {
		return fmt.Errorf("state %s does not exist", state)
	}
	sm.currentState = state
	return nil
}

// CurrentState returns the current state of the state machine.
func (sm *StateMachine) CurrentState() State {
	return sm.currentState
}
