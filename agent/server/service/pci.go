package service

import (
	"fmt"
	"github.com/jaypipes/pcidb"
)

func main() {
	pci, err := pcidb.New()
	if err != nil {
		fmt.Printf("Error getting PCI info: %v", err)
	}

	for _, devClass := range pci.Classes {
		fmt.Printf(" Device class: %v ('%v')\n", devClass.Name, devClass.ID)
		for _, devSubclass := range devClass.Subclasses {
			fmt.Printf("    Device subclass: %v ('%v')\n", devSubclass.Name, devSubclass.ID)
			for _, progIface := range devSubclass.ProgrammingInterfaces {
				fmt.Printf("        Programming interface: %v ('%v')\n", progIface.Name, progIface.ID)
			}
		}
	}
}
