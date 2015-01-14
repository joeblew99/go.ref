package rt

import (
	"os"
	"os/signal"
	"syscall"

	"v.io/core/veyron2/vlog"
)

func (r *vrt) initSignalHandling() {
	// TODO(caprita): Given that our device manager implementation is to
	// kill all child apps when the device manager dies, we should
	// enable SIGHUP on apps by default.

	// Automatically handle SIGHUP to prevent applications started as
	// daemons from being killed.  The developer can choose to still listen
	// on SIGHUP and take a different action if desired.
	r.signals = make(chan os.Signal, 1)
	signal.Notify(r.signals, syscall.SIGHUP)
	go func() {
		for {
			sig, ok := <-r.signals
			if !ok {
				break
			}
			vlog.Infof("Received signal %v", sig)
		}
	}()
}

func (r *vrt) shutdownSignalHandling() {
	signal.Stop(r.signals)
	close(r.signals)
}
