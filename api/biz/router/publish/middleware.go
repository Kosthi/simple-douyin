// Code generated by hertz generator.

package publish

import (
	"github.com/cloudwego/hertz/pkg/app"
	mw "simple-douyin/api/biz/middleware"
)

func rootMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _douyinMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _publishMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _actionMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _publishactionMw() []app.HandlerFunc {
	// token未找到，但是客户端发送了...?
	return []app.HandlerFunc{mw.JwtMiddleware.MiddlewareFunc()}
	// return nil
}

func _listMw() []app.HandlerFunc {
	// your code...
	return nil
}

func _publishlistMw() []app.HandlerFunc {
	return []app.HandlerFunc{mw.JwtMiddleware.MiddlewareFunc()}
}
