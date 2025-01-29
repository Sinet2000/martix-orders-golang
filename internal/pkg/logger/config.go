package logger

type LoggerConfig struct {
	Development LoggerEnvironmentConfig `json:"development"`
	Production  LoggerEnvironmentConfig `json:"production"`
}

type LoggerEnvironmentConfig struct {
	Level            string              `json:"level"`
	Encoding         string              `json:"encoding"`
	OutputPaths      []string            `json:"outputPaths"`
	ErrorOutputPaths []string            `json:"errorOutputPaths"`
	EncoderConfig    LoggerEncoderConfig `json:"encoderConfig"`
}

type LoggerEncoderConfig struct {
	MessageKey    string `json:"messageKey"`
	LevelKey      string `json:"levelKey"`
	LevelEncoder  string `json:"levelEncoder"`
	TimeKey       string `json:"timeKey"`
	TimeEncoder   string `json:"timeEncoder"`
	CallerKey     string `json:"callerKey"`
	CallerEncoder string `json:"callerEncoder"`
}
