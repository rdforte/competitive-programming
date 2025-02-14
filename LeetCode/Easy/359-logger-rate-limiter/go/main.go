package main

func main() {
}

type Logger struct {
	logs map[string]int
}

func Constructor() Logger {
	return Logger{make(map[string]int)}
}

func (l *Logger) ShouldPrintMessage(timestamp int, message string) bool {
	t, ok := l.logs[message]
	if !ok {
		l.logs[message] = timestamp
		return true
	}

	if timestamp >= t+10 {
		l.logs[message] = timestamp
		return true
	}

	return false
}
