package validation

import (
	"fmt"
	"jsontostruct/ironic"
)

// Valid returns true or false according to the test condition
func Valid(hostList ironic.Data, expectedHardwareConfig ironic.ExpectedHardwareConfiguration) []ironic.Host {
	var validHost []ironic.Host
	// Your Logic here
	noOfHosts := len(hostList.Host)
	fmt.Println("Number OF Hosts to Compare:", noOfHosts)

	n := 0
	for n = 0; n < noOfHosts; n++ {
		if expectedHardwareConfig.ExpectedCPU.Count < hostList.Host[n].Spec.HardwareDetails.CPU.Count {
			fmt.Println("CPU Valid, Now Check Memory")
			if expectedHardwareConfig.ExpectedDisk.SizeBytesGB < hostList.Host[n].Spec.HardwareDetails.Storage.SizeBytes*1024*1024 {
				fmt.Println("Memory Valid, Now Check Nics")
				if expectedHardwareConfig.ExpectedNICS.NumberOfNICS < hostList.Host[n].Spec.HardwareDetails.Nics.NoOfNics {
					fmt.Println("Number of Nics are avialale, Now Check RAM")
					if expectedHardwareConfig.ExpectedRAM < hostList.Host[n].Spec.HardwareDetails.Memory.PhysicalMb {
						fmt.Println("RAM is avialale, Now Check Vendor Name")
						if expectedHardwareConfig.ExpectedSystemVendor.Name == hostList.Host[n].Spec.HardwareDetails.SystemVendor.Manufacturer {
							fmt.Println("Host Found")
							validHost = append(validHost, hostList.Host[n])
						} else {
							fmt.Println("System vendor different, Check next Host")
							continue
						}
					} else {
						fmt.Println("Not sufficient RAM, Check next Host")
						continue
					}
				} else {
					fmt.Println("Not sufficient Nics, Check next Host")
					continue
				}
			} else {
				fmt.Println("Not sufficient memory, Check next Host")
				continue
			}
		} else {
			fmt.Println("CPU Invalid, Check next Host")
			continue
		}
		//	fmt.Println("Inside Validation function", expectedHardwareConfig.ExpectedCPU.Count)
		//	fmt.Println("Storage in bytes count in first Host:", hostList.Host[n].Spec.HardwareDetails.Storage.SizeBytes)
		//	fmt.Println("Storage in bytes count in first Host:", hostList.Host[n].Spec.HardwareDetails.Storage.NumberOfDisks)
		//	fmt.Println("Nics count in first Host:", hostList.Host[n].Spec.HardwareDetails.Nics.NoOfNics)

	}
	return validHost
}
