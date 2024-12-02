#ifndef STATE_MACHINE_H
#define STATE_MACHINE_H

#include <stddef.h>

// Forward declarations
typedef struct State State;
typedef struct StateMachine StateMachine;
typedef struct Model Model;

// Type definitions
typedef const char* StateKey;
typedef StateKey (*ActionFunc)(const State*, Model*, const void*);

// Model structure for testing
struct Model {
    int value;
};

// State structure
struct State {
    StateKey key;
    ActionFunc action;
    void* data;
};

// State Machine structure
struct StateMachine {
    State* current_state;
    State** states;
    size_t num_states;
    size_t capacity;
    const char* name;
};

// Function declarations
StateMachine* new_state_machine(const char* name);
void free_state_machine(StateMachine* sm);

State* new_state(StateKey key, ActionFunc action, void* data);
void free_state(State* state);

int add_state(StateMachine* sm, State* state);
int set_initial_state(StateMachine* sm, StateKey key);
StateKey execute(StateMachine* sm, Model* model, const void* input);

State* get_current_state(const StateMachine* sm);
const char* state_machine_to_string(const StateMachine* sm, char* buffer, size_t size);
const char* state_to_string(const State* state, char* buffer, size_t size);

#endif // STATE_MACHINE_H
