package txixwebservice

import "github.com/txix-open/isp-kit/http"

func (ws *Server) AddMiddleWare(mw []http.Middleware) {
	ws.wrapper.Middlewares = append(ws.wrapper.Middlewares, mw...)
}
