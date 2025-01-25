package main

import (
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"

	//"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func main() {
	myApp := app.New()
	myWindow := myApp.NewWindow("TabContainer Widget")

	tabs := container.NewAppTabs(
		container.NewTabItem("Tab 1", widget.NewLabel("Hello")),
		container.NewTabItem("Tab 2", widget.NewLabel("World!")),
		container.NewTabItem("Tab 3", widget.NewLabel("This")),
		container.NewTabItem("Tab 4", widget.NewLabel("is")),
		container.NewTabItem("Tab 5", widget.NewLabel("a")),
		container.NewTabItem("Tab 6", widget.NewLabel("set")),
		container.NewTabItem("Tab 7", widget.NewLabel("of")),
		container.NewTabItem("Tab 8", widget.NewLabel("Tabs")),
	)

	//tabs.Append(container.NewTabItemWithIcon("Home", theme.HomeIcon(), widget.NewLabel("Home tab")))

	tabs.SetTabLocation(container.TabLocationLeading)

	myWindow.SetContent(tabs)
	myWindow.ShowAndRun()
}
