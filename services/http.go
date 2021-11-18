package services

import (
	"log"
	"net/http"
	"strconv"

	"github.com/wbwax/wax-engine/g"
	"github.com/wbwax/wax-engine/utils"

	"github.com/wbwax/logger"
)

// StartHTTP starts http server and pprof
func StartHTTP(cfg g.HTTPConfig) {
	if !cfg.Enable {
		logger.Infof("msg=%s", "http is disabled")
		return
	}

	portStr := strconv.FormatInt(int64(cfg.Port), 10)
	address := utils.JoinStrWithSep(":", "0.0.0.0", portStr)

	s := &http.Server{
		Addr:           address,
		MaxHeaderBytes: 1 << 30,
	}
	s.SetKeepAlivesEnabled(true) // By default, keep-alives are always enabled

	logger.Infof("msg=%s||address=%s", "http server will start", address)
	if err := s.ListenAndServe(); err != nil { // will block when success
		logger.Errorf("msg=%s||err=%s", "failed to start http server", err.Error())
		log.Fatalln(err)
	}
}
