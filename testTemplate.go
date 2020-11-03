package main

import (
	"database/sql"
	_ "github.com/lib/pq"
	"html/template"
	"log"
	"net/http"
	"time"
)

var Db *sql.DB

type Thread struct {
	Id        int
	Uuid      string
	Topic     string
	UserId    int
	CreatedAt time.Time
}

type User struct {
	Id        int
	Uuid      string
	Name      string
	Email     string
	Password  string
	CreatedAt time.Time
}

func init() {
	var err error
	Db, err = sql.Open("postgres", "user=postgres password=123456 dbname=chitchat")
	if err != nil {
		log.Fatal(err)
	}
	return
}

// format the CreatedAt date to display nicely on the screen
func (thread *Thread) CreatedAtDate() string {
	return thread.CreatedAt.Format("Jan 2, 2006 at 3:04pm")
}

// get the number of posts in a thread
func (thread *Thread) NumReplies() (count int) {
	rows, err := Db.Query("SELECT count(*) FROM posts where thread_id = $1", thread.Id)
	if err != nil {
		return
	}
	for rows.Next() {
		if err = rows.Scan(&count); err != nil {
			return
		}
	}
	rows.Close()
	return
}

// Get the user who started this thread
func (thread *Thread) User() (user User) {
	user = User{}
	_ = Db.QueryRow("SELECT id, uuid, name, email, created_at FROM users WHERE id = $1", thread.UserId).
		Scan(&user.Id, &user.Uuid, &user.Name, &user.Email, &user.CreatedAt)
	return
}

func index(writer http.ResponseWriter, _ *http.Request) {
	thread := Thread{
		Id:        1,
		Uuid:      "99999999999999",
		Topic:     "bear",
		UserId:    3,
		CreatedAt: time.Now(),
	}

	t := template.Must(template.ParseFiles("t1.html"))
	_ = t.Execute(writer, &thread)
}

func main() {
	http.HandleFunc("/test", index)
	_ = http.ListenAndServe("127.0.0.1:9000", nil)
}
