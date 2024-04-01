package actions

import (
	"fmt"
	"regexp"
	"sync"

	"github.com/gookit/gcli/v3/progress"
	"github.com/stellaraf/panw-device-state-export/internal/panw"
)

const (
	ExportSuccess = iota
	ExportFailure
)

type ExportResult struct {
	Device panw.Device
	Status int
	File   string
	Error  *error
}

func CollectDeviceStateExport(prog *progress.Progress, rchan chan ExportResult, _wg *sync.WaitGroup, d panw.Device, fu, fp, outDir string) {
	defer _wg.Done()
	defer func() {
		if prog != nil {
			prog.Advance()
		}
	}()
	fw, err := panw.NewFirewall(d, fu, fp)
	if err != nil {
		rchan <- ExportResult{Device: d, Status: ExportFailure, Error: &err}
		return
	}
	outFile, err := fw.ExportDeviceState(outDir)
	if err != nil {
		rchan <- ExportResult{Device: d, Status: ExportFailure, Error: &err}
		return
	}

	rchan <- ExportResult{Device: d, File: outFile, Status: ExportSuccess, Error: nil}
}

func CollectDeviceStateExports(showProgress bool, outDir, hostnamePattern, praHost, praUsername, praPassword, fwUsername, fwPassword string) ([]ExportResult, error) {
	pra, err := panw.NewPanorama(praHost, praUsername, praPassword)
	if err != nil {
		return nil, err
	}

	devices, err := pra.ConnectedDevices()
	if err != nil {
		return nil, err
	}
	var p *progress.Progress
	if showProgress {
		p = progress.Bar(len(devices))
		p.Start()
	}

	pattern := regexp.MustCompile(fmt.Sprintf("^.*%s.*$", hostnamePattern))

	results := make(chan ExportResult, len(devices))
	var wg sync.WaitGroup
	for _, device := range devices {
		if !pattern.MatchString(device.Hostname) {
			if p != nil {
				p.Advance()
			}
			continue
		}
		wg.Add(1)
		go CollectDeviceStateExport(p, results, &wg, device, fwUsername, fwPassword, outDir)
	}
	go func() {
		defer close(results)
		wg.Wait()
	}()
	if p != nil {
		p.Finish()
	}
	resultsSlice := make([]ExportResult, 0, len(devices))
	for result := range results {
		resultsSlice = append(resultsSlice, result)
	}
	return resultsSlice, nil
}
