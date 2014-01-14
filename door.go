package door

import (
	"fmt"

	"github.com/chimera/rs232"
)

type doorlock struct {
	baud  int
	ports []string
}

func (d *doorlock) connect(port string) (conn *rs232.Port, err error) {

	// fmt.Println("Connecting to door lock...")

	// Configure the serial connection.
	options := rs232.Options{
		BitRate:  uint32(d.baud),
		DataBits: 8,
		StopBits: 1,
		Parity:   rs232.PARITY_NONE,
		Timeout:  0,
	}

	// Open a connection to the serial port.
	conn, err = rs232.Open(port, options)
	if err != nil {
		return &rs232.Port{}, fmt.Errorf("Could not connect to port %s", port)
	}

	return conn, nil
}

func (d *doorlock) disconnect(conn *rs232.Port) {
	// fmt.Println("Disconnecting from door lock...")
	conn.Close()
}

func (d *doorlock) Unlock() (err error) {

	// Loop over the available ports and try to connect in order.
	for _, port := range d.ports {

		// Attempt to connect to the given serial port.
		conn, err := d.connect(port)
		if err != nil {
			// fmt.Println(err)
			continue
		}
		defer d.disconnect(conn)

		// Attempt to unlock door
		_, err = conn.Write([]byte("1"))
		if err != nil {
			return fmt.Errorf("Could not unlock door, with error: %s", err)
		}

		return nil
	}

	// None of the expected ports could be connected to.
	return fmt.Errorf("Failed to connect to all available ports!")
}

func NewDoorLock() doorlock {
	return doorlock{
		baud: 19200,
		// TODO: This shouldn't be hard coded, ideally.
		ports: []string{
			"/dev/ttyACM0",
			"/dev/ttyACM1",
			"/dev/ttyACM2",
			"/dev/tty.usbmodem411",
			"/dev/tty.usbmodem621",
		},
	}
}
