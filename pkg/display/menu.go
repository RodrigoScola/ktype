package display

type Displayable interface {
    Display() string
}

type Menu struct {
    displayable Displayable
}
func NewMenu() Menu{
    return Menu{}
}

func (m *Menu) Display() string {
    if m.displayable != nil {
        return m.Display()
    }
    return ""
}

func (m *Menu) ToDisplay(displayable Displayable) {
    m.displayable = displayable


}
