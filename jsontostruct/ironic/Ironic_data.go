package ironic

// Data ironic introspected data
type Data struct {
	Host []Host `json:"host"`
}

// Host Spec of the host
type Host struct {
	Spec Spec `json:"Spec"`
}

// Spec hardware details for the spec
type Spec struct {
	HardwareDetails HardwareDetails `json:"HardwareDetails"`
}

// HardwareDetails config details
type HardwareDetails struct {
	Boot         Boot         `json:"boot"`
	SystemVendor SystemVendor `json:"system_vendor"`
	Memory       Memory       `json:"memory"`
	CPU          CPU          `json:"cpu"`
	Storage      Storage      `json:"storage"`
	Nics         Nics         `json:"nics"`
}

// Boot boot details
type Boot struct {
	CurrentBootMode string `json:"current_boot_mode"`
	PxeInterface    string `json:"pxe_interface"`
}

// SystemVendor vendor details
type SystemVendor struct {
	SerialNumber string `json:"serial_number"`
	ProductName  string `json:"product_name"`
	Manufacturer string `json:"manufacturer"`
}

// CPU cpu details
type CPU struct {
	Count        int      `json:"count"`
	Frequency    string   `json:"frequency"`
	Flags        []string `json:"flags"`
	ModelName    string   `json:"model_name"`
	Architecture string   `json:"architecture"`
}

// Memory memory details
type Memory struct {
	PhysicalMb int   `json:"physical_mb"`
	Total      int64 `json:"total"`
}

// Storage details
type Storage struct {
	SizeBytes     int64 `json:"size_bytes"`
	NumberOfDisks int64 `json:"noOfDisks"`
}

// NICS count
type Nics struct {
	NoOfNics int `json:"noOfNics"`
}

// ExpectedHardwareConfiguration details to match with the host
type ExpectedHardwareConfiguration struct {
	ExpectedCPU          ExpectedCPU          `json:"expectedCPU"`
	ExpectedDisk         ExpectedDisk         `json:"expectedDisk"`
	ExpectedNICS         ExpectedNICS         `json:"expectedNICS"`
	ExpectedRAM          int                  `json:"expectedRAM"`
	ExpectedSystemVendor ExpectedSystemVendor `json:"expectedSystemVendor"`
}

// ExpectedCPU expected cpu count
type ExpectedCPU struct {
	Count int `json:"count"`
}

// ExpectedDisk size and number of disk needed
type ExpectedDisk struct {
	SizeBytesGB   int64 `json:"sizeBytesGB"`
	NumberOfDisks int   `json:"numberOfDisks"`
}

// ExpectedNICS count of nics cards
type ExpectedNICS struct {
	NumberOfNICS int `json:"numberOfNICS"`
}

// ExpectedSystemVendor Vendor details
type ExpectedSystemVendor struct {
	Name string `json:"name"`
}
