package cli

import (
	"github.com/gookit/gcli/v3"
	"github.com/gookit/gcli/v3/gflag"
	"github.com/stellaraf/panw-device-state-export/internal/actions"
)

func CreateDeviceStateExport() *gcli.Command {

	var outDir string
	var hostnamePattern string
	var praHost string
	var praUsername string
	var praPassword string
	var fwUsername string
	var fwPassword string

	cmd := &gcli.Command{
		Name: "device-state-export",
		Desc: "Initiate a Device State export for all connected devices in Panorama",
		Func: func(c *gcli.Command, args []string) error {
			results, err := actions.CollectDeviceStateExports(true, outDir, hostnamePattern, praHost, praUsername, praPassword, fwUsername, fwPassword)
			if err != nil {
				return err
			}
			for _, device := range results {
				if device.Status == actions.ExportSuccess {
					c.Println("SUCCESS:", device.Device.Hostname, device.File)
				}
				if device.Status == actions.ExportFailure {
					c.Println("FAILURE:", device.Device.Hostname)
				}
			}
			return nil
		},
	}
	cmd.StrOpt2(&outDir, "out-directory", "Directory in which device state exports will be written", gflag.WithRequired(), gflag.WithShortcut("o"))
	cmd.StrOpt2(&hostnamePattern, "hostname-pattern", "Regex pattern to match firewall hostnames. Only matched hostnames will have their device states exported.", gflag.WithRequired(), gflag.WithShortcut("p"))
	cmd.StrOpt2(&praHost, "pra-host", "Panorama Hostname", gflag.WithRequired(), gflag.WithShortcut("ph"))
	cmd.StrOpt2(&praUsername, "pra-username", "Panorama Username", gflag.WithRequired(), gflag.WithShortcut("pu"))
	cmd.StrOpt2(&praPassword, "pra-password", "Panorama Password", gflag.WithRequired(), gflag.WithShortcut("pp"))
	cmd.StrOpt2(&fwUsername, "fw-username", "Firewall Username", gflag.WithRequired(), gflag.WithShortcut("fu"))
	cmd.StrOpt2(&fwPassword, "fw-password", "Firewall Password", gflag.WithRequired(), gflag.WithShortcut("fp"))
	return cmd
}

func Run(args []string) {
	app := gcli.NewApp()
	app.Add(CreateDeviceStateExport())
	app.Run(args[1:])
}
