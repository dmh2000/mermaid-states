#pragma once

#include <string>
#include <memory>
#include <map>
#include <functional>
#include <stdexcept>

namespace sm {

using StateKey = std::string;

template<typename Model, typename Input>
class State;

template<typename Model, typename Input>
using ActionFunc = std::function<StateKey(const State<Model, Input>&, Model&, const Input&)>;

template<typename Model, typename Input>
class State {
public:
    State(const StateKey& key, ActionFunc<Model, Input> action)
        : key_(key), action_(action) {
        if (key.empty() || !action) {
            throw std::invalid_argument("Invalid state parameters");
        }
    }

    StateKey execute(Model& model, const Input& input) const {
        return action_(*this, model, input);
    }

    StateKey getKey() const { return key_; }
    std::string toString() const { return "key: " + key_; }

private:
    StateKey key_;
    ActionFunc<Model, Input> action_;
};

template<typename Model, typename Input>
class StateMachine {
public:
    StateMachine(const std::string& name)
        : name_(name) {}

    void addState(std::shared_ptr<State<Model, Input>> state) {
        if (states_.find(state->getKey()) != states_.end()) {
            throw std::runtime_error("State " + state->getKey() + " already exists");
        }
        
        states_[state->getKey()] = state;
        
        if (!currentState_) {
            setInitialState(state->getKey());
        }
    }

    void setInitialState(const StateKey& key) {
        auto it = states_.find(key);
        if (it == states_.end()) {
            throw std::runtime_error("State " + key + " does not exist");
        }
        currentState_ = it->second;
    }

    StateKey execute(Model& model, const Input& input) {
        if (!currentState_) {
            throw std::runtime_error("No current state set");
        }

        StateKey nextKey = currentState_->execute(model, input);
        
        if (nextKey != currentState_->getKey()) {
            auto it = states_.find(nextKey);
            if (it == states_.end()) {
                throw std::runtime_error("State " + nextKey + " does not exist");
            }
            currentState_ = it->second;
        }

        return currentState_->getKey();
    }

    std::shared_ptr<State<Model, Input>> getCurrentState() const {
        return currentState_;
    }

    std::map<StateKey, std::shared_ptr<State<Model, Input>>> getStates() const {
        return states_;
    }

    std::string toString() const {
        return "name: " + name_ + " , states: " + std::to_string(states_.size());
    }

private:
    std::string name_;
    std::shared_ptr<State<Model, Input>> currentState_;
    std::map<StateKey, std::shared_ptr<State<Model, Input>>> states_;
};

} // namespace sm
