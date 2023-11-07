package logger

import "go.uber.org/zap"

func New(name string) (*zap.Logger, error) {
	logConfig := zap.NewProductionConfig()
	logConfig.Level = zap.NewAtomicLevelAt(zap.InfoLevel)

	logger, err := logConfig.Build()
	if err != nil {
		return nil, err
	}

	return logger.Named(name), nil
}
