/*

  Basic example of how to use the door library.

  Note that you must have a USB device connected to one of the available ports.

*/
package main

import (
	"fmt"

	"github.com/chimera/door"
)

func main() {

	// Create a new door lock instance.
	d := door.NewDoorLock()

	// Unlock door
	err := d.Unlock()
	if err != nil {
		fmt.Printf("Error unlocking door! %s", err.Error())
	}
}
