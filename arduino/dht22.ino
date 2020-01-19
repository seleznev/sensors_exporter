#include <DHT.h>;

#define DHTPIN 7 // Arduino pin we're connected to
#define DHTTYPE DHT22 // DHT 22 (AM2302)
DHT dht(DHTPIN, DHTTYPE); // Initialize DHT sensor for normal 16mhz Arduino

float humidity;
float temperature;

void setup() {
  Serial.begin(9600);
  dht.begin();
}

void loop() {
  delay(10000); // 10s
  
  // Read data
  temperature = dht.readTemperature();
  humidity = dht.readHumidity();
  
  // Print temperature and humidity values to serial monitor
  Serial.print("{\"temperature\": ");
  Serial.print(temperature);
  Serial.print(", \"humidity\": ");
  Serial.print(humidity);
  Serial.println("}");
}
