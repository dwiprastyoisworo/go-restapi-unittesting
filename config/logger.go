package config

import log "github.com/sirupsen/logrus"

func init() {
	// set format log
	log.SetFormatter(&log.JSONFormatter{})
	//set time full log
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})
}

type Logger struct {
	FileName string
	Action   string
}

func NewLogger(fileName string, action string) *Logger {
	return &Logger{
		FileName: fileName,
		Action:   action,
	}
}

// Fungsi SendLogger untuk logging dengan optimasi map level log
func (l Logger) SendLogger(message string, level log.Level, metaData map[string]interface{}) {
	// Tambahkan fields ke log
	logger := log.WithFields(log.Fields{
		"file_name": l.FileName,
		"action":    l.Action,
		"meta_data": metaData,
	})

	// Map level log untuk efisiensi
	levelMap := map[log.Level]func(...interface{}){
		log.InfoLevel:  logger.Info,
		log.WarnLevel:  logger.Warn,
		log.ErrorLevel: logger.Error,
		log.DebugLevel: logger.Debug,
	}

	// Cek level log, jika ada di map
	if logFunc, exists := levelMap[level]; exists {
		logFunc(message)
	} else {
		// Default ke log warning jika level tidak dikenali
		logger.Warn(message)
	}
}
