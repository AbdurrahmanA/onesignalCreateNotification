package onesignal

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

//IPushNotification IPushNotification
type IPushNotification interface {
	pushNotification(appID, msg, title string, id [][]string) error
}

//NotificationForUser NotificationForUser
type NotificationForUser struct {
	AppID            string   `json:"app_id"`
	IncludePlayerIds []string `json:"include_player_ids"`
	Data             struct {
		Foo string `json:"foo"`
	} `json:"data"`
	Contents struct {
		En string `json:"en"`
	} `json:"contents"`
	Heading struct {
		En string `json:"en"`
	} `json:"headings"`
}

func (a NotificationForUser) pushNotification(appID, msg, title string, id [][]string) error {
	var s []string
	for i := 0; i < len(id[0]); i++ {
		s = append(s, id[0][i])
	}
	a.AppID = appID
	a.Data.Foo = "bar"
	a.IncludePlayerIds = s
	a.Contents.En = msg
	a.Heading.En = title

	out, err := json.Marshal(a)
	if err != nil {
		//return false
	}
	reqBody := strings.NewReader(string(out))
	request, err := http.NewRequest("POST", "https://onesignal.com/api/v1/notifications", reqBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Basic NDBjM2I0YTMtNDNkNS00NTgwLWE2MWYtOGNkY2MxNzUyYTdk")
	if err != nil {
		//return false
	}
	resp, err := client.Client.Do(request)
	if err != nil {
		//return false
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println(result)
	if resp.StatusCode != 200 {
		//return false
	}
	//return true
	return nil
}

//NotificationForAllUsers NotificationForAllUsers
type NotificationForAllUsers struct {
	AppID            string   `json:"app_id"`
	IncludedSegments []string `json:"included_segments"`
	Data             struct {
		Foo string `json:"foo"`
	} `json:"data"`
	Contents struct {
		En string `json:"en"`
	} `json:"contents"`
	Heading struct {
		En string `json:"en"`
	} `json:"headings"`
}

func (w NotificationForAllUsers) pushNotiForAllUser(appID, msg, title string) error {
	var s []string
	s = append(s, "All")

	w.AppID = appID
	w.Data.Foo = "bar"
	w.IncludedSegments = s
	w.Contents.En = msg
	w.Heading.En = title

	out, err := json.Marshal(w)
	if err != nil {
		return errors.New("Json Marshal error")
	}

	reqBody := strings.NewReader(string(out))
	request, err := http.NewRequest("POST", "https://onesignal.com/api/v1/notifications", reqBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Basic NDBjM2I0YTMtNDNkNS00NTgwLWE2MWYtOGNkY2MxNzUyYTdk")
	if err != nil {
		return errors.New("http request error")
	}
	resp, err := client.Client.Do(request)
	if err != nil {
		return errors.New("Client Do error")
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	fmt.Println(result)
	if resp.StatusCode != 200 {
		return errors.New(strconv.Itoa(resp.StatusCode))
	}
	return nil
}

//NotificationForAllUsersAdapter NotificationForAllUsersAdapter
type NotificationForAllUsersAdapter struct {
	users NotificationForAllUsers
}

func (w NotificationForAllUsersAdapter) pushNotification(appID, msg, title string, id [][]string) error {
	err := w.users.pushNotiForAllUser(appID, msg, title)
	return err
}

//Push Push
func Push(i IPushNotification, n Notification, a AppCreate) error {
	err := i.pushNotification(a.AppID, n.getMessage(), n.getTitle(), n.getID())
	return err
}
