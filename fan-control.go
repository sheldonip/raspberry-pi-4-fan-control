package main

import (
	"bufio"
	"io"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"github.com/stianeikeland/go-rpio/v4"
)

func ReadInts(r io.Reader) ([]int, error) {
	scanner := bufio.NewScanner(r)
	scanner.Split(bufio.ScanWords)
	var result []int
	for scanner.Scan() {
		x, err := strconv.Atoi(scanner.Text())
		if err != nil {
			return result, err
		}
		result = append(result, x)
	}
	return result, scanner.Err()
}

func main() {
	lowerTemperatureThreshold := float64(36)
	upperTemperatureThreshold := float64(48)
	for {
		data, err := os.ReadFile("/sys/class/thermal/thermal_zone0/temp")
		if err != nil {
			log.Fatal(err)
		}

		ints, err := ReadInts(strings.NewReader(string(data)))
		if err != nil {
			log.Fatal(err)
		}
		temperature := float64(ints[0]) / 1000

		err = rpio.Open()
		if err != nil {
			log.Fatal(err)
		}

		txd := rpio.Pin(14)
		txd.Output()
		if lowerTemperatureThreshold >= temperature {
			txd.Low()
		}
		if temperature >= upperTemperatureThreshold {
			txd.High()
		}
		rpio.Close()
		time.Sleep(5 * time.Second)
	}
}
