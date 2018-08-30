package watchdog

import (
	"flag"
	"fmt"
	"log"
	"os"
	"time"
)

func Start() error {

	devicePath := flag.String("device", "/dev/watchdog", "Watchdog device")
	devStat, err := os.Stat(*devicePath)
	if err != nil {
		log.Fatal("Can not stat watchdog device:", err)
	}
	if devStat.Mode()&os.ModeDevice == 0 {
		log.Fatal("Watchdog device should be a device file")
	}

	wdt, err := os.OpenFile(*devicePath, os.O_WRONLY, 0)
	if err != nil {
		log.Fatal("Couldn't open watchdog:", err)
		return err
	}
	defer wdt.Close()

	tasty := []byte{100}
	feed := time.Tick(60e9)
	fmt.Printf("Launch watchdog daemon\n")
	for {
		select {
		case <-feed:
			wdt.Seek(0, os.SEEK_SET)
			_, err := wdt.Write(tasty)
			if err != nil {
				log.Fatal("Failed to feed dog :", err)
			}
		}
	}

	return nil
}
