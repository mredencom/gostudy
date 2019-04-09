package counters

// counters 包提供告警计数器的功能
// 这个类型用于保存告警计数
type alertCounter int

func New(value int) alertCounter {
	return alertCounter(value)
}
