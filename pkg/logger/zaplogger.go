package logger

// var Log *zap.Logger

// func InitLogger() {
// 	encoderConfig := zap.NewProductionEncoderConfig()
// 	encoderConfig.EncodeTime = zapcore.ISO8601TimeEncoder

// 	fileEncoder := zapcore.NewJSONEncoder(encoderConfig)

// 	consoleConfig := encoderConfig
// 	consoleConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
// 	consoleEncoder := zapcore.NewConsoleEncoder(consoleConfig)

// 	logFile, err := os.OpenFile("ZapLogger.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
// 	if err != nil {
// 		panic("Failed to open log file: " + err.Error())
// 	}

// 	fileWriter := zapcore.AddSync(logFile)
// 	consoleWriter := zapcore.AddSync(os.Stdout)

// 	defaultLevel := zap.NewAtomicLevelAt(zap.InfoLevel)

// 	core := zapcore.NewTee(
// 		zapcore.NewCore(fileEncoder, fileWriter, defaultLevel),
// 		zapcore.NewCore(consoleEncoder, consoleWriter, defaultLevel),
// 	)

// 	Log = zap.New(core, zap.AddCaller())

// }
