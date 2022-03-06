package aproxiable

import (
	"testing"
)

func TestParseConfig(t *testing.T) {

	t.Log("Testing parsing config from file")

	config, err := parseConfigFromPath("testfiles/config_test.yaml")

	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", config)
}
