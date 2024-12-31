package main

// Inspired by https://github.com/NimbleMarkets/ntcharts/blob/main/examples/linechart/scatter/main.go

import (
	"fmt"
	"os"

	"github.com/NimbleMarkets/ntcharts/canvas"
	"github.com/NimbleMarkets/ntcharts/linechart"
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
	lc     linechart.Model
	cursor canvas.Point
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit
		}
	}
	m.lc.Clear()
	point := canvas.Float64Point{X: float64(50), Y: float64(50)}
	m.lc.DrawBrailleCircle(point, 10)
	return m, nil
}

func (m model) View() string {
	s := "Use the arrow keys to move the cursor, `q/ctrl+c` to quit\n"
	s += fmt.Sprintf("Cursor position: (%d, %d)\n", m.cursor.X, m.cursor.Y)
	s += whiteBoxStyle.Render(m.lc.View())
	s += "\n"
	return s
}

func main() {
	w := 50
	h := 20
	minX := 0.0
	maxX := 100.0
	minY := 0.0
	maxY := 100.0
	lc := linechart.New(w, h, minX, maxX, minY, maxY)

	m := model{lc, canvas.Point{X: 0, Y: 0}}

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
