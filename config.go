package main

import (
	"encoding/json"
	"io/ioutil"
)

type PumpConfig struct {
	Name string                 `json:"name"`
	Meta map[string]interface{} `json:"meta"`
}

type TykPumpConfiguration struct {
	PurgeDelay             int                   `json:"purge_delay"`
	DontPurgeUptimeData    bool                  `json:"dont_purge_uptime_data"`
	UptimePumpConfig       interface{}           `json:"uptime_pump_config"`
	Pumps                  map[string]PumpConfig `json:"pumps"`
	AnalyticsStorageType   string                `json:"analytics_storage_type"`
	AnalyticsStorageConfig interface{}           `json:"analytics_storage_config"`
}

func LoadConfig(filePath *string, configStruct *TykPumpConfiguration) {
	configuration, err := ioutil.ReadFile(*filePath)
	if err != nil {
		log.Fatal("Couldn't load configuration file: ", err)

	}

	marshalErr := json.Unmarshal(configuration, &configStruct)
	if marshalErr != nil {
		log.Fatal("Couldn't unmarshal configuration: ", marshalErr)
	}
}
