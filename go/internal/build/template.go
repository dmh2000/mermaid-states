package build

const stateTemplate = `
NewState(
    "{{STATEKEY}}",
	func (current *State[{{MODEL}}, {{INPUT}}], model *{{MODEL}}, input {{INPUT}}) (key StateKey, err error) {
		// COMMENT
		return {{NEWSTATE}},nil
	},
    nil,
)
`
