package httpserver

import (
	"fmt"
	"mime"
	"net/http"

	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/errors/gcode"
	"github.com/gogf/gf/v2/errors/gerror"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/kearth/klib/klog"
	"github.com/kearth/tea/frame/base"
)

// Middleware HTTP 中间件
type Middleware = func(r *Request)

type logFormat struct {
	Status   int                 `json:"status"`
	Method   string              `json:"method"`
	Path     string              `json:"path"`
	Proto    string              `json:"proto"`
	Duration string              `json:"duration"`
	RemoteIP string              `json:"remote_ip"`
	Headers  map[string][]string `json:"headers"`
	Response any                 `json:"response"`
}

const (
	contentTypeEventStream  = "text/event-stream"
	contentTypeOctetStream  = "application/octet-stream"
	contentTypeMixedReplace = "multipart/x-mixed-replace"
)

// MiddlewareHandlerResponse HTTP 中间件 - 响应处理
func MiddlewareHandlerResponse(r *Request) {
	r.Middleware.Next()

	// There's custom buffer content, it then exits current handler.
	if r.Response.BufferLength() > 0 || r.Response.Writer.BytesWritten() > 0 {
		f := LogFormat(r, map[string]any{})
		klog.Info(r.Context(), f)
		return
	}

	// It does not output common response content if it is stream response.
	mediaType, _, _ := mime.ParseMediaType(r.Response.Header().Get("Content-Type"))

	for _, ct := range []string{contentTypeEventStream, contentTypeOctetStream, contentTypeMixedReplace} {
		if mediaType == ct {
			return
		}
	}

	var (
		msg  string
		err  = r.GetError()
		res  = r.GetHandlerResponse()
		code = gerror.Code(err)
	)
	if err != nil {
		if code == gcode.CodeNil {
			code = gcode.CodeInternalError
		}
		msg = err.Error()
	} else {
		if r.Response.Status > 0 && r.Response.Status != http.StatusOK {
			switch r.Response.Status {
			case http.StatusNotFound:
				code = gcode.CodeNotFound
			case http.StatusForbidden:
				code = gcode.CodeNotAuthorized
			default:
				code = gcode.CodeUnknown
			}
			// It creates an error as it can be retrieved by other middlewares.
			err = gerror.NewCode(code, msg)
			r.SetError(err)
		} else {
			code = gcode.CodeOK
		}
		msg = code.Message()
	}

	o := base.Response()
	// TODO 这里可以根据业务场景自定义响应格式
	o.SetCode(code.Code())
	o.SetMsg(msg)
	o.SetData(res)
	f := LogFormat(r, o)
	klog.Info(r.Context(), f)
	r.Response.WriteJson(o)
}

// MiddlewareHandlerCORS HTTP 中间件 - CORS 处理
func MiddlewareHandlerCORS(r *Request) {
	ghttp.MiddlewareCORS(r)
}

// LogFormat 日志格式
func LogFormat(r *Request, content any) string {
	ns := r.LeaveTime.Sub(r.EnterTime).Nanoseconds()
	ms := ns / 1000000
	lf := logFormat{
		Status:   r.Response.Status,
		Method:   r.Request.Method,
		Path:     r.URL.Path,
		Proto:    r.Proto,
		Duration: fmt.Sprintf("%d ms", ms),
		RemoteIP: r.GetRemoteIp(),
		Headers:  r.Header,
		Response: content,
	}

	str, _ := gjson.EncodeString(lf)
	return str
}
