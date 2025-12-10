package main

import (
	"fmt"
	"slices"
	"strconv"
	"strings"

	"github.com/wimspaargaren/aoc"
)

func main() {
	input := aoc.MustReadFile("input.txt")
	machines := []Machine{}
	for _, line := range strings.Split(input, "\n") {
		machine := parseMachine(line)
		machines = append(machines, machine)
	}
	part1 := 0
	for i, machine := range machines {
		fmt.Println("Processing machine", i+1, "of", len(machines))
		part1 += process(machine)
	}
	fmt.Println("Part 1:", part1)
}

type Set struct {
	counter []int
	presses []int
}

func NewSet(config []string) *Set {
	return &Set{
		counter: make([]int, len(config)),
		presses: []int{},
	}
}

func (s *Set) Add(buttons [][]int, press int) {
	for _, index := range buttons[press] {
		s.counter[index]++
	}
	s.presses = append(s.presses, press)
	slices.Sort(s.presses)
}

func (s *Set) Clone() *Set {
	newCounter := make([]int, len(s.counter))
	copy(newCounter, s.counter)
	newPresses := make([]int, len(s.presses))
	copy(newPresses, s.presses)
	return &Set{
		counter: newCounter,
		presses: newPresses,
	}
}

func (s *Set) Check(config []string) bool {
	for i, conf := range config {
		if conf == "." {
			s.counter[i] %= 2
			if s.counter[i] != 0 {
				return false
			}
		}
		if conf == "#" {
			s.counter[i] %= 2
			if s.counter[i] != 1 {
				return false
			}
		}
	}
	return true
}

func process(machine Machine) int {
	trackMap := make(map[string]bool)
	initialSet := []*Set{}
	for i := range machine.Buttons {
		newSet := NewSet(machine.Config)
		newSet.Add(machine.Buttons, i)
		initialSet = append(initialSet, newSet)
	}
	for {
		for _, set := range initialSet {
			if set.Check(machine.Config) {
				return len(set.presses)
			}
		}
		initialSet = increaseSets(trackMap, initialSet, machine.Buttons)
		if len(initialSet) == 0 {
			panic("NOTHIGN FOUND!")
		}
	}
	return -1
}

// adds all combinations of button presses to the sets
func increaseSets(trackMap map[string]bool, sets []*Set, buttons [][]int) []*Set {
	newSets := []*Set{}
	for _, set := range sets {
		for i := range buttons {
			clonedSet := set.Clone()
			clonedSet.Add(buttons, i)
			_, ok := trackMap[fmtIntSlice(clonedSet.presses)]
			if ok {
				continue
			}
			trackMap[fmtIntSlice(clonedSet.presses)] = true
			newSets = append(newSets, clonedSet)
		}
	}
	return newSets
}

func fmtIntSlice(slice []int) string {
	strs := []string{}
	for _, num := range slice {
		strs = append(strs, strconv.Itoa(num))
	}
	return strings.Join(strs, ",")
}

func checkSet(config []string, set []int, buttons [][]int) bool {
	counter := make([]int, len(config))
	for _, pres := range set {
		for _, index := range buttons[pres] {
			counter[index]++
		}
	}
	for i, conf := range config {
		if conf == "." {
			counter[i] %= 2
			if counter[i] != 0 {
				return false
			}
		}
		if conf == "#" {
			counter[i] %= 2
			if counter[i] != 1 {
				return false
			}
		}
	}
	return true
}

func parseMachine(input string) Machine {
	parts := strings.Split(input, " ")
	config := parts[0]
	jotlageRequirement := parts[len(parts)-1]

	buttons := [][]int{}
	for i := 1; i < len(parts)-1; i++ {
		part := parts[i]
		temp := part[1 : len(part)-1]
		button := []int{}
		for _, s := range strings.Split(temp, ",") {
			index, err := strconv.Atoi(s)
			if err != nil {
				panic(err)
			}
			button = append(button, index)
		}
		buttons = append(buttons, button)
	}
	config = config[1 : len(config)-1]

	return Machine{
		Config:              strings.Split(config, ""),
		JoltageRequirements: jotlageRequirement,
		Buttons:             buttons,
	}
}

type Machine struct {
	Config              []string
	Buttons             [][]int
	JoltageRequirements string
}
