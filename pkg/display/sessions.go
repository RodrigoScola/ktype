package display

import (

	huh "github.com/charmbracelet/huh"
)

type sessionModel struct {
	names []string
}

func NewSessionMenu(names []string, sessionname *string) *huh.Form {
	opts := make([]huh.Option[string], len(names))
	for i := range names {
		opts[i] = huh.NewOption(names[i], names[i])
	}

	return huh.NewForm(
		huh.NewGroup(
			huh.NewSelect[string]().
				Title("Choose your session").
				Options(opts...).Value(sessionname),
		),
	)
}



func NewCreateSessionMenu(sessionfilepath *string) *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Title("what is the file path?").
                Suggestions([]string{"./Another-Kingdom-by-Andrew-Klavan.txt"}).
				Value(sessionfilepath),
		),
	)
}
