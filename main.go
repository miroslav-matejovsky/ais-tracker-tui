package main

import (
	"fmt"
	"os"

	"github.com/NimbleMarkets/ntcharts/canvas"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

const (
	Purple = "63"
	Blue   = "4"
	White  = "15"
)

var whiteBoxStyle = lipgloss.NewStyle().
	BorderStyle(lipgloss.NormalBorder()).
	BorderForeground(lipgloss.Color(White))

var whiteStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color(White))

var blueStyle = lipgloss.NewStyle().
	Foreground(lipgloss.Color(Blue))

type model struct {
	c      canvas.Model
	cursor canvas.Point
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "up":
			m.cursor.Y--
			if m.cursor.Y < 0 {
				m.cursor.Y = 0
			}
		case "down":
			m.cursor.Y++
			if m.cursor.Y > m.c.Height()-1 {
				m.cursor.Y = m.c.Height() - 1
			}
		case "right":
			m.cursor.X++
			if m.cursor.X > m.c.Width()-1 {
				m.cursor.X = m.c.Width() - 1
			}
		case "left":
			m.cursor.X--
			if m.cursor.X < 0 {
				m.cursor.X = 0
			}
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}

	return m, nil
}

func (m model) View() string {
	s := "Use the arrow keys to move the cursor, `q/ctrl+c` to quit\n"
	s += fmt.Sprintf("Cursor position: (%d, %d)\n", m.cursor.X, m.cursor.Y)
	s += whiteBoxStyle.Render(m.c.View())
	s += "\n"
	return s
}

func main() {
	w := 50
	h := 20
	c := canvas.New(w, h)
	cursor := canvas.Point{X: 0, Y: 0}
	m := model{c, cursor}

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
