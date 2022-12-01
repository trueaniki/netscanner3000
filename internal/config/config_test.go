package config

import (
	"testing"
)

func ParseConfigExample() Config {
	return Parse("../../config-example.yaml")
}

func TestGetDaemon(t *testing.T) {

}

func TestGetScanners(t *testing.T) {
	c := ParseConfigExample()

	for _, s := range c.GetScanners() {
		t.Logf("%+v\n", s)
	}
}

func TestGetObservers(t *testing.T) {
	c := ParseConfigExample()

	for _, o := range c.GetObservers() {
		t.Logf("%+v\n", o)
	}
}
