#ifndef STATE_MACHINE_H
#define STATE_MACHINE_H

#include <stdint.h>
#include <stdbool.h>

// Opaque types for the state machine implementation
typedef struct state_t state_t;
typedef struct state_machine_t state_machine_t;
typedef int32_t state_key_t;

// Function pointer type for state actions
typedef state_key_t (*action_func_t)(void* model, void* input);

// Create a new state
state_t* state_new(state_key_t key, const char* name, action_func_t action);

// Free a state
void state_free(state_t* state);

// Get state properties
state_key_t state_get_key(const state_t* state);
const char* state_get_name(const state_t* state);

// Execute state action
state_key_t state_execute(state_t* state, void* model, void* input);

// Create a new state machine
state_machine_t* state_machine_new(void* model, const char* name);

// Free a state machine
void state_machine_free(state_machine_t* sm);

// Add a state to the machine
bool state_machine_add_state(state_machine_t* sm, state_t* state);

// Set initial state
bool state_machine_set_initial_state(state_machine_t* sm, state_key_t key);

// Execute current state
state_key_t state_machine_execute(state_machine_t* sm, void* input);

// Get current state
state_t* state_machine_get_current_state(const state_machine_t* sm);

#endif // STATE_MACHINE_H
