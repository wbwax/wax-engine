package g

import (
	"github.com/wbwax/logger"
)

func InitLog(cfg *logCfg) error {
	config := logger.Config{
		MaxSize:    cfg.MaxSize,
		MaxAge:     cfg.MaxAge,
		MaxBackups: cfg.MaxBackups,
		Level:      cfg.Level,
		Path:       cfg.Path,
		Encoding:   cfg.Encoding,
	}
	logger.Init(config)
	defer logger.Sync() // flush buffer, if any
	return nil
}
