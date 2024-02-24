package port

type IMonitorFunctionStart func()

type IMonitor interface {
	StartOnce()
}
