package main

import (
	"fmt"
	"os"

	"github.com/NimbleMarkets/ntcharts/canvas"
	"github.com/NimbleMarkets/ntcharts/canvas/graph"
	"github.com/NimbleMarkets/ntcharts/canvas/runes"
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

	leftCircle := graph.GetFullCirclePoints(m.cursor.Add(canvas.Point{X: -1, Y: -1}), 3)
	rightCircle := graph.GetFullCirclePoints(m.cursor.Add(canvas.Point{X: 1, Y: 1}), 3)
	leftSet := map[canvas.Point]struct{}{}
	rightSet := map[canvas.Point]struct{}{}

	for _, p := range leftCircle {
		leftSet[p] = struct{}{}
	}
	for _, p := range rightCircle {
		rightSet[p] = struct{}{}
	}

	m.c.Clear()
	for _, p := range leftCircle {
		if _, ok := rightSet[p]; !ok {
			m.c.SetCell(p, canvas.NewCellWithStyle(runes.FullBlock, blueStyle))
		}
	}
	for _, p := range rightCircle {
		if _, ok := leftSet[p]; !ok {
			m.c.SetCell(p, canvas.NewCellWithStyle(runes.FullBlock, blueStyle))
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
	cursor := canvas.Point{X: w / 2, Y: h / 2}
	m := model{c, cursor}

	p := tea.NewProgram(m)
	if _, err := p.Run(); err != nil {
		fmt.Println("Error running program:", err)
		os.Exit(1)
	}
}
