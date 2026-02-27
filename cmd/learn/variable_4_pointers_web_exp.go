// 本文件演示：Web 后端里指针的常见用法，并串成一条完整请求流（PATCH 更新用户）
// 运行整条链路：在 main() 里调用 RunDemoServer()，用 curl 测试：
//   curl -X PATCH http://localhost:8080/user/1 -H "Content-Type: application/json" -d '{"nickname":"小明","age":18}'

package main

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

// ---------- 1. 共享配置：指针让全进程用同一份 ----------
type Config struct {
	Port int
	DB   string
}

func LoadConfig() *Config {
	return &(Config{Port: 8080, DB: "postgres://..."})
}

type Server struct {
	cfg *Config
	// 模拟 DB：map[id]*UserRow，用指针存一行，方便多处读同一行
	db map[int]*UserRow
}

func NewServer(cfg *Config) *Server {
	return &Server{cfg: cfg, db: mockDB()}
}

func mockDB() map[int]*UserRow {
	nick := "张三"
	age := 20
	return map[int]*UserRow{
		1: {ID: 1, Nickname: &nick, Age: &age},
		2: {ID: 2, Nickname: nil, Age: nil}, // 昵称、年龄都是 NULL
	}
}

// ---------- 2. 可选字段：PATCH 用 *string *int 区分“没传”和“传了零值” ----------
type PatchUserRequest struct {
	Nickname *string `json:"nickname"`
	Age      *int    `json:"age"`
}

// ---------- 3. 库表 NULL：用指针，nil 表示数据库 NULL ----------
type UserRow struct {
	ID       int     `json:"id"`
	Nickname *string `json:"nickname,omitempty"`
	Age      *int    `json:"age,omitempty"`
}

// ---------- 4. 业务模型：指针接收者才能改自己；返回 *User 方便继续改 ----------
type User struct {
	ID       int     `json:"id"`
	Name     string  `json:"name"`
	Nickname *string `json:"nickname,omitempty"`
	Age      *int    `json:"age,omitempty"`
	Valid    bool    `json:"-"`
}

func (u *User) Validate() bool {
	if u.Name == "" {
		u.Valid = false
		return false
	}
	u.Valid = true
	return true
}

// rowToUser 把 DB 行转成 API 用的 User（*User 让调用方可以继续改、统一用一份）
func rowToUser(row *UserRow, name string) *User {
	u := &User{ID: row.ID, Name: name, Nickname: row.Nickname, Age: row.Age}
	u.Validate()
	return u
}

// ---------- 整条链路：PATCH /user/:id ----------
func (s *Server) HandlePatchUser(w http.ResponseWriter, r *http.Request) {
	// 1) 读 ID（这里简化，实际可用 gorilla/mux 等）
	idStr := strings.TrimPrefix(r.URL.Path, "/user/")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "invalid id", http.StatusBadRequest)
		return
	}

	// 2) JSON 反序列化：必须传指针 &req，Decode 才能往 req 里写
	var req PatchUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "bad json", http.StatusBadRequest)
		return
	}

	// 3) 从“DB”取行（共享的是指针，拿到的就是那一行）
	row, ok := s.db[id]
	if !ok {
		http.Error(w, "not found", http.StatusNotFound)
		return
	}

	// 4) 行转 User（返回 *User，后面直接改同一份）
	u := rowToUser(row, "用户"+strconv.Itoa(id))

	// 5) 只更新“传了”的字段（nil=没传，非 nil=传了）
	if req.Nickname != nil {
		row.Nickname = req.Nickname
		u.Nickname = req.Nickname
	}
	if req.Age != nil {
		row.Age = req.Age
		u.Age = req.Age
	}

	// 6) 再次校验（指针接收者，改 u.Valid）
	u.Validate()

	// 7) 返回当前用户（同一份 *User 序列化出去）
	w.Header().Set("Content-Type", "application/json")
	_ = json.NewEncoder(w).Encode(u)
}

// RunDemoServer 启动演示服务，串起上述所有指针用法
func RunDemoServer() {
	cfg := LoadConfig()
	s := NewServer(cfg)
	http.HandleFunc("/user/", s.HandlePatchUser)
	_ = http.ListenAndServe(":"+strconv.Itoa(cfg.Port), nil)
}
