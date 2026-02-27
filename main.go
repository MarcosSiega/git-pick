package main

import (
	"fmt"
	"os"
	"os/exec"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

type item struct {
	title, desc string
}

func (i item) Title() string       { return i.title }
func (i item) Description() string { return i.desc }
func (i item) FilterValue() string { return i.title }

type model struct {
	list     list.Model
	selected map[string]bool
	quitting bool
}

var (
	delegate = list.NewDefaultDelegate()
	baseStyle = lipgloss.NewStyle().Padding(1)
	borderStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		Padding(1, 2).
		Margin(1).
		BorderForeground(lipgloss.Color("63"))
)

func getGitChanges() ([]list.Item, error) {
	cmd := exec.Command("git", "status", "--porcelain")
	output, err := cmd.Output()
	if err != nil {
		return nil, err
	}

	lines := strings.Split(string(output), "\n")
	var items []list.Item
	for _, line := range lines {
		if len(line) < 4 {
			continue
		}
		status := strings.TrimSpace(line[:2])
		file := strings.TrimSpace(line[3:])
		items = append(items, item{title: file, desc: fmt.Sprintf("Status: %s", status)})
	}
	return items, nil
}

func initialModel() model {
	items, err := getGitChanges()
	if err != nil {
		fmt.Println("Error getting git changes:", err)
		os.Exit(1)
	}

	l := list.New(items, delegate, 0, 0)
	l.Title = "gitpick"
	return model{list: l, selected: make(map[string]bool)}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			m.quitting = true
			return m, tea.Quit
		case " ":
			curItem := m.list.SelectedItem().(item)
			m.selected[curItem.title] = !m.selected[curItem.title]
		case "enter":
			for file, selected := range m.selected {
				if selected {
					exec.Command("git", "add", file).Run()
				}
			}
			m.quitting = true
			return m, tea.Quit
		}
	}
	var cmd tea.Cmd
	m.list, cmd = m.list.Update(msg)
	return m, cmd
}

func (m model) View() string {
	if m.quitting {
		return baseStyle.Render("\n✔ Done! Exiting...\n")
	}

	var b strings.Builder
	header := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("205")).
		Render("🌱 gitpick — select files with [space], confirm with [enter], or [q] to quit")

	b.WriteString(header)
	b.WriteString("\n\n")

	for i, listItem := range m.list.Items() {
		it := listItem.(item)
		prefix := "[ ]"
		if m.selected[it.title] {
			prefix = "[x]"
		}

		line := fmt.Sprintf("%s %s", prefix, it.title)
		if i == m.list.Index() {
			cursor := lipgloss.NewStyle().
				MarginLeft(1).
				Foreground(lipgloss.Color("57")).
				Render("➤")
			line = cursor + " " + line
		} else {
			line = "  " + line
		}
		b.WriteString(line + "\n")
	}

	footer := lipgloss.NewStyle().
		MarginTop(1).
		Faint(true).
		Render("\n↑↓ navigate · space select · enter add · q quit")

	return borderStyle.Render(b.String() + footer)
}

func main() {
	p := tea.NewProgram(initialModel())
	if err := p.Start(); err != nil {
		fmt.Printf("Error starting the program: %v\n", err)
		os.Exit(1)
	}
}