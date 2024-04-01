package panw

import (
	"fmt"
	"strings"

	"github.com/antchfx/xmlquery"
	"github.com/go-resty/resty/v2"
)

type Panorama struct {
	Host   string
	client *resty.Client
}

type Device struct {
	Hostname     string
	SerialNumber string
	IPAddress    string
	Version      string
	Platform     string
}

func (pra *Panorama) Check() (bool, error) {
	req := pra.client.R().SetQueryParams(map[string]string{
		"type": "op",
		"cmd":  "<show><system><software><status></status></software></system></show>",
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

func (pra *Panorama) ConnectedDevices() ([]Device, error) {
	req := pra.client.R().SetQueryParams(map[string]string{
		"type": "op",
		"cmd":  "<show><devices><connected></connected></devices></show>",
	})
	res, err := req.Get("/api")
	if err != nil {
		return nil, err
	}
	if res.IsError() {
		return nil, fmt.Errorf(res.Status())
	}
	xml, err := xmlquery.Parse(strings.NewReader(res.String()))
	if err != nil {
		return nil, err
	}
	nodes := xmlquery.Find(xml, "//devices/entry")
	devices := make([]Device, 0, len(nodes))
	for _, node := range nodes {
		serial := node.SelectElement("serial").InnerText()
		hostname := node.SelectElement("hostname").InnerText()
		ip := node.SelectElement("ip-address").InnerText()
		version := node.SelectElement("sw-version").InnerText()
		platform := node.SelectElement("model").InnerText()
		device := Device{
			SerialNumber: serial,
			Hostname:     hostname,
			IPAddress:    ip,
			Version:      version,
			Platform:     platform,
		}
		devices = append(devices, device)
	}
	if len(devices) == 0 {
		return nil, fmt.Errorf("no devices found")
	}
	return devices, nil
}

func NewPanorama(host, username, password string) (*Panorama, error) {
	client := resty.New()
	client.SetBasicAuth(username, password)
	client.SetBaseURL(fmt.Sprintf("https://%s", host))
	pra := &Panorama{
		Host:   host,
		client: client,
	}
	_, err := pra.Check()
	if err != nil {
		return nil, err
	}
	return pra, nil
}
