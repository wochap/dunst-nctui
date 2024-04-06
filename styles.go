package main

import (
	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
	"github.com/spf13/viper"
)

const (
	bullet   = "•"
	ellipsis = "…"
)

// DefaultListStyles returns a set of default style definitions for this list
// component.
func DefaultListStyles() (s list.Styles) {
	verySubdued := viper.GetString("theme.verySubdued")
	subdued := viper.GetString("theme.subdued")
	titleFg := viper.GetString("theme.titleFg")
	spinnerFg := viper.GetString("theme.spinnerFg")
	filterPromptFg := viper.GetString("theme.filterPromptFg")
	filterCursorFg := viper.GetString("theme.filterCursorFg")
	statusBarFg := viper.GetString("theme.statusBarFg")
	statusBarActiveFilterFg := viper.GetString("theme.statusBarActiveFilterFg")
	noItemsFg := viper.GetString("theme.noItemsFg")
	verySubduedColor := lipgloss.Color(verySubdued)
	subduedColor := lipgloss.Color(subdued)

	s.TitleBar = lipgloss.NewStyle().Padding(0, 0, 1, 2)

	s.Title = lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(titleFg)).
		MarginTop(1)

	s.Spinner = lipgloss.NewStyle().
		Foreground(lipgloss.Color(spinnerFg))

	s.FilterPrompt = lipgloss.NewStyle().
		Foreground(lipgloss.Color(filterPromptFg))

	s.FilterCursor = lipgloss.NewStyle().
		Foreground(lipgloss.Color(filterCursorFg))

	s.DefaultFilterCharacterMatch = lipgloss.NewStyle().Underline(true)

	s.StatusBar = lipgloss.NewStyle().
		Foreground(lipgloss.Color(statusBarFg)).
		Padding(0, 0, 1, 2)

	s.StatusEmpty = lipgloss.NewStyle().Foreground(subduedColor)

	s.StatusBarActiveFilter = lipgloss.NewStyle().
		Foreground(lipgloss.Color(statusBarActiveFilterFg))

	s.StatusBarFilterCount = lipgloss.NewStyle().Foreground(verySubduedColor)

	s.NoItems = lipgloss.NewStyle().
		Foreground(lipgloss.Color(noItemsFg))

	s.ArabicPagination = lipgloss.NewStyle().Foreground(subduedColor)

	s.PaginationStyle = lipgloss.NewStyle().PaddingLeft(2) //nolint:gomnd

	s.HelpStyle = lipgloss.NewStyle().Padding(1, 0, 0, 2)

	s.ActivePaginationDot = lipgloss.NewStyle().
		Foreground(subduedColor).
		SetString(bullet)

	s.InactivePaginationDot = lipgloss.NewStyle().
		Foreground(verySubduedColor).
		SetString(bullet)

	s.DividerDot = lipgloss.NewStyle().
		Foreground(verySubduedColor).
		SetString(" " + bullet + " ")

	return s
}

func DefaultItemStyles() (s list.DefaultItemStyles) {
	normalTitleFg := viper.GetString("theme.normalTitleFg")
	normalDescFg := viper.GetString("theme.normalDescFg")
	selectedTitleFg := viper.GetString("theme.selectedTitleFg")
	selectedDescFg := viper.GetString("theme.selectedDescFg")
	dimmedTitleFg := viper.GetString("theme.dimmedTitleFg")
	dimmedDescFg := viper.GetString("theme.dimmedDescFg")

	s.NormalTitle = lipgloss.NewStyle().
		Foreground(lipgloss.Color(normalTitleFg)).
		Padding(0, 2, 0, 2)

	s.NormalDesc = s.NormalTitle.Copy().
		Foreground(lipgloss.Color(normalDescFg))

	s.SelectedTitle = lipgloss.NewStyle().
		Foreground(lipgloss.Color(selectedTitleFg)).
		Padding(0, 2, 0, 2)

	s.SelectedDesc = s.SelectedTitle.Copy().
		Foreground(lipgloss.Color(selectedDescFg))

	s.DimmedTitle = lipgloss.NewStyle().
		Foreground(lipgloss.Color(dimmedTitleFg)).
		Padding(0, 2, 0, 2)

	s.DimmedDesc = s.DimmedTitle.Copy().
		Foreground(lipgloss.Color(dimmedDescFg))

	s.FilterMatch = lipgloss.NewStyle().Underline(true)

	return s
}

func DefaultHelpStyles() (s help.Styles) {
	keyStyleFg := viper.GetString("theme.subdued")
	descStyleFg := viper.GetString("theme.verySubdued")
	sepStyleFg := viper.GetString("theme.verySubdued")

	keyStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(keyStyleFg))
	descStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(descStyleFg))
	sepStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(sepStyleFg))

	return help.Styles{
		ShortKey:       keyStyle,
		ShortDesc:      descStyle,
		ShortSeparator: sepStyle,
		Ellipsis:       sepStyle.Copy(),
		FullKey:        keyStyle.Copy(),
		FullDesc:       descStyle.Copy(),
		FullSeparator:  sepStyle.Copy(),
	}
}
