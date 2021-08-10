package main
import (
//	"encoding/json"
	"strings"
	"fmt"
	"net/http"
	"io"
	"time"
)

type human struct{
	first_name string
	last_name string
	photo_200 string
}
type  items  struct{
	Count int
	Items human
}
func main(){
	token:="ed79dfb1e3c7726771d07c582e2ce306cac14312d7d293d5e58cb158f94d6029a0405745a083afa30128a"
	req := fmt.Sprintf("https://api.vk.com/method/execute.getMems?group_id=80270762&access_token=%s&v=5.52", token)
	resp, _ := http.Get(req)
	b, _ := io.ReadAll(resp.Body)
	var result []string
	i := 0
	for i := 0; i<len(b); i+=24{}
		t := string(b[i:i+2])
		t = strings.Replace(t, "[", "", -1)
		t = strings.Replace(t, "]", "", -1)
		s:= fmt.Sprintf("https://api.vk.com/method/execute.getFriends?members=%s&access_token=%s&v=5.52", t,token)
		resp, _ = http.Get(s)
		time.Sleep(1*time.Second)
		res, _:= io.ReadAll(resp.Body)
		result = append(result, string(res))
		fmt.Println(string(res))
//	}
}

