package main

import (
	"fmt"
	"os"

	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
	"github.com/charmbracelet/lipgloss"
)

type model struct {
	viewport     viewport.Model
	ready        bool
	rawContent   string
	errorMessage string
}

func initialModel(filePath string) model {
	fileBytes, err := os.ReadFile(filePath)
	if err != nil {
		return model{errorMessage: fmt.Sprintf("❌ File read error: %v", err)}
	}

	// Render Markdown with Glamour dark style
	rendered, err := glamour.Render(string(fileBytes), "dark")
	if err != nil {
		return model{errorMessage: fmt.Sprintf("❌ Markdown parsing error: %v", err)}
	}

	return model{
		rawContent:   rendered,
		ready:        false,
		errorMessage: "",
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmds []tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c":
			return m, tea.Quit

		case "k", "up":
			m.viewport.LineUp(1)
		case "j", "down":
			m.viewport.LineDown(1)
		case "u", "pgup":
			m.viewport.HalfPageUp()
		case "d", "pgdown":
			m.viewport.HalfPageDown()
		}

	case tea.WindowSizeMsg:
		headerHeight := 3
		footerHeight := 2
		verticalMarginHeight := headerHeight + footerHeight

		if !m.ready {
			m.viewport = viewport.New(msg.Width, msg.Height-verticalMarginHeight)
			m.viewport.YPosition = headerHeight

			if m.errorMessage != "" {
				m.viewport.SetContent(m.errorMessage)
			} else {
				m.viewport.SetContent(m.rawContent)
			}
			m.ready = true
		} else {
			m.viewport.Width = msg.Width
			m.viewport.Height = msg.Height - verticalMarginHeight
		}
	}

	var cmd tea.Cmd
	m.viewport, cmd = m.viewport.Update(msg)
	cmds = append(cmds, cmd)

	return m, tea.Batch(cmds...)
}

func (m model) View() string {
	if m.errorMessage != "" && !m.ready {
		return m.errorMessage
	}
	if !m.ready {
		return "🔄 Loading and calculating terminal dimensions..."
	}

	header := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#5A5CD6")).
		Padding(0, 1).
		Render("🎨 COLORMARK TUI")

	footer := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#626262")).
		Render(fmt.Sprintf("\n[Arrow Keys / j / k to scroll] • [Press 'q' to quit]"))

	return fmt.Sprintf("%s\n\n%s\n%s", header, m.viewport.View(), footer)
}

func main() {
	if len(os.Args) < 2 {
		fmt.Println("👉 Usage: github.com/mehmetalidsy/colormark <markdown-file-path>")
		os.Exit(1)
	}

	p := tea.NewProgram(
		initialModel(os.Args[1]),
		tea.WithAltScreen(),
	)

	if _, err := p.Run(); err != nil {
		fmt.Printf("Oh no, an error occurred: %v\n", err)
		os.Exit(1)
	}
}