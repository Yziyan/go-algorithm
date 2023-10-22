// @Author: Ciusyan 10/22/23

package other

// Options 模式比较适用于构建不是特别复杂的对象，
// 我们想要限制用户必须传入某些参数，但是有些参数是可选的，那么我们就可以使用 options 模式来限制用户

type Config struct {
	host    string
	port    int
	timeout int
	isLog   bool

	// 这个必须不能为 nil
	mustNotNil *int
}

// ConfigOption 这里可以使用函数式选项来做
type ConfigOption func(*Config)

// WithTimeoutOption 比如这里给定一个 timeout 的选项
func WithTimeoutOption(timeout int) ConfigOption {
	return func(config *Config) {
		config.timeout = timeout
	}
}

// WithIsLogOption 比如这里给定一个 isLog 的选项
func WithIsLogOption(isLog bool) ConfigOption {
	return func(config *Config) {
		config.isLog = isLog
	}
}

// NewConfig 比如这里用户想要使用 Config ，就必须要传入 host 和 port，但是 timeout 和 isLog 就是可选的。
func NewConfig(host string, port int, opts ...ConfigOption) *Config {
	cfg := &Config{
		host: host,
		port: port,

		// 下面可以给个默认值，也可以不给默认值
		timeout: 2,
		isLog:   false,
	}

	// 关键，遍历调用传入的 options
	for _, opt := range opts {
		opt(cfg)
	}

	return cfg
}
