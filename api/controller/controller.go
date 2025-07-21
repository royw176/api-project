package controller

import (
    "encoding/json"
    "fmt"
    "io"
//    "io/ioutil"
    "strconv"
    "github.com/gorilla/mux"
    "net/http"
    "api/services"
)

var TodoList []Todo

type Todo struct {
    Id int64 `json:"id"`
    Item string `json:"item"`
}

type ApiResponse struct {
    ResultCode string
    ResultMessage interface{}
}

// [1]go HTTP處理函式
func AddTodo(w http.ResponseWriter, r *http.Request) {
    //[2]多重附值，並設定io.LimitReader限制大小
    body, err := io.ReadAll(io.LimitReader(r.Body, 1024))

    if err != nil {
        fmt.Println(err)
    }

    var addTodo Todo
    //[3]匿名變數:_(底線)，這邊用於轉為Json，但忽略json解析錯誤。
    _ = json.Unmarshal(body, &addTodo)
    TodoList = append(TodoList, addTodo)
    //[4]defer
    defer r.Body.Close()
    response := ApiResponse{"200", TodoList}

    // 回傳
    services.ResponseWithJson(w, http.StatusOK, response)
}

func GetTodoById(w http.ResponseWriter, r *http.Request) {
    vars := mux.Vars(r)
    // 獲取url參數
    queryId := vars["id"]
    var targetTodo Todo
    for _, Todo := range TodoList {
        intQueryId, _ := strconv.ParseInt(queryId, 10, 64)
        if Todo.Id == intQueryId {
            targetTodo = Todo
        }
    }
    response := ApiResponse{"200", targetTodo}
    services.ResponseWithJson(w, http.StatusOK, response)
}

// 一些說明
// [1]go HTTP處理函式基本語法: func handler(w http.ResponseWriter, r *http.Request)
// w：輸出
// r：輸入（HTTP 請求）

// [2]多重附值: body, err := function()
// 意思是: 執行函式 ioutil.ReadAll(...)，它會回傳兩個值；
// 第一個值給 body，第二個值給 err。

//[3]匿名變數:_(底線): 
// 忽略你不需要使用的變數值，當一個函式回傳多個值，但你只想要其中幾個，就可以使用 _ 忽略不要的。

//[4]defer:
// 延遲執行某段程式碼，直到函式結束時(return)才執行。
// 常用於: 關檔案、關網路連線、Unlock() mutex、Recovery from panic、釋放資源等等。