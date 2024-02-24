package main

import (
	"fmt"
	"os"

	"github.com/stevenwilkin/deribit-funding/deribit"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	d      = &deribit.Deribit{}
	margin = lipgloss.NewStyle().Margin(1, 2, 0, 2)
)

type model struct {
	funding float64
}

type fundingMsg float64

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyCtrlC:
			return m, tea.Quit
		}
	case fundingMsg:
		m.funding = float64(msg)
	}

	return m, nil
}

func (m model) View() string {
	if m.funding == 0 {
		return ""
	}

	return margin.Render(fmt.Sprintf("%f%%", m.funding*100))
}

func main() {
	m := model{}
	p := tea.NewProgram(m, tea.WithAltScreen(), tea.WithFPS(10))

	go func() {
		for funding := range d.Funding() {
			p.Send(fundingMsg(funding))
		}
		os.Exit(1)
	}()

	if err := p.Start(); err != nil {
		fmt.Printf("Error: %v", err)
		os.Exit(1)
	}
}
