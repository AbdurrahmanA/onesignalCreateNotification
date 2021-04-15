package onesignal

import (
	b64 "encoding/base64"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"
)

//IPushNotification IPushNotification
type iPushNotification interface {
	pushNotification(appID, msg, title string, id [][]string) error
}

//NotificationForUser NotificationForUser
type notificationForUser struct {
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

//PushNotification PushNotification
func (a notificationForUser) pushNotification(appID, msg, title string, id [][]string) error {
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
		return errors.New("Json Marshal error")
	}
	reqBody := strings.NewReader(string(out))
	request, err := http.NewRequest("POST", "https://onesignal.com/api/v1/notifications", reqBody)
	request.Header.Set("Content-Type", "application/json")
	request.Header.Set("Authorization", "Basic "+b64.StdEncoding.EncodeToString([]byte(appID)))
	if err != nil {
		return errors.New("Http request error")
	}
	resp, err := client.Client.Do(request)
	if err != nil {
		return errors.New("Client Do error")
	}

	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	last := fmt.Sprintf("%v", result)
	if resp.StatusCode != 200 {
		return errors.New(last)
	}
	return nil
}

//NotificationForAllUsers NotificationForAllUsers
type notificationForAllUsers struct {
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

//PushNotiForAllUser PushNotiForAllUser
func (w notificationForAllUsers) pushNotiForAllUser(appID, msg, title string) error {
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
	request.Header.Set("Authorization", "Basic "+b64.StdEncoding.EncodeToString([]byte(appID)))
	if err != nil {
		return errors.New("http request error")
	}
	resp, err := client.Client.Do(request)
	if err != nil {
		return errors.New("Client Do error")
	}
	var result map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&result)
	last := fmt.Sprintf("%v", result)
	if resp.StatusCode != 200 {
		return errors.New(last)
	}
	return nil
}

//NotificationForAllUsersAdapter NotificationForAllUsersAdapter
type notificationForAllUsersAdapter struct {
	users notificationForAllUsers
}

//PushNotification PushNotification
func (w notificationForAllUsersAdapter) pushNotification(appID, msg, title string, id [][]string) error {
	err := w.users.pushNotiForAllUser(appID, msg, title)
	return err
}

//Push Push
func push(i iPushNotification, n notification, a AppCreate) error {
	err := i.pushNotification(a.AppID, n.getMessage(), n.getTitle(), n.getID())
	return err
}
