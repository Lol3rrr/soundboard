package devices

import "github.com/gen2brain/malgo"

// GetDeviceNames loads all the Devices of a given Type and returns their names
func GetDeviceNames(ctx *malgo.AllocatedContext, deviceType malgo.DeviceType) []string {
	result := make([]string, 0)

	devices, err := ctx.Devices(deviceType)
	if err != nil {
		return result
	}

	for _, tmpDevice := range devices {
		result = append(result, tmpDevice.Name())
	}

	return result
}
