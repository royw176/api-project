A:  
你現在是一位世界前10大軟體公司的devops/SRE工程師，你具有豐富的帶領經驗，擅長教導新手，將他訓練成能夠獨當一面的高手。

今天我是一個新手，正在學習如何開發golang API。
我的主管給我一個任務: 使用golang開發API並與kong串接，並且可以在程式呼叫時，讓人輸入kong service name, kong service host url, kong service url path等等資訊。

我不需要程式框架，我需要你告訴我，我該怎麼分階段學習，每個階段要有什麼進度來完成這個任務?

B:  
## 階段 1：熟悉基本 API 與 Golang HTTP 模組
### 目標：
能用純 Golang 建立一個基本 API Server，處理 GET、POST 請求。

### 學習內容：
net/http 模組用法（handler, ServeMux）

http.Request 和 http.ResponseWriter 的使用

JSON 的解析與輸出（使用 encoding/json）

### 任務：
寫一個簡單的 API，例如：

GET /hello 回傳 { "message": "hello" }

POST /echo 接收 JSON，回傳一模一樣的資料

## 階段 2：理解 Kong API 管理與概念
### 目標：
理解 Kong 是什麼、它如何接收後端服務的註冊資訊。

### 學習內容：
Kong 的基本架構（Service, Route, Upstream, Target）

Kong Admin API 的用法（如 /services, /routes 等）

Kong 的 REST API 如何操作（用 curl 或 Postman 測試）

### 任務：
使用 curl 呼叫 Kong Admin API：

新增一個 Service

綁定一條 Route

使用 GET /services, GET /routes 查看結果

## 階段 3：使用 Golang 呼叫 Kong Admin API
### 目標：
學會如何用 Golang 建立、刪除、查詢 Kong 的服務。

### 學習內容：
使用 net/http 發送 POST/GET 請求

http.NewRequest 與設定 Header

JSON 結構體與 json.Marshal/Unmarshal

### 任務：
開發一個 CLI 工具，讓使用者輸入以下參數後：

kongServiceName

serviceHost（e.g., httpbin.org）

servicePath（e.g., /anything）

然後：

在 Kong 中建立對應的 Service

建立 Route

顯示呼叫結果

## 階段 4：撰寫 RESTful API 對外開放操作 Kong
### 目標：
從 Golang 自己寫的 API 來控制 Kong，而不是 CLI 工具。

### 學習內容：
Golang Handler 中如何接收 JSON body

怎麼把收到的參數轉成 HTTP 請求送給 Kong

如何回應錯誤訊息或成功訊息給使用者

### 任務：
建立一支 API：

POST /register-kong-service，Body 包含：

json
複製
編輯
{
  "service_name": "example",
  "host": "httpbin.org",
  "path": "/anything"
}
回傳：

成功訊息與 ID

或錯誤內容

## 階段 5：最佳化與擴充性設計
### 目標：
能寫出具有錯誤處理、重試邏輯、清楚日誌與可維護性的 API 程式碼。

### 學習內容：
封裝 HTTP 操作（用 function 分離邏輯）

撰寫結構化錯誤與 Log（可以考慮 log 或 zap 套件）

使用環境變數儲存 Kong Admin URL

### 任務：
將 Kong 操作抽成一個 package，例如 kongclient

所有邏輯寫成小函式，如 CreateService(), CreateRoute()

可以使用 log 列印每步操作狀態

支援 .env 檔或 os.Getenv 方式管理設定

### 延伸學習（之後再學也可以）：
使用 Swagger/OpenAPI 描述文件

撰寫單元測試（使用 httptest）

使用 Docker 包裝你的 Golang API

實際將 API 部署到測試環境並透過 Kong gateway 代理出去

## 前置作業
1. 安裝docker desktop
2. 使用docker compose 運行kong