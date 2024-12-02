#include <stdio.h>
#include <string.h>
#include <assert.h>
#include "state_machine.h"

#define BUFFER_SIZE 256

static const char* s1 = "state1";
static const char* s2 = "state2";
static const char* s3 = "state3";
static const char* s4 = "state4";

static StateKey state1_action(const State* state, Model* model, const void* input) {
    model->value = *(const int*)input;
    return s2;
}

static StateKey state2_action(const State* state, Model* model, const void* input) {
    model->value = *(const int*)input;
    return s3;
}

static StateKey state3_action(const State* state, Model* model, const void* input) {
    model->value = *(const int*)input;
    return s4;
}

static StateKey state4_action(const State* state, Model* model, const void* input) {
    model->value = *(const int*)input;
    return s1;
}

static void test_empty_state_machine(void) {
    char buffer[BUFFER_SIZE];
    StateMachine* sm = new_state_machine("test");
    
    const char* str = state_machine_to_string(sm, buffer, BUFFER_SIZE);
    assert(strcmp(str, "name: test , states: 0") == 0);
    
    free_state_machine(sm);
    printf("test_empty_state_machine: PASS\n");
}

static void test_add_state(void) {
    char buffer[BUFFER_SIZE];
    StateMachine* sm = new_state_machine("test");
    
    State* state1 = new_state(s1, state1_action, NULL);
    assert(add_state(sm, state1) == 0);
    
    const char* str = state_machine_to_string(sm, buffer, BUFFER_SIZE);
    assert(strcmp(str, "name: test , states: 1") == 0);
    
    State* current = get_current_state(sm);
    assert(strcmp(current->key, s1) == 0);
    
    free_state_machine(sm);
    printf("test_add_state: PASS\n");
}

static void test_state_machine_execution(void) {
    StateMachine* sm = new_state_machine("test");
    
    State* states[] = {
        new_state(s1, state1_action, NULL),
        new_state(s2, state2_action, NULL),
        new_state(s3, state3_action, NULL),
        new_state(s4, state4_action, NULL)
    };
    
    for (int i = 0; i < 4; i++) {
        assert(add_state(sm, states[i]) == 0);
    }
    
    Model model = {0};
    int inputs[] = {10, 20, 30, 40};
    const char* expected_keys[] = {s2, s3, s4, s1};
    
    for (int i = 0; i < 4; i++) {
        StateKey key = execute(sm, &model, &inputs[i]);
        assert(strcmp(key, expected_keys[i]) == 0);
        assert(model.value == inputs[i]);
    }
    
    free_state_machine(sm);
    printf("test_state_machine_execution: PASS\n");
}

int main(void) {
    test_empty_state_machine();
    test_add_state();
    test_state_machine_execution();
    
    printf("All tests passed!\n");
    return 0;
}
