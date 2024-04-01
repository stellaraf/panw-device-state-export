package panw

import (
	"crypto/tls"
	"fmt"
	"path/filepath"
	"time"

	"github.com/go-resty/resty/v2"
)

type Firewall struct {
	Name   string
	client *resty.Client
	Device Device
}

func (fw *Firewall) Check() (bool, error) {
	req := fw.client.R().SetQueryParams(map[string]string{
		"type": "op",
		"cmd":  "<show><clock></clock></show>",
	})
	res, err := req.Get("/api")
	if err != nil {
		return false, err
	}
	if res.IsError() {
		return false, fmt.Errorf(res.Status())
	}
	return true, nil
}

func (fw *Firewall) ExportDeviceState(dir string) (string, error) {
	fw.client.SetOutputDirectory(dir)
	now := time.Now()
	outFile := filepath.Join(dir,
		fmt.Sprintf(
			"%s_%s_devicestate.tgz",
			now.Format("20060102"),
			fw.Device.Hostname,
		))
	req := fw.client.R().SetOutput(outFile).SetQueryParams(map[string]string{
		"type":     "export",
		"category": "device-state",
	})
	res, err := req.Get("/api")
	if err != nil {
		return "", err
	}
	if res.IsError() {
		return "", fmt.Errorf(res.Status())
	}
	return outFile, nil
}

func NewFirewall(device Device, username, password string) (*Firewall, error) {
	client := resty.New()
	client.SetBasicAuth(username, password)
	client.SetBaseURL(fmt.Sprintf("https://%s", device.IPAddress))
	client.SetTLSClientConfig(&tls.Config{InsecureSkipVerify: true})
	fw := &Firewall{
		Name:   device.Hostname,
		Device: device,
		client: client,
	}
	_, err := fw.Check()
	if err != nil {
		return nil, err
	}
	return fw, nil
}
