package domain

type DevicePhysicalStatus string

const (
	DevicePhysicalStatusOperational DevicePhysicalStatus = "OPT"
	DevicePhysicalStatusMaintenance DevicePhysicalStatus = "MNT"
	DevicePhysicalStatusDamaged     DevicePhysicalStatus = "DMG"
	DevicePhysicalStatusRetired     DevicePhysicalStatus = "RTD"
)

type Device struct {
	IMEI            string
	Model           string
	FirmwareVersion string
	BatteryLevel    uint8
	PhysicalStatus  DevicePhysicalStatus
}
