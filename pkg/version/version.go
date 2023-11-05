package version

import (
	"fmt"
	"runtime"
)

const (
	Version     = "1.0.0" // Crare version
	CoreVersion = "5.0.0" // TeleBot version
)

var UA = fmt.Sprintf(
	"Mozilla/5.0 (X11; %v %v; en-GB; compatible; +https://crare.pkg.one) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/119.0.0.0 Safari/537.36 Crare/%v Telebot/%v",
	runtime.GOOS,
	runtime.GOARCH,
	Version,
	CoreVersion,
)
