package database

import (
	"sync"
	"time"
)

var (
	rackBobinaMap = make(map[string][]string)
	mapMutex      = &sync.Mutex{}
)

func AddBobinaToRack(hexRack, hexBobina string) {
	mapMutex.Lock()
	rackBobinaMap[hexRack] = append(rackBobinaMap[hexRack], hexBobina)
	mapMutex.Unlock()
}

func InitPeriodicCleanup(interval time.Duration) {
	go func() {
		for {
			time.Sleep(interval)
			mapMutex.Lock()
			rackBobinaMap = make(map[string][]string) // Limpa todo o mapa
			mapMutex.Unlock()
		}
	}()
}
