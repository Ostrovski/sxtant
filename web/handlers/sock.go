package handlers

import (
	"io/ioutil"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

func SockHandler() gin.HandlerFunc {
	upgrader := websocket.Upgrader{
		ReadBufferSize:  1024,
		WriteBufferSize: 1024,
	}

	url := "/questions?pagesize=5&order=desc&sort=creation&site=stackoverflow&access_token=QkPrVaNLSoxuefluUo3P7g))&key=0bSmpUxKaZoDijYvROtLlA(("
	// https://stackexchange.com/oauth?client_id=8429&redirect_uri=http://sxtant.micromind.me/auth

	return func(c *gin.Context) {
		conn, _ := upgrader.Upgrade(c.Writer, c.Request, nil)

		for {
			conn.WriteMessage(websocket.TextMessage, []byte(fetchQuestions(url)))
			time.Sleep(30 * time.Second)
		}
	}
}

func fetchQuestions(url string) string {
	res, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	questions, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		log.Fatal(err)
	}

	return string(questions)
}
