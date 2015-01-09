package monitod

import(
	"flag"
	_ "github.com/sevlyar/go-daemon"
)

var (
	signal = flag.String("s", "",
		`send a signal to the daemon`
	)
)

//start the daemon with args passed
//https://github.com/sevlyar/go-daemon/blob/master/sample/sample.go
func startDaemon() {
	
}
