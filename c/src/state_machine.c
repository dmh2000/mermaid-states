#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include "state_machine.h"

#define INITIAL_CAPACITY 10

StateMachine* new_state_machine(const char* name) {
    StateMachine* sm = malloc(sizeof(StateMachine));
    if (!sm) return NULL;

    sm->states = malloc(sizeof(State*) * INITIAL_CAPACITY);
    if (!sm->states) {
        free(sm);
        return NULL;
    }

    sm->current_state = NULL;
    sm->num_states = 0;
    sm->capacity = INITIAL_CAPACITY;
    sm->name = strdup(name);
    
    return sm;
}

void free_state_machine(StateMachine* sm) {
    if (!sm) return;
    
    for (size_t i = 0; i < sm->num_states; i++) {
        free_state(sm->states[i]);
    }
    
    free(sm->states);
    free((void*)sm->name);
    free(sm);
}

State* new_state(StateKey key, ActionFunc action, void* data) {
    if (!key || !action) return NULL;
    
    State* state = malloc(sizeof(State));
    if (!state) return NULL;
    
    state->key = strdup(key);
    state->action = action;
    state->data = data;
    
    return state;
}

void free_state(State* state) {
    if (!state) return;
    free((void*)state->key);
    free(state);
}

static State* find_state(const StateMachine* sm, StateKey key) {
    for (size_t i = 0; i < sm->num_states; i++) {
        if (strcmp(sm->states[i]->key, key) == 0) {
            return sm->states[i];
        }
    }
    return NULL;
}

int add_state(StateMachine* sm, State* state) {
    if (!sm || !state) return -1;
    
    if (find_state(sm, state->key)) {
        return -1;  // State already exists
    }
    
    if (sm->num_states >= sm->capacity) {
        size_t new_capacity = sm->capacity * 2;
        State** new_states = realloc(sm->states, sizeof(State*) * new_capacity);
        if (!new_states) return -1;
        
        sm->states = new_states;
        sm->capacity = new_capacity;
    }
    
    sm->states[sm->num_states++] = state;
    
    if (!sm->current_state) {
        sm->current_state = state;
    }
    
    return 0;
}

int set_initial_state(StateMachine* sm, StateKey key) {
    if (!sm) return -1;
    
    State* state = find_state(sm, key);
    if (!state) return -1;
    
    sm->current_state = state;
    return 0;
}

StateKey execute(StateMachine* sm, Model* model, const void* input) {
    if (!sm || !sm->current_state) return NULL;
    
    StateKey next_key = sm->current_state->action(sm->current_state, model, input);
    if (!next_key) return NULL;
    
    if (strcmp(next_key, sm->current_state->key) != 0) {
        State* next_state = find_state(sm, next_key);
        if (!next_state) return NULL;
        sm->current_state = next_state;
    }
    
    return sm->current_state->key;
}

State* get_current_state(const StateMachine* sm) {
    return sm ? sm->current_state : NULL;
}

const char* state_machine_to_string(const StateMachine* sm, char* buffer, size_t size) {
    if (!sm || !buffer) return NULL;
    snprintf(buffer, size, "name: %s , states: %zu", sm->name, sm->num_states);
    return buffer;
}

const char* state_to_string(const State* state, char* buffer, size_t size) {
    if (!state || !buffer) return NULL;
    snprintf(buffer, size, "key: %s", state->key);
    return buffer;
}
