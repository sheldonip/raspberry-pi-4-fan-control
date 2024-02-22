# raspberry-pi-4-fan-control
Small script for controlling a pwm fan of raspberry pi 4 

## Wiring
- Physical pin 4 to 5V
- Physical pin 6 to ground
- Physical pin 8 (GPIO 14, TXD) to pwm

## Before running
```sh
sudo apt-get install golang
go mod init fan-control.go
go get github.com/stianeikeland/go-rpio/v4
```

## Run
```sh
go run fan-control.go
```
