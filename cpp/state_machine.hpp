#ifndef STATE_MACHINE_HPP
#define STATE_MACHINE_HPP

#include <string>
#include <map>
#include <memory>
#include <functional>
#include <stdexcept>

namespace sm {

using StateKey = int32_t;

template<typename Model, typename Input>
class State {
public:
    using ActionFunc = std::function<StateKey(Model&, const Input&)>;

    State(StateKey key, const std::string& name, ActionFunc action)
        : key_(key), name_(name), action_(action) {
        if (!action_) {
            throw std::invalid_argument("action function cannot be null");
        }
    }

    StateKey getKey() const { return key_; }
    std::string getName() const { return name_; }
    std::string toString() const;
    
    StateKey execute(Model& model, const Input& input) {
        return action_(model, input);
    }

private:
    StateKey key_;
    std::string name_;
    ActionFunc action_;
};

template<typename Model, typename Input>
class StateMachine {
public:
    StateMachine(Model& model, const std::string& name)
        : model_(model), name_(name), currentState_(nullptr) {}

    std::string toString() const;
    
    State<Model, Input>* getCurrentState() { return currentState_; }
    
    std::map<StateKey, std::unique_ptr<State<Model, Input>>> getStates() const;
    
    void addState(std::unique_ptr<State<Model, Input>> state);
    
    void setInitialState(StateKey key);
    
    StateKey execute(const Input& input);

private:
    Model& model_;
    std::string name_;
    State<Model, Input>* currentState_;
    std::map<StateKey, std::unique_ptr<State<Model, Input>>> states_;
};

} // namespace sm

#include "state_machine.inl"
#endif // STATE_MACHINE_HPP
