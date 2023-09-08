package terrors

var (
	// 0
	Succ = New(0, "success")

	// 100 default 默认
	ParamsEmpty = New(100000, "params empty") // 参数为空
	ParamsError = New(100001, "params error") // 参数错误
	FormatError = New(100002, "format error") // 参数格式错误
	AssertError = New(100003, "assert error") // 类型断言错误
	TransError  = New(100004, "trans error")  // 类型转换错误

	SystemError = New(100005, "system error") // 系统错误
	RPCError    = New(100006, "rpc error")    // 远程调用错误
	ConfError   = New(100007, "conf error")   // 内部配置错误
	NoData      = New(100008, "no data")      // 没有数据

	// 200 core 核心
	NameRegistered    = New(200000, "the name has registered") // 容器已注册
	NameNotRegistered = New(200001, "the name not regisered")  // 容器未注册

	// 300 user 用户自定义

	// 400

)
