package go_seelog

import log "github.com/cihub/seelog"

func LogOutput() {
	log.Debug("-- debug log --")
	log.Info("-- info log --")
	log.Warn("-- warn log --")
	log.Error("-- error log --")
}
