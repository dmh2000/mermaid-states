#ifndef STATE_MACHINE_INL
#define STATE_MACHINE_INL

namespace sm {

template<typename Model, typename Input>
std::string State<Model, Input>::toString() const {
    return "key: " + std::to_string(key_) + " , name: " + name_;
}

template<typename Model, typename Input>
std::string StateMachine<Model, Input>::toString() const {
    return "name: " + name_ + " , states: " + std::to_string(states_.size());
}

template<typename Model, typename Input>
std::map<StateKey, std::unique_ptr<State<Model, Input>>> StateMachine<Model, Input>::getStates() const {
    std::map<StateKey, std::unique_ptr<State<Model, Input>>> states;
    for (const auto& pair : states_) {
        states[pair.first] = std::make_unique<State<Model, Input>>(*pair.second);
    }
    return states;
}

template<typename Model, typename Input>
void StateMachine<Model, Input>::addState(std::unique_ptr<State<Model, Input>> state) {
    if (states_.find(state->getKey()) != states_.end()) {
        throw std::runtime_error("state " + std::to_string(state->getKey()) + " already exists");
    }

    State<Model, Input>* statePtr = state.get();
    states_[state->getKey()] = std::move(state);

    if (!currentState_) {
        currentState_ = statePtr;
    }
}

template<typename Model, typename Input>
void StateMachine<Model, Input>::setInitialState(StateKey key) {
    auto it = states_.find(key);
    if (it == states_.end()) {
        throw std::runtime_error("state " + std::to_string(key) + " does not exist");
    }
    currentState_ = it->second.get();
}

template<typename Model, typename Input>
StateKey StateMachine<Model, Input>::execute(const Input& input) {
    if (!currentState_) {
        throw std::runtime_error("no current state set");
    }

    StateKey nextKey = currentState_->execute(model_, input);

    if (nextKey == currentState_->getKey()) {
        return nextKey;
    }

    auto it = states_.find(nextKey);
    if (it == states_.end()) {
        throw std::runtime_error("state " + std::to_string(nextKey) + " does not exist");
    }

    currentState_ = it->second.get();
    return nextKey;
}

} // namespace sm

#endif // STATE_MACHINE_INL
