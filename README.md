# Mqtt-cli
This a simple mqtt-cli application written in [Go](https://golang.org/). It has pub sub command which is suite for the shell script pipelining.

# Install
If you have golang ready environment, you can easily install it
```
go get github.com/shuza/mqtt-cli
```

# Usage
You can set host, port, clientId, topic and qos on the Environment variables.
```
export MQTT_HOST="localhost"
export MQTT_PORT="1883"
export MQTT_CLIENT_ID="mqtt-cli"
export MQTT_TOPIC="/test/mqtt/cli"
export MQTT_QOS="1"
```
Or you can use flags to do it.

|   Name   |    Short Hand |    Description     |
|:----------:|:-------------:|:------------:|
| address    |      a        | Set MQTT host address |
| port       |      p        | Set MQTT host port    |
| clientId   |      i        | Set MQTT client ID    |
| topic      |      t        | Set MQTT topic        |
| qos        |      q        | Set MQTT service quality |
| message    |      m        | Set MQTT message palyload |
 
# Sub
```
mqtt-cli sub -a "localhost" -p 1883 -t "/tmp/a"
```

# Pub
```
mqtt-cli pub -a "localhost" -p 1883 -t "/tmp/a" -m "thsi is to test"
```
![Solid](https://media.giphy.com/media/JpGXXxiamqV53gqleB/giphy.gif)
