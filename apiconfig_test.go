package apiconfig

import (
	"encoding/json"
	"fmt"
	"testing"
)

type TestConfig struct {
	Name    string `json:"name"`
	Address string `json:"address"`
}

func TestSample(t *testing.T) {
	tc := TestConfig{
		Name:    "Test",
		Address: "Sample",
	}
	tcData, err := json.Marshal(tc)
	if err != nil {
		t.Fatal("Failed to create config")
	}
	err = CreateAPIConfig("api_config.json", "TestKey", tcData)
	if err != nil {
		t.Fatal("Failed to create config")
	}
	var c TestConfig
	err = LoadAPIConfig("api_config.json", "TestKey", &c)
	if err != nil {
		t.Fatal("Failed to load config")
	}
	if c.Name != tc.Name || c.Address != tc.Address {
		t.Fatal("Value mismatch")
	}
}

func TestSimple(t *testing.T) {

	var c map[string]interface{}
	err := LoadAPIConfig("api_config.json", "TestKey@Prod#$erver^2021", &c)
	if err != nil {
		t.Fatal("Failed to load config")
	}
	fmt.Printf("Config : %v\n", c)
}
