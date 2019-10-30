package test

import (
	"flag"
	"testing"

	"github.com/golang/glog"
)

// glog write log file to local
func TestGolog(t *testing.T) {
	flag.Set("v", "1")
	// Create logfiles folder first, then write log to local dir
	flag.Set("log_dir", "../log")
	flag.Parse() // 1

	glog.Info("This is a Info log") // 2
	glog.Warning("This is a Vincent Warning log")
	glog.Error("This is a Error log")

	glog.V(1).Infoln("level 1") // 3
	glog.V(2).Infoln("level 2")

	glog.Flush() // 4
}
