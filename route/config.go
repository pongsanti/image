package route

type Config struct {
	MaxMemory int64
	FormKey   string
	FileDir   string
}

func DefaultConfig() Config {
	return Config{
		MaxMemory: 10 << 20, // 10 MB
		FormKey:   "file",
		FileDir:   "images",
	}
}
