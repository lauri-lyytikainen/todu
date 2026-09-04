package main

import (
	"log"

	tea "github.com/charmbracelet/bubbletea"
)

const (
	listView uint = iota
	taskView
)

type model struct {
	state     uint
	store     *Store
	tasks     []Task
	listIndex int
}

func NewModel(store *Store) model {
	tasks, err := store.GetTasks()
	if err != nil {
		log.Fatalf("unable to get notes: %v", err)
	}
	return model{
		state: listView,
		store: store,
		tasks: tasks,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		key := msg.String()
		switch m.state {
		case listView:
			switch key {
			case "q":
				return m, tea.Quit
			case "up", "k":
				if m.listIndex > 0 {
					m.listIndex--
				}
			case "down", "j":
				if m.listIndex > len(m.tasks) {
					m.listIndex++
				}
			}
		}
	}
	return m, nil
}
