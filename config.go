package mxload

import "strings"

type MxloadConfig struct {
	Enabled       bool
	CollectorURL  string
	ProjectID     int
	ServiceKey    string
	CustomHeaders []string
	IgnorePaths   []string
}

func (c *MxloadConfig) IsCustomHeader(header string) bool {
	for _, h := range c.CustomHeaders {
		if strings.EqualFold(h, header) {
			return true
		}
	}
	return false
}

func (c *MxloadConfig) IsIgnoredPath(path string) bool {
	for _, p := range c.IgnorePaths {
		if strings.Contains(path, p) {
			return true
		}
	}
	return false
}
