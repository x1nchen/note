package main

import (
	"encoding/base64"
	"log"

	giDevice "github.com/electricbubble/gidevice"
)

func main() {
	// libimobiledevice.SetDebug(true)
	usbmux, err := giDevice.NewUsbmux()
	if err != nil {
		log.Fatalln(err)
	}

	devices, err := usbmux.Devices()
	if err != nil {
		log.Fatalln(err)
	}

	if len(devices) == 0 {
		log.Fatalln("No Device")
	}

	d := devices[0]

	imageSignatures, err := d.Images()
	if err != nil {
		log.Fatalln(err)
	}

	for i, imgSign := range imageSignatures {
		log.Printf("[%d] %s\n", i+1, base64.StdEncoding.EncodeToString(imgSign))
	}

	// dmgPath := "/Applications/Xcode.app/Contents/Developer/Platforms/iPhoneOS.platform/DeviceSupport/14.0/DeveloperDiskImage.dmg"
	// signaturePath := "/Applications/Xcode.app/Contents/Developer/Platforms/iPhoneOS.platform/DeviceSupport/14.0/DeveloperDiskImage.dmg.signature"
	//
	// err = d.MountDeveloperDiskImage(dmgPath, signaturePath)
	// if err != nil {
	// 	log.Fatalln(err)
	// }
	// log.Println("mount developer disk image success")

	// https://developer.amap.com/tools/picker
	// https://lbs.qq.com/tool/getpoint/index.html
	// if err = d.SimulateLocationUpdate(113.925734, 22.511836, giDevice.CoordinateSystemWGS84); err != nil {
	// 	log.Fatalln(err)
	// }

	err = d.SimulateLocationRecover()
	if err != nil {
		log.Fatalln(err)
	}
	log.Println("simul loc recover")
}

