package go_seelog

import log "github.com/cihub/seelog"

var Logger log.LoggerInterface

func init() {
	Logger, err := log.LoggerFromConfigAsFile("./config/seelog_config.xml")
	if err != nil {
		log.Error("Fatal error ", err.Error())
		return
	}
	log.ReplaceLogger(Logger)
	Logger.Info("-- 配置文件加载成功 --")
}
