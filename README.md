# sensors_exporter

Simple golang application that reads temperature/humidity from Arduino (via serial interface) and export them as Prometheus metrics.

```
                      +----------------------+
+--------------+      | +------------------+ |      +---------+
|              |      | |                  | |      |         |      +--------+
|  Prometheus  <--------+ sensors_exporter <--------+ Arduino <------+ DHT 22 |
|              |      | |                  | |      |         |      +--------+
+--------------+      | +------------------+ |      +---------+ 
                      +----------------------+
                                  Raspberry Pi
```

## Usage

```
Usage of sensors_exporter:
  -listen string
    	listen address (default "localhost:9612")
  -metrics_path string
    	path under which metrics are served (default "/metrics")
  -serial string
    	path to serial device (default "/dev/ttyACM0")
```
