#include "state_machine.h"
#include <stdio.h>
#include <assert.h>

enum {
    INVALID = 0,
    STATE1,
    STATE2,
    STATE3,
    STATE4
};

typedef struct {
    int value;
} test_model_t;

static state_key_t state1_action(void* model, void* input) {
    test_model_t* m = (test_model_t*)model;
    m->value = *(int*)input;
    return STATE2;
}

static state_key_t state2_action(void* model, void* input) {
    test_model_t* m = (test_model_t*)model;
    m->value = *(int*)input;
    return STATE3;
}

static state_key_t state3_action(void* model, void* input) {
    test_model_t* m = (test_model_t*)model;
    m->value = *(int*)input;
    return STATE4;
}

static state_key_t state4_action(void* model, void* input) {
    test_model_t* m = (test_model_t*)model;
    m->value = *(int*)input;
    return STATE4;
}

void test_empty(void) {
    test_model_t model = {0};
    state_machine_t* machine = state_machine_new(&model, "test");
    assert(machine != NULL);
    assert(state_machine_get_current_state(machine) == NULL);
    state_machine_free(machine);
    printf("Empty test passed\n");
}

void test_add_state(void) {
    test_model_t model = {0};
    state_machine_t* machine = state_machine_new(&model, "test");
    
    state_t* state = state_new(STATE1, "state1", state1_action);
    assert(state != NULL);
    
    assert(state_machine_add_state(machine, state));
    assert(state_machine_get_current_state(machine) == state);
    assert(state_get_key(state) == STATE1);
    
    state_machine_free(machine);
    printf("Add state test passed\n");
}

void test_state_machine_execution(void) {
    test_model_t model = {0};
    state_machine_t* machine = state_machine_new(&model, "test");

    // Create and add states
    state_t* states[] = {
        state_new(STATE1, "state1", state1_action),
        state_new(STATE2, "state2", state2_action),
        state_new(STATE3, "state3", state3_action),
        state_new(STATE4, "state4", state4_action)
    };

    for (int i = 0; i < 4; i++) {
        assert(state_machine_add_state(machine, states[i]));
    }

    // Test execution
    int inputs[] = {10, 20, 30, 40};
    state_key_t expected[] = {STATE2, STATE3, STATE4, STATE4};

    for (int i = 0; i < 4; i++) {
        state_key_t key = state_machine_execute(machine, &inputs[i]);
        assert(key == expected[i]);
        assert(model.value == inputs[i]);
    }

    state_machine_free(machine);
    printf("Execution test passed\n");
}

int main(void) {
    test_empty();
    test_add_state();
    test_state_machine_execution();
    printf("All tests passed!\n");
    return 0;
}
