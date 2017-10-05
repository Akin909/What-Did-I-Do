package main

import ui "github.com/gizak/termui"

//QuitMenu Closes the app
func QuitMenu() (u *ui.Par, ga *ui.Gauge) {
	p := ui.NewPar(":Press q to Quit")
	p.Height = 3
	p.Width = 50
	p.TextFgColor = ui.ColorWhite
	p.BorderLabel = "Latest Commit"
	p.BorderFg = ui.ColorCyan

	g := ui.NewGauge()
	g.Percent = 50
	g.Width = 50
	g.Height = 3
	g.Y = 11
	g.BorderLabel = "Gauge"
	g.BarColor = ui.ColorRed
	g.BorderFg = ui.ColorWhite
	return p, g

}

// ShowGitOutput renders output of a Git command
func ShowGitOutput(s string) (u *ui.Par) {
	p := ui.NewPar(s)
	p.Height = 10
	p.Width = 50
	p.TextFgColor = ui.ColorWhite
	p.BorderFg = ui.ColorCyan

	return p
}
