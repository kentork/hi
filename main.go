package main // package main

import (
	"os"
	"runtime"

	"github.com/martinlindhe/notify"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()

	app.Action = func(c *cli.Context) error {
		appname := "hi"
		if runtime.GOOS == "windows" {
			// https://github.com/go-toast/toast/issues/9
			appname = "{1AC14E77-02E7-4E5D-B744-2EB1AE5198B7}\\WindowsPowerShell\\v1.0\\powershell.exe"
		}

		title := "hi"
		contents := "there"

		if len(c.Args()) >= 2 {
			title = c.Args().Get(0)
			contents = c.Args().Get(1)
		} else if len(c.Args()) >= 1 {
			title = "hi there"
			contents = c.Args().Get(0)
		}

		notify.Notify(appname, title, contents, "dorobou_jikan.png")
		return nil
	}

	app.Name = "hi"
	app.Version = "0.1.0"
	app.Usage = "desktop notifier"
	app.Run(os.Args)
}
