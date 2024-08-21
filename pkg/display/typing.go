package display

import (
	"time"

	"github.com/RodrigoScola/ktype/pkg/book"
	filesessions "github.com/RodrigoScola/ktype/pkg/file_sessions"
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
	Book   *book.Book
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
    Cursor lipgloss.Style
	Border     lipgloss.Color
	InputField lipgloss.Style
	Primary    lipgloss.Style
	Secondary  lipgloss.Style
	Text       lipgloss.Style
}

func DefaultStyles() *styles {
	s := new(styles)
	s.Border = lipgloss.Color("37")
    s.Cursor = lipgloss.NewStyle().Background(lipgloss.Color("#f2f2f2")).Foreground(lipgloss.Color("#000000"))
	s.Primary = lipgloss.NewStyle().Foreground(lipgloss.Color("#ffe300"))
	s.Secondary = lipgloss.NewStyle().Foreground(lipgloss.Color("#d44729")).
		Bold(true).Italic(true)
	s.Text = lipgloss.NewStyle()
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
		default:
			m.Book.Current().Current.Add(book.Letter{
				Char:      msg.String()[0],
				Ignore:    false,
				CreatedAt: time.Now(),
			})
			if m.Book.Current().Complete() {
				filesessions.Save(m.Book)
				m.Book.Next()
			}
		}
	}
	m.Input, cmd = m.Input.Update(msg)

	return m, cmd
}

func (m model) View() string {
	if m.width == 0 {
		return "loading..."
	}
	view := m.Book.Current().Display(m.styles.Primary, m.styles.Secondary, m.styles.Text, m.styles.Cursor)
	timestamps := []time.Time{}

	for i := range m.Book.Current().Current.Letters {
		item := m.Book.Current().Current.Letters[i]
		timestamps = append(timestamps, item.CreatedAt)
	}

   // wc := statistics.GetWords(m.Book.Current().Current.String())

	//var wpm string = ""
	// if len(timestamps) > 2 {
	// 	wpm = fmt.Sprintf("%.3f", statistics.CalculateWPM(timestamps[0], timestamps[len(timestamps)-1],wc))
	// }

	return lipgloss.Place(
		m.width, m.height, lipgloss.Center, lipgloss.Center,
		lipgloss.JoinVertical(lipgloss.Center,
			m.styles.InputField.Render(view),
			//wpm,
		))
}
