package main

import (
	"fmt"
	"os"

	"github.com/NimbleMarkets/ntcharts/canvas"
	tea "github.com/charmbracelet/bubbletea"
)

type model struct {
	c canvas.Model
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		if msg.String() == "q" {
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	return m.c.View()
}

func main() {
	w := 20
	h := 11
	c := canvas.New(w, h)
	m := model{c}

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
