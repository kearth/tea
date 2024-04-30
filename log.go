package tea

type ILog interface {
	Info(msg string)
}
type tlog struct {
}
