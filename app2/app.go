package app2

import (
	"fmt"
	"io"

	"github.com/kudrykv/latex-yearly-planner/app2/flags"
	"github.com/urfave/cli/v2"
)

type App struct {
	app        *cli.App
	deviceFlag *flags.DeviceFlag
}

func New(reader io.Reader, writer, errWriter io.Writer) *App {
	return (&App{}).
		prepareFlags().
		setupCli(reader, writer, errWriter)
}

func (r *App) prepareFlags() *App {
	r.deviceFlag = flags.NewDeviceFlag()

	return r
}

func (r *App) setupCli(reader io.Reader, writer, errWriter io.Writer) *App {
	r.app = &cli.App{
		Name: "plannergen",

		Reader:    reader,
		Writer:    writer,
		ErrWriter: errWriter,
		Flags:     r.flags(),

		Action: r.mainAction,
	}

	return r
}

func (r *App) flags() []cli.Flag {
	return []cli.Flag{
		r.deviceFlag,
	}
}

func (r *App) mainAction(appContext *cli.Context) error {
	// dimensions:
	// - device (supernote, remarkable (vanilla, ddvk), ?boox)
	//     this defines dimensions and internal layout boundaries: leave empty space for control elements
	// - template - which template to use
	//     sub-dependent are enabled template sections
	//     based on what is selected, some links should or should not be displayed, etc
	device, err := r.deviceFlag.Device()
	if err != nil {
		return fmt.Errorf("get device: %w", err)
	}

	_ = device

	// get / parse configs
	// create tex files
	// run latex

	return nil
}

func (r App) Run(args []string) error {
	return r.app.Run(args)
}