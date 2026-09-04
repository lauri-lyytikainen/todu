package main

import (
	tea "github.com/charmbracelet/bubbletea"
	"log"
)

func main() {
	store := new(Store)
	if err := store.Init(); err != nil {
		log.Fatalf("unable to init store: %v", err)
	}
	m := NewModel(store)

	p := tea.NewProgram(m)

	_, err := p.Run()

	if err != nil {
		log.Fatalf("unable to run tui: %v", err)
	}
}
