#include <gtest/gtest.h>
#include "state_machine.hpp"

struct TestModel {
    int value;
    explicit TestModel(int v = 0) : value(v) {}
};

class StateMachineTest : public ::testing::Test {
protected:
    const sm::StateKey s1 = "state1";
    const sm::StateKey s2 = "state2";
    const sm::StateKey s3 = "state3";
    const sm::StateKey s4 = "state4";
};

TEST_F(StateMachineTest, EmptyStateMachine) {
    sm::StateMachine<TestModel, int> machine("test");
    EXPECT_EQ(machine.toString(), "name: test , states: 0");
}

TEST_F(StateMachineTest, AddState) {
    sm::StateMachine<TestModel, int> machine("test");
    
    auto state1 = std::make_shared<sm::State<TestModel, int>>(
        s1,
        [](const sm::State<TestModel, int>& state, TestModel& model, const int& input) {
            model.value = input;
            return state.getKey();
        }
    );

    machine.addState(state1);
    EXPECT_EQ(machine.toString(), "name: test , states: 1");
    EXPECT_EQ(machine.getCurrentState()->getKey(), s1);
}

TEST_F(StateMachineTest, StateMachineExecution) {
    sm::StateMachine<TestModel, int> machine("test");
    
    auto state1 = std::make_shared<sm::State<TestModel, int>>(
        s1,
        [this](const sm::State<TestModel, int>&, TestModel& model, const int& input) {
            model.value = input;
            return s2;
        }
    );

    auto state2 = std::make_shared<sm::State<TestModel, int>>(
        s2,
        [this](const sm::State<TestModel, int>&, TestModel& model, const int& input) {
            model.value = input;
            return s3;
        }
    );

    auto state3 = std::make_shared<sm::State<TestModel, int>>(
        s3,
        [this](const sm::State<TestModel, int>&, TestModel& model, const int& input) {
            model.value = input;
            return s4;
        }
    );

    auto state4 = std::make_shared<sm::State<TestModel, int>>(
        s4,
        [this](const sm::State<TestModel, int>&, TestModel& model, const int& input) {
            model.value = input;
            return s1;
        }
    );

    machine.addState(state1);
    machine.addState(state2);
    machine.addState(state3);
    machine.addState(state4);

    TestModel model;
    std::vector<int> inputs = {10, 20, 30, 40};
    std::vector<sm::StateKey> expectedKeys = {s2, s3, s4, s1};

    for (size_t i = 0; i < inputs.size(); ++i) {
        sm::StateKey key = machine.execute(model, inputs[i]);
        EXPECT_EQ(key, expectedKeys[i]);
        EXPECT_EQ(model.value, inputs[i]);
    }
}
