#include "state_machine.hpp"
#include <cassert>
#include <iostream>

enum States {
    INVALID = 0,
    STATE1,
    STATE2,
    STATE3,
    STATE4
};

struct TestModel {
    int value;
};

void test_empty() {
    TestModel model{0};
    sm::StateMachine<TestModel, int> machine(model, "test");
    assert(machine.toString() == "name: test , states: 0");
    std::cout << "Empty test passed\n";
}

void test_add_state() {
    TestModel model{0};
    sm::StateMachine<TestModel, int> machine(model, "test");
    
    auto state = std::make_unique<sm::State<TestModel, int>>(
        STATE1,
        "state1",
        [](TestModel& m, const int& input) {
            m.value = input;
            std::cout << "State 1 (key=" << STATE1 << ", name=state1): model.value = " << m.value << "\n";
            return STATE2;
        }
    );
    
    machine.addState(std::move(state));
    assert(machine.toString() == "name: test , states: 1");
    assert(machine.getCurrentState()->toString() == "key: 1 , name: state1");
    std::cout << "Add state test passed\n";
}

void test_state_machine_execution() {
    TestModel model{0};
    sm::StateMachine<TestModel, int> machine(model, "test");

    // Add states
    machine.addState(std::make_unique<sm::State<TestModel, int>>(
        STATE1,
        "state1",
        [](TestModel& m, const int& input) {
            m.value = input;
            return STATE2;
        }
    ));

    machine.addState(std::make_unique<sm::State<TestModel, int>>(
        STATE2,
        "state2",
        [](TestModel& m, const int& input) {
            m.value = input;
            std::cout << "State 2 (key=" << STATE2 << ", name=state2): model.value = " << m.value << "\n";
            return STATE3;
        }
    ));

    machine.addState(std::make_unique<sm::State<TestModel, int>>(
        STATE3,
        "state3",
        [](TestModel& m, const int& input) {
            m.value = input;
            std::cout << "State 3 (key=" << STATE3 << ", name=state3): model.value = " << m.value << "\n";
            return STATE4;
        }
    ));

    machine.addState(std::make_unique<sm::State<TestModel, int>>(
        STATE4,
        "state4",
        [](TestModel& m, const int& input) {
            m.value = input;
            std::cout << "State 4 (key=" << STATE4 << ", name=state4): model.value = " << m.value << "\n";
            return STATE4;
        }
    ));

    // Execute state machine
    int inputs[] = {10, 20, 30, 40};
    States expected[] = {STATE2, STATE3, STATE4, STATE4};

    for (size_t i = 0; i < 4; ++i) {
        auto key = machine.execute(inputs[i]);
        assert(key == expected[i]);
        assert(model.value == inputs[i]);
    }
    
    std::cout << "Execution test passed\n";
}

int main() {
    try {
        test_empty();
        test_add_state();
        test_state_machine_execution();
        std::cout << "All tests passed!\n";
        return 0;
    } catch (const std::exception& e) {
        std::cerr << "Test failed: " << e.what() << std::endl;
        return 1;
    }
}
