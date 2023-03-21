package handler

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"log"
	"net/http"
)

// TestSimpleApi ecs1的上层应用
func TestSimpleApi(ip string, port int, path string) *gin.Engine {
	r := gin.Default()
	r.GET(path, FindEcs2(ip, port, path))
	return r
}

func FindEcs2(ip string, port int, path string) func(c *gin.Context) {

	return func(c *gin.Context) {
		url := fmt.Sprintf("http://%s:%d%s", ip, port, path)
		resp, err := http.Get(url)
		if err != nil {
			log.Println("Failed to Get resp from rpc test")
			return
		}
		var m map[string]string
		dataByte, _ := ioutil.ReadAll(resp.Body)
		err = json.Unmarshal(dataByte, &m)
		if err != nil {
			log.Println("Failed to unmarshal body to byte")
			return
		}
		log.Printf("%+v", m)

		c.Data(200, "json", dataByte)
	}
}
