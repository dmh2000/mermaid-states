#include "state_machine.h"
#include <stdlib.h>
#include <string.h>

#define MAX_STATES 256
#define MAX_NAME_LEN 64

struct state_t {
    state_key_t key;
    char name[MAX_NAME_LEN];
    action_func_t action;
};

struct state_machine_t {
    void* model;
    char name[MAX_NAME_LEN];
    state_t* current_state;
    state_t* states[MAX_STATES];
    size_t num_states;
};

state_t* state_new(state_key_t key, const char* name, action_func_t action) {
    if (!name || !action) {
        return NULL;
    }

    state_t* state = malloc(sizeof(state_t));
    if (!state) {
        return NULL;
    }

    state->key = key;
    strncpy(state->name, name, MAX_NAME_LEN - 1);
    state->name[MAX_NAME_LEN - 1] = '\0';
    state->action = action;

    return state;
}

void state_free(state_t* state) {
    free(state);
}

state_key_t state_get_key(const state_t* state) {
    return state ? state->key : 0;
}

const char* state_get_name(const state_t* state) {
    return state ? state->name : NULL;
}

state_key_t state_execute(state_t* state, void* model, void* input) {
    return state && state->action ? state->action(model, input) : 0;
}

state_machine_t* state_machine_new(void* model, const char* name) {
    if (!name) {
        return NULL;
    }

    state_machine_t* sm = malloc(sizeof(state_machine_t));
    if (!sm) {
        return NULL;
    }

    sm->model = model;
    strncpy(sm->name, name, MAX_NAME_LEN - 1);
    sm->name[MAX_NAME_LEN - 1] = '\0';
    sm->current_state = NULL;
    sm->num_states = 0;
    memset(sm->states, 0, sizeof(sm->states));

    return sm;
}

void state_machine_free(state_machine_t* sm) {
    if (sm) {
        for (size_t i = 0; i < sm->num_states; i++) {
            state_free(sm->states[i]);
        }
        free(sm);
    }
}

bool state_machine_add_state(state_machine_t* sm, state_t* state) {
    if (!sm || !state || sm->num_states >= MAX_STATES) {
        return false;
    }

    // Check for duplicate key
    for (size_t i = 0; i < sm->num_states; i++) {
        if (sm->states[i]->key == state->key) {
            return false;
        }
    }

    sm->states[sm->num_states++] = state;
    
    // Set as initial state if first state added
    if (sm->num_states == 1) {
        sm->current_state = state;
    }

    return true;
}

bool state_machine_set_initial_state(state_machine_t* sm, state_key_t key) {
    if (!sm) {
        return false;
    }

    for (size_t i = 0; i < sm->num_states; i++) {
        if (sm->states[i]->key == key) {
            sm->current_state = sm->states[i];
            return true;
        }
    }

    return false;
}

state_key_t state_machine_execute(state_machine_t* sm, void* input) {
    if (!sm || !sm->current_state) {
        return 0;
    }

    state_key_t next_key = state_execute(sm->current_state, sm->model, input);
    
    // Stay in same state if key matches current
    if (next_key == sm->current_state->key) {
        return next_key;
    }

    // Find and transition to next state
    for (size_t i = 0; i < sm->num_states; i++) {
        if (sm->states[i]->key == next_key) {
            sm->current_state = sm->states[i];
            return next_key;
        }
    }

    return 0; // Error - invalid next state
}

state_t* state_machine_get_current_state(const state_machine_t* sm) {
    return sm ? sm->current_state : NULL;
}
