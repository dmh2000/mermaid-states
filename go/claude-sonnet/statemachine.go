package main

// State represents a state in the state machine
type State struct {
	Name        string
	Transitions map[string]*Transition
}

// Transition represents a transition between states
type Transition struct {
	From      string
	To        string
	Event     string
	Action    func() error
}

// StateMachine represents the finite state machine
type StateMachine struct {
	states       map[string]*State
	currentState *State
}

// NewStateMachine creates a new state machine
func NewStateMachine() *StateMachine {
	return &StateMachine{
		states: make(map[string]*State),
	}
}

// AddState adds a new state to the machine
func (sm *StateMachine) AddState(name string) *State {
	state := &State{
		Name:        name,
		Transitions: make(map[string]*Transition),
	}
	sm.states[name] = state
	return state
}

// AddTransition adds a transition between states
func (sm *StateMachine) AddTransition(from, to, event string, action func() error) error {
	fromState, exists := sm.states[from]
	if !exists {
		return fmt.Errorf("state '%s' does not exist", from)
	}

	if _, exists := sm.states[to]; !exists {
		return fmt.Errorf("state '%s' does not exist", to)
	}

	transition := &Transition{
		From:   from,
		To:     to,
		Event:  event,
		Action: action,
	}

	fromState.Transitions[event] = transition
	return nil
}

// SetInitialState sets the initial state of the machine
func (sm *StateMachine) SetInitialState(name string) error {
	state, exists := sm.states[name]
	if !exists {
		return fmt.Errorf("state '%s' does not exist", name)
	}
	sm.currentState = state
	return nil
}

// Trigger handles an event and performs the state transition
func (sm *StateMachine) Trigger(event string) error {
	if sm.currentState == nil {
		return fmt.Errorf("no current state set")
	}

	transition, exists := sm.currentState.Transitions[event]
	if !exists {
		return fmt.Errorf("no transition for event '%s' from state '%s'", 
			event, sm.currentState.Name)
	}

	if transition.Action != nil {
		if err := transition.Action(); err != nil {
			return err
		}
	}

	sm.currentState = sm.states[transition.To]
	return nil
}

// CurrentState returns the name of the current state
func (sm *StateMachine) CurrentState() string {
	if sm.currentState == nil {
		return ""
	}
	return sm.currentState.Name
}
