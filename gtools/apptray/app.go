package app

import (
	"github.com/getlantern/systray"
)

func AppTray() {
	systray.Run(onReady, onExit)
}

func onReady() {
	systray.SetIcon(ico)
	systray.SetTitle("App-Tool")
	systray.SetTooltip("App-Tool-About")

	go func() {
		//AddMenuItem

		systray.AddSeparator()
		mQuit := systray.AddMenuItem("Quit", "Quit app")

		for {
			select {
			case <-mQuit.ClickedCh:
				systray.Quit()
				return
			}
		}
	}()

}

func onExit() {

}
