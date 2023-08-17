package timeular

import (
	"context"
	"errors"
	"fmt"

	"tinygo.org/x/bluetooth"
)

// Timeular represents a device contract.
type Timeular interface {
	// GetOrientation returns the current device orientation.
	GetOrientation() (uint8, error)
	// GetBatteryLevel returns the current device battery level.
	GetBatteryLevel() (uint8, error)
	// OnOrientationChanged registers a callback for the actor which is triggered
	// by an orientation change notification.
	OnOrientationChanged(func(uint8))
	// OnBatteryLevelChanged registers a callback for the actor which is triggered
	// by battery level change notification.
	OnBatteryLevelChanged(func(uint8))
	// Run the event monitor.
	Run(ctx context.Context) error
}

var (
	serviceUUIDTimeularOrientation        = bluetooth.NewUUID([16]byte{0xc7, 0xe7, 0x00, 0x10, 0xc8, 0x47, 0x11, 0xe6, 0x81, 0x75, 0x8c, 0x89, 0xa5, 0x5d, 0x40, 0x3c})
	characteristicUUIDTimeularOrientation = bluetooth.NewUUID([16]byte{0xc7, 0xe7, 0x00, 0x12, 0xc8, 0x47, 0x11, 0xe6, 0x81, 0x75, 0x8c, 0x89, 0xa5, 0x5d, 0x40, 0x3c})
)

// New timeular device driver to handle common device services.
func New(device *bluetooth.Device) (Timeular, error) {
	// Check arguments
	if device == nil {
		return nil, errors.New("bluetooth device must not be nil")
	}

	// Discover required services
	services, err := device.DiscoverServices([]bluetooth.UUID{
		bluetooth.ServiceUUIDBattery,
		serviceUUIDTimeularOrientation,
	})
	switch {
	case err != nil:
		return nil, fmt.Errorf("unable to retrieve device services: %w", err)
	case len(services) != 2:
		return nil, fmt.Errorf("retrieved services count is not as expected got %d: %w", len(services), err)
	}

	return &timeularDevice{
		device:                      device,
		orientationChangedCallback:  nil,
		batteryLevelChangedCallback: nil,
		batteryService:              services[0],
		orientionService:            services[1],
	}, nil
}

type timeularDevice struct {
	device                      *bluetooth.Device
	orientationChangedCallback  func(uint8)
	batteryLevelChangedCallback func(uint8)
	orientionService            bluetooth.DeviceService
	batteryService              bluetooth.DeviceService
}

func (td *timeularDevice) GetOrientation() (uint8, error) {
	characteristic, err := td.getOrientationCharacteristic()
	if err != nil {
		return 0, fmt.Errorf("unable to read orientation characteristic: %w", err)
	}

	out := [1]byte{}
	if err := td.readValue(characteristic, out[:]); err != nil {
		return 0, fmt.Errorf("unable to read orientation value: %w", err)
	}

	return uint8(out[0]), nil
}

func (td *timeularDevice) GetBatteryLevel() (uint8, error) {
	characteristic, err := td.getBatteryLevelCharacteristic()
	if err != nil {
		return 0, fmt.Errorf("unable to read battery level characteristic: %w", err)
	}

	out := [1]byte{}
	if err := td.readValue(characteristic, out[:]); err != nil {
		return 0, fmt.Errorf("unable to read battery level value: %w", err)
	}

	return uint8(out[0]), nil
}

func (td *timeularDevice) OnOrientationChanged(callback func(uint8)) {
	td.orientationChangedCallback = callback
}

func (td *timeularDevice) OnBatteryLevelChanged(callback func(uint8)) {
	td.batteryLevelChangedCallback = callback
}

func (td *timeularDevice) Run(ctx context.Context) error {
	// Monitor orientation changes
	if td.orientationChangedCallback != nil {
		// Subscribe to event
		orientation, err := td.getOrientationCharacteristic()
		if err != nil {
			return fmt.Errorf("unable to retrieve orientation characteristic: %w", err)
		}

		if err := orientation.EnableNotifications(func(buf []byte) {
			if len(buf) == 1 {
				td.orientationChangedCallback(uint8(buf[0]))
			}
		}); err != nil {
			return fmt.Errorf("unable to register orientation notification handler: %w", err)
		}
	}

	// Monitor battery level changes
	if td.batteryLevelChangedCallback != nil {
		// Subscribe to event
		battery, err := td.getBatteryLevelCharacteristic()
		if err != nil {
			return fmt.Errorf("unable to retrieve battery level characteristic: %w", err)
		}

		if err := battery.EnableNotifications(func(buf []byte) {
			if len(buf) == 1 {
				td.batteryLevelChangedCallback(uint8(buf[0]))
			}
		}); err != nil {
			return fmt.Errorf("unable to register battery level notification handler: %w", err)
		}
	}

	select {
	case <-ctx.Done():
		return ctx.Err()
	}
}

// -----------------------------------------------------------------------------

func (td *timeularDevice) getBatteryLevelCharacteristic() (*bluetooth.DeviceCharacteristic, error) {
	chars, err := td.batteryService.DiscoverCharacteristics([]bluetooth.UUID{
		bluetooth.CharacteristicUUIDBatteryLevel,
	})
	switch {
	case err != nil:
		return nil, fmt.Errorf("unable to retrieve characteristic from the device: %w", err)
	case len(chars) != 1:
		return nil, fmt.Errorf("retrieved characteristic count is not as expected got %d: %w", len(chars), err)
	}

	// Extract characteristic
	char := chars[0]

	return &char, nil
}

func (td *timeularDevice) getOrientationCharacteristic() (*bluetooth.DeviceCharacteristic, error) {
	chars, err := td.orientionService.DiscoverCharacteristics([]bluetooth.UUID{
		characteristicUUIDTimeularOrientation,
	})
	switch {
	case err != nil:
		return nil, fmt.Errorf("unable to retrieve characteristic from the device: %w", err)
	case len(chars) != 1:
		return nil, fmt.Errorf("retrieved characteristic count is not as expected got %d: %w", len(chars), err)
	}

	// Extract characteristic
	char := chars[0]

	return &char, nil
}

func (td *timeularDevice) readValue(char *bluetooth.DeviceCharacteristic, out []byte) error {
	// It looks like the characteristic doesn't implement correctly the io.Reader interface - https://github.com/tinygo-org/bluetooth/pull/136

	// Read characteristic
	if _, err := char.Read(out[:]); err != nil {
		return err
	}

	return nil
}
