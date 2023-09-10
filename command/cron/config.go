package cron

type Config struct {
	Exp        string // 表达式
	Zone       string // 时区
	Start      string // 开始时间
	WithSecond bool   // 是否包含秒数
	Num        int    // 生成的数量
}

type Option func(*Config)
