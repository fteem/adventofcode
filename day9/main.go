package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

type Machine struct {
	groups       int
	ignoreNext   bool
	score        int
	inGarbage    bool
	garbageCount int
}

func (m *Machine) info() {
	fmt.Printf("groups: %v, ignoreNext: %v, score: %v, inGarbage: %v, garbageCount: %v\n", m.groups, m.ignoreNext, m.score, m.inGarbage, m.garbageCount)
}

func (m *Machine) process(char string) {
	if !m.ignoreNext {
		if m.inGarbage {
			if !m.ignoreNext && char == ">" {
				m.inGarbage = false
			} else if char == "!" {
				m.ignoreNext = true
			} else {
				m.garbageCount++
			}
		} else {
			if char == "{" {
				m.groups++
			} else if char == "}" {
				m.score += m.groups
				m.groups--
			} else if char == "<" {
				m.inGarbage = true
			} else if m.inGarbage && char == "!" {
				m.ignoreNext = true
			}
		}
	} else {
		m.ignoreNext = false
	}
}

func main() {
	contents, err := ioutil.ReadFile(os.Args[1])
	if err != nil {
		fmt.Println("Error opening file:", err)
		os.Exit(1)
	}

	data := strings.Trim(string(contents), "\n")
	chars := strings.Split(data, "")

	machine := Machine{}

	for _, char := range chars {
		machine.process(char)
	}

	machine.info()
}
