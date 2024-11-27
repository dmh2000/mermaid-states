package build

import (
	"fmt"
	"strings"
	"testing"
)

type pair struct {
	key string
	val string
}

const model = "XModel"
const input = "XInput"

var states = [][]pair{
	{
		{"STATEKEY", "1"},
		{"MODEL", model},
		{"INPUT", input},
		{"NEWSTATE", "2"},
	},
	{
		{"STATEKEY", "2"},
		{"MODEL", model},
		{"INPUT", input},
		{"NEWSTATE", "3"},
	},
	{
		{"STATEKEY", "3"},
		{"MODEL", model},
		{"INPUT", input},
		{"NEWSTATE", "4"},
	},
	{
		{"STATEKEY", "4"},
		{"MODEL", model},
		{"INPUT", input},
		{"NEWSTATE", "5"},
	},
	{
		{"STATEKEY", "5"},
		{"MODEL", model},
		{"INPUT", input},
		{"NEWSTATE", "6"},
	},
}

func TestStateTemplate(t *testing.T) {

	const t1 = stateTemplate

	for _, state := range states {
		b := t1
		for _, p := range state {
			b = strings.ReplaceAll(b, "{{"+p.key+"}}", p.val)
		}

		fmt.Println(b, ",")
	}
}
