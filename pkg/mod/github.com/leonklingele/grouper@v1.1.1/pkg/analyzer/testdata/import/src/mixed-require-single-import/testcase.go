package testcase

import (
	"fmt"
	"log"
)

import "sync" // want "should only use a single 'import' declaration, 3 found"
import "time"

func dummy() { fmt.Println("dummy"); log.Println("dummy"); _ = sync.Mutex{}; _ = time.Nanosecond }
