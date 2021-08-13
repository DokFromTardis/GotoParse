package main
import (
	"strconv"
	"os"
	"encoding/json"
	"strings"
	"fmt"
	"net/http"
	"io"
	"time"
)
type gotushnick struct{
	Response []struct {
		FirstName    string `json:"first_name"`
		ID           int    `json:"id"`
		LastName     string `json:"last_name"`
		Photo200Orig string `json:"photo_200_orig"`
		Deactivated  string `json:"deactivated,omitempty"`
	} `json:"response"`
}

type member struct {
	Response []struct {
		Usr gotushnick
		Friends []struct {
			FirstName    string `json:"first_name"`
			ID           int    `json:"id"`
			LastName     string `json:"last_name"`
			Photo200Orig string `json:"photo_200_orig"`
			TrackCode    string `json:"track_code"`
			Deactivated  string `json:"deactivated,omitempty"`
		} `json:"friends"`
	} `json:"response"`
}
type response struct{
	Response []int `json:"response"`
}
func main(){
	tok, _ = os.Open("token.txt")
	defer file.Close()
	data := make([]byte,64)
	for{
		n,err: = tok.Read(data)
		if err == io.EOF{
			break
		}
	}
	token:=string(data[:n])
	req := fmt.Sprintf("https://api.vk.com/method/execute.getMems?group_id=80270762&access_token=%s&v=5.52", token)
	resp, _ := http.Get(req)
	b, _ := io.ReadAll(resp.Body)
	var response response
	_ =  json.Unmarshal(b, &response)
	var members []member
	var response_strings []string
	for i := 0; i<len(response.Response); i++{
		response_strings = append(response_strings,strconv.Itoa(response.Response[i]))
	}
	for i := 0; i<len(response.Response)-23; i+=23{
		t := strings.Join(response_strings[i:i+23], ",")
		s:= fmt.Sprintf("https://api.vk.com/method/execute.getFriends?members=%s&access_token=%s&v=5.52",t,token)
		resp, _ = http.Get(s)
		time.Sleep(1*time.Second)
		res, _:= io.ReadAll(resp.Body)
		var mem member
		err := json.Unmarshal(res, &mem)
		for i:=0;i<len(mem.Response); i++{
			s=fmt.Sprintf("https://api.vk.com/method/users.get?user_ids=%d&fields=photo_200_orig&access_token=%s&v=5.52",mem.Response[i].User, token)
			resp, _ = http.Get(s)
			time.Sleep(time.Second)
			res,_= io.ReadAll(resp.Body)
			var mmbr gotushnick
			err = json.Unmarshal(res,&mmbr)
			mem.Response[i].Usr = mmbr
		}
		if err != nil{fmt.Println(err.Error())}
		members = append(members, mem)
		fmt.Println(mem.Response[0].Usr)
		fmt.Println(mem.Response[1].Usr)
	}
	mems, err := json.Marshal(&members)
	if err  != nil {
		fmt.Println(err.Error())
	}
	file, _ := os.Create("members.json")
	file.Write(mems)
}
