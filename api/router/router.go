package router

import (
    "net/http"

    // 自定義模組，路徑為$GOPATH/src
    controller "api/controller"

    // 第三方套件
    "github.com/gorilla/mux"
)

// 自定義Route內容型態。
type Route struct {
    Method string
    Pattern string
    Handler http.HandlerFunc
    Middleware mux.MiddlewareFunc
}

// [1]切片: 定義變數routes，它可以存放多筆Route資料。
var routes []Route

func init() {
    // register("HTTP method", "URL路徑", "呼叫controller裡的AddTodo函式", "可空")
    register("POST", "/api/todo", controller.AddTodo, nil)
    register("GET", "/api/todo/{id}", controller.GetTodoById, nil)
}

// [2] pointer
func NewRouter() *mux.Router {
    r := mux.NewRouter()
    for _, route := range routes {
        // [2]鍊式呼叫
        r.Methods(route.Method).
          Path(route.Pattern).
          Handler(route.Handler)

        if route.Middleware != nil {
            r.Use(route.Middleware)
        }
    }
    return r
}

func register(method, pattern string, handler http.HandlerFunc, middleware mux.MiddlewareFunc) {
    routes = append(routes, Route{method, pattern, handler, middleware})
}

// 一些說明
// [1]切片（slice）:是用來儲存「可變長度」資料的一種集合型別。

// [2]指標 (pointer): 來儲存某個變數記憶體位址的資料型別。
// example: 
// var x int = 10
// var p *int = &x (用 & 取 x 的記憶體位址)

// [3]鍊式呼叫 (method chaining): 函式回傳的是自己（或某個物件），讓你可以「點點點」一直串接呼叫。
//     在Go 中常見的風格，也稱為 fluent interface，用來讓 API 語法簡潔、易讀。