module github.com/aibotsoft/latency-metric

go 1.17

require (
	github.com/adshao/go-binance/v2 v2.3.4
	github.com/aibotsoft/ftx-api v0.0.0-20220218163717-ebf7b8ba8402
	github.com/cristalhq/aconfig v0.16.8
	github.com/cristalhq/aconfig/aconfigyaml v0.16.1
	github.com/tcnksm/go-httpstat v0.2.0
	go.uber.org/zap v1.21.0
)

require (
	github.com/bitly/go-simplejson v0.5.0 // indirect
	github.com/gorilla/websocket v1.5.0 // indirect
	github.com/klauspost/compress v1.13.1 // indirect
	go.uber.org/atomic v1.9.0 // indirect
	go.uber.org/multierr v1.7.0 // indirect
	gopkg.in/yaml.v2 v2.3.0 // indirect
	nhooyr.io/websocket v1.8.7 // indirect
)

replace github.com/aibotsoft/ftx-api => ../ftx-api
