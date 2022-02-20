package config

import (
	"github.com/cristalhq/aconfig"
	"github.com/cristalhq/aconfig/aconfigyaml"
)

type Config struct {
	Count   int `json:"count" flag:"c"`
	Service struct {
		Name string `json:"name"`
	} `json:"service"`
	Zap struct {
		//debug, info, warn, error, fatal, panic
		Level string `json:"level" default:"info"`
		//console, json
		Encoding string `json:"encoding" default:"console"`
		//disable, short, full
		Caller string `json:"caller" default:"disable"`
	} `json:"zap"`
	Binance struct {
		//Name   string `json:"name" default:"binance"`
		//Debug  bool   `json:"debug" default:"false"`
		Key    string `json:"-"`
		Secret string `json:"-"`
	} `json:"binance"`
	Ftx struct {
		//Name string `json:"name" default:"ftx"`
		//WsHost string `default:"wss://ftx.com/ws/"`
		Key    string `json:"-"`
		Secret string `json:"-"`
	} `json:"ftx"`
}

func NewConfig() *Config {
	var cfg Config
	loader := aconfig.LoaderFor(&cfg, aconfig.Config{
		//SkipFlags:          true,
		AllFieldRequired:   true,
		AllowUnknownFlags:  true,
		AllowUnknownEnvs:   true,
		AllowUnknownFields: true,
		AllowDuplicates:    true,
		SkipEnv:            false,
		FileFlag:           "config",
		FailOnFileNotFound: false,
		MergeFiles:         true,
		Files:              []string{"config.yaml", "latency-metric.yaml"},
		FileDecoders: map[string]aconfig.FileDecoder{
			".yaml": aconfigyaml.New(),
		},
	})
	err := loader.Load()
	if err != nil {
		panic(err)
	}

	return &cfg
}
