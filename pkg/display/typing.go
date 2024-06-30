package display

import (
	"fmt"
	"time"

	"github.com/RodrigoScola/ktype/pkg/book"
	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)


type model struct {
	width  int
	index  int
	styles *styles
	height int
	Input  textinput.Model
    Book *book.Book
}

func New(session *book.Book) *model {
	answerField := textinput.New()
	styles := DefaultStyles()
	return &model{
		Book: session, Input: answerField,
		styles: styles,
	}
}

type styles struct {
	Border     lipgloss.Color
	InputField lipgloss.Style
	Primary    lipgloss.Style
	Secondary  lipgloss.Style
}

func DefaultStyles() *styles {
	s := new(styles)

	s.Border = lipgloss.Color("36")
	s.Primary = lipgloss.NewStyle().Foreground(lipgloss.Color("81"))
	s.Secondary = lipgloss.NewStyle().Foreground(lipgloss.Color("84")).
		Bold(true).Italic(true)
	s.InputField = lipgloss.NewStyle().
		BorderForeground(s.Border).
		BorderStyle(lipgloss.NormalBorder()).
		Padding(1).Width(80)

	return s
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.height = msg.Height
		m.width = msg.Width

	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c":
			return m, tea.Quit
		case "backspace":
			m.Book.Current().Current.RemoveLast()
		case "enter":
			m.index++
			m.Input.SetValue("done!")
			return m, nil
		default:
			m.Book.Current().Current.Add(book.Letter{
				Char:      msg.String()[0],
				Ignore:    false,
				CreatedAt: time.Now(),
			})
		}
	}
    if m.Book.Current().Complete() {
        fmt.Println("saving")
        //m.Book.Session.Save()
        fmt.Println("saved")
        m.Book.Next()
    }
	m.Input, cmd = m.Input.Update(msg)

	return m, cmd
}

func (m model) View() string {
	if m.width == 0 {
		return "loading..."
	}
	view := m.Book.Current().Display(m.styles.Primary, m.styles.Secondary)

	return lipgloss.Place(
		m.width, m.height, lipgloss.Center, lipgloss.Center,
		lipgloss.JoinVertical(lipgloss.Center,
			m.styles.InputField.Render(view),
		))
}
