package tui

import (
	"fmt"
	"log/slog"
	"strings"

	tea "charm.land/bubbletea/v2"
	"charm.land/lipgloss/v2"
)

type Page struct {
	id    uint
	title string
}

type RootWindow struct {
	activePage uint
	pages      []Page
	height     int
	width      int
}

func NewRootWindow() RootWindow {
	return RootWindow{activePage: 0, pages: []Page{
		{id: 1, title: "Тесты элементов"},
		{id: 2, title: "Тесты команд"},
		{id: 3, title: "Много активных элементов"},
	}}
}

func (m RootWindow) Init() tea.Cmd {
	return nil
}

func (m RootWindow) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width
		return m, nil
	case tea.KeyMsg:
		if msg.String() == "q" {
			if m.activePage != RootPage {
				m.activePage = RootPage
				return m, nil
			}
			return m, tea.Quit
		}
	}
	return m, nil
}

func (m RootWindow) View() tea.View {
	if m.activePage == RootPage {
		headerStyle := lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("#FAFAFA")).
			Background(lipgloss.Color("#7D56F4")).
			Padding(0, 1).
			MarginBottom(1)

		numStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#7D56F4")).
			Bold(true)

		titleStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#C4C4C4"))

		var listBuilder strings.Builder
		for _, p := range m.pages {
			item := fmt.Sprintf(
				"%s %s\n",
				numStyle.Render(fmt.Sprintf("%d:", p.id)),
				titleStyle.Render(p.title),
			)
			_, err := listBuilder.WriteString(item)
			if err != nil {
				slog.Error("ROOT WINDOW. PAGES BUILDER. ERROR", slog.String("error", err.Error()))
				return tea.NewView(lipgloss.NewStyle().
					Foreground(lipgloss.Color("9")).
					Render("Произошла ошибка во время загрузки страницы :("))
			}
		}

		header := headerStyle.Render("НАВИГАЦИЯ")
		content := lipgloss.JoinVertical(lipgloss.Center, header, listBuilder.String())

		boxStyle := lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).            // Закруглённая граница
			BorderForeground(lipgloss.Color("#7D56F4")). // Цвет границы
			Padding(1, 3).                               // Внутренние отступы
			Width(40).                                   // Ширина блока
			Align(lipgloss.Left)                         // Выравнивание контента внутри блока

		innerBox := boxStyle.Render(content)

		centeredUI := lipgloss.Place(
			m.width,         // Полная ширина терминала
			m.height,        // Полная высота терминала
			lipgloss.Center, // Горизонтальное выравнивание (Left, Center, Right)
			lipgloss.Center, // Вертикальное выравнивание (Top, Center, Bottom)
			innerBox,
		)
		v := tea.NewView(centeredUI)
		v.AltScreen = true
		return v
	}
	content := fmt.Sprintf("ТЕБЯ ТУТ БЫТЬ НЕ ДОЛЖНО")
	view := tea.NewView(content)
	view.AltScreen = true
	return view
}
