package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"net/http"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/prometheus/common/log"

	"github.com/tarm/serial"
)

type SensorData struct {
	Temperature float64 `json:"temperature"`
	Humidity    float64 `json:"humidity"`
}

var (
	listen = flag.String("listen",
		"0.0.0.0:9612",
		"listen address")
	metricsPath = flag.String("metrics_path",
		"/metrics",
		"path under which metrics are served")
	serialPath = flag.String("serial",
		"/dev/ttyACM0",
		"path to serial device")
)

var (
	temperature = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "temperature_celsius",
		Help: "Current temperature",
	})
	humidity = prometheus.NewGauge(prometheus.GaugeOpts{
		Name: "humidity_percent",
		Help: "Current humidity",
	})
)

func init() {
	prometheus.MustRegister(temperature)
	prometheus.MustRegister(humidity)
}

func watch() {
	c := &serial.Config{Name: *serialPath, Baud: 9600}
	s, err := serial.OpenPort(c)
	if err != nil {
		log.Fatal(err)
	}
	defer s.Close()

	scanner := bufio.NewScanner(s)
	sd := SensorData{}
	for scanner.Scan() {
		line := scanner.Text()

		err = json.Unmarshal([]byte(line), &sd)
		if err != nil {
			log.Error("Can't unmarshar sensors data: ", err)
			log.Error("Received data: ", line)
			continue
		}

		temperature.Set(sd.Temperature)
		humidity.Set(sd.Humidity)
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
}

func main() {
	flag.Parse()

	go watch()

	http.Handle(*metricsPath, promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`<html>
			<head><title>Sensors Exporter</title></head>
			<body>
			<h1>Sensors Exporter</h1>
			<p><a href="` + *metricsPath + `">Metrics</a></p>
			</body></html>`))
	})
	log.Infoln("Listening on", *listen)
	log.Infoln("Serving metrics under", *metricsPath)
	log.Fatal(http.ListenAndServe(*listen, nil))
}
