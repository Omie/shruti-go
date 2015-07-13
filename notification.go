package shrutigo

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path"
	"time"
)

const (
	// Priorities
	PRIO_LOW  = 10
	PRIO_MED  = 20
	PRIO_HIGH = 30

	// Actions
	ACT_POLL = 10
	ACT_PUSH = 20
)

type Notification struct {
	Id        int       `json:"id" db:"id"`
	Title     string    `json:"title" db:"title"`
	Url       string    `json:"url, omitempty" db:"url"`
	Key       string    `json:"key" db:"key"`
	Provider  string    `json:"provider" db:"provider"`
	CreatedOn time.Time `json:"created_on, omitempty" db:"created_on"`
	Priority  int       `json:"priority" db:"priority"`
	Action    int       `json:"action" db:"action"`
}

func (client *Client) GetNotificationsSince(since *time.Time) (n []*Notification, err error) {

	url := client.Protocol + path.Join(client.Host, "notifications", fmt.Sprintf("%d", since.Unix()))

	response, err := http.Get(url)
	if err != nil {
		return
	}
	defer response.Body.Close()

	contents, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return
	}

	n = make([]*Notification, 0)
	err = json.Unmarshal(contents, &n)
	if err != nil {
		return
	}

	return
}
