package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"sort"
)

type Users struct {
	Members []Member `json:"members"`
}

type Profile struct {
	RealName string `json:"real_name"`
	Title    string `json:"title,omitempty"`
}

type Member struct {
	Name    string  `json:"name"`
	Profile Profile `json:"profile"`
	Deleted bool    `json:"deleted,omitempty"`
	Updated int     `json:"updated"`
	IsBot   bool    `json:"is_bot,omitempty"`
}

func newRequest(endpoint string) *http.Request {
	req, _ := http.NewRequest(http.MethodGet, fmt.Sprintf("https://%s", endpoint), nil)
	req.Header.Add("Authorization", "Bearer {SLACK_APP_TOKEN}")
	return req
}

func main() {
	client := http.Client{}
	req := newRequest("slack.com/api/users.list?pretty=1")
	resp, err := client.Do(req)
	if err != nil {
		panic(err.Error())
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err.Error())
	}
	var users Users
	err = json.Unmarshal(body, &users)
	if err != nil {
		panic(err.Error())
	}
	var deactivated []Member
	for _, member := range users.Members {
		if member.Deleted == true && member.IsBot == false {
			deactivated = append(deactivated, member)
		}
	}
	sort.Slice(deactivated, func(i, j int) bool {
		return deactivated[i].Updated > deactivated[j].Updated
	})
	fmt.Println(len(deactivated))
	file, _ := json.MarshalIndent(deactivated, "", " ")

	_ = ioutil.WriteFile("deactivated_users.json", file, 0644)
}
