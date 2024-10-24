package help

import (
	"fmt"
	"testing"

	"github.com/charmbracelet/x/ansi"
	"github.com/charmbracelet/x/exp/golden"

	"github.com/charmbracelet/bubbles/v2/key"
)

func TestFullHelp(t *testing.T) {
	m := New()
	m.Styles.FullSeparator = m.Styles.FullSeparator.SetString(" | ")
	k := key.WithKeys("x")
	kb := [][]key.Binding{
		{
			key.NewBinding(k, key.WithHelp("enter", "continue")),
		},
		{
			key.NewBinding(k, key.WithHelp("esc", "back")),
			key.NewBinding(k, key.WithHelp("?", "help")),
		},
		{
			key.NewBinding(k, key.WithHelp("H", "home")),
			key.NewBinding(k, key.WithHelp("ctrl+c", "quit")),
			key.NewBinding(k, key.WithHelp("ctrl+l", "log")),
		},
	}

	for _, w := range []int{20, 30, 40} {
		t.Run(fmt.Sprintf("full help %d width", w), func(t *testing.T) {
			m.Width = w
			s := m.FullHelpView(kb)
			s = ansi.Strip(s)
			golden.RequireEqual(t, []byte(s))
		})
	}
}
