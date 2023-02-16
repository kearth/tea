package env

var (
	ModeDebug = "debug"

	ModeTest = "test"

	ModeRelease = "release"
)

type Env struct {
	RootPath string
}

func init() {

}

// var DefaultEnv
/*
// AppName
func AppName() string {
	return appConf.AppName
}

// RunMode
func RunMode() string {
	return appConf.RunMode
}
*/

/*
func Root() string  { return "" }
func PID() string   { return "" }
func LogID() string { return "" }
*/
