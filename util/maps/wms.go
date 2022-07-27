package util

type WMSParams struct {
	Service     string `form:"service"`
	Version     string `form:"version"`
	Request     string `form:"request"`
	Layers      string `form:"layers"`
	BBox        string `form:"bbox"`
	Width       string `form:"width"`
	Height      string `form:"height"`
	SRS         string `form:"srs"`
	Styles      string `form:"styles"`
	Format      string `form:"format"`
	Transparent string `form:"transparent"`
}
