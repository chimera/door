package door

import (
	"fmt"
	"os"
	"os/signal"
	"time"

	"github.com/davecheney/gpio"
	"github.com/davecheney/gpio/rpi"
)

const (
	openDelay = 6
)

type doorlock struct {
	pin gpio.Pin
}

func (d *doorlock) connect() (err error) {

	// Lock door on exit
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	go func() {
		for _ = range c {
			fmt.Printf("\nClearing and unexporting the pin.\n")
			d.disconnect()
		}
	}()

	// Create a connection with the door.
	pin, err := gpio.OpenPin(rpi.GPIO25, gpio.ModeOutput)
	if err != nil {
		return fmt.Errorf("Error opening pin! %s\n", err)
	}

	d.pin = pin
	return nil
}

func (d *doorlock) disconnect() {
	d.pin.Clear()
	d.pin.Close()
	os.Exit(0)
}

func (d *doorlock) Unlock() (err error) {

	// Connect to the door
	err = d.connect()
	if err != nil {
		return err
	}

	// Open the door
	d.pin.Set()

	// Wait to lock the door
	time.Sleep(openDelay * time.Second)

	// Lock the door after the delay
	d.pin.Clear()

	return nil
}

func NewDoorLock() doorlock {
	return doorlock{}
}
