package shrutigo

import (
	"math/rand"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

var (
	sClient Client
	letters = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
)

func randSeq(n int) string {
	rand.Seed(time.Now().UnixNano())

	b := make([]rune, n)
	for i := range b {
		b[i] = letters[rand.Intn(len(letters))]
	}
	return string(b)
}

func init() {
	sClient = Client{"http://", "localhost:9574"}
}

func TestProvider(t *testing.T) {

	pname := randSeq(8)

	/*	Test Insert */
	p := Provider{Name: pname,
		DisplayName: "Test provider",
		Description: "test is in progress",
		WebURL:      "http://shrutiapp.com",
		IconURL:     "#",
		Voice:       "Brian",
	}

	err := sClient.RegisterProvider(p)
	assert.Nil(t, err)

	testProvider, err := sClient.GetSingleProvider(pname)
	assert.Nil(t, err)
	if assert.NotNil(t, testProvider) {

		assert.Equal(t, p.Name, testProvider.Name)
		assert.Equal(t, p.DisplayName, testProvider.DisplayName)
		assert.Equal(t, p.Description, testProvider.Description)
		assert.Equal(t, p.WebURL, testProvider.WebURL)
		assert.Equal(t, p.Voice, testProvider.Voice)
		assert.Equal(t, true, testProvider.Active)

	}

	/* Test Update */
	p.DisplayName = "Test2"
	p.Description = "testing update"
	p.WebURL = "#"
	p.IconURL = "http://foo.com"
	p.Voice = "Nicole"
	p.Active = true
	err = sClient.UpdateProvider(p)
	assert.Nil(t, err)

	testProvider, err = sClient.GetSingleProvider(pname)
	assert.Nil(t, err)
	if assert.NotNil(t, testProvider) {

		assert.Equal(t, p.Name, testProvider.Name)
		assert.Equal(t, p.DisplayName, testProvider.DisplayName)
		assert.Equal(t, p.Description, testProvider.Description)
		assert.Equal(t, p.WebURL, testProvider.WebURL)
		assert.Equal(t, p.Voice, testProvider.Voice)
		assert.Equal(t, true, testProvider.Active)

	}

	/* Test Delete */

	err = sClient.DeleteProvider(pname)
	assert.Nil(t, err)

	/* Test GetAllProviders */

	pname = randSeq(8)

	p = Provider{Name: pname,
		DisplayName: "Test provider",
		Description: "test is in progress",
		WebURL:      "http://shrutiapp.com",
		IconURL:     "#",
		Voice:       "Brian",
	}

	err = sClient.RegisterProvider(p)
	assert.Nil(t, err)

	pname2 := randSeq(8)
	p2 := Provider{Name: pname2,
		DisplayName: "Test2",
		Description: "testing update",
		WebURL:      "#",
		IconURL:     "http://foo.com",
		Voice:       "Nicole",
	}

	err = sClient.RegisterProvider(p2)
	assert.Nil(t, err)

	allProviders, err := sClient.GetAllProviders()
	assert.Nil(t, err)

	if assert.NotNil(t, allProviders) {
		assert.Equal(t, 2, len(allProviders))
	}

	err = sClient.DeleteProvider(pname)
	assert.Nil(t, err)

	err = sClient.DeleteProvider(pname2)
	assert.Nil(t, err)

	return

}

/* Test Notifications */

func TestNotification(t *testing.T) {

	pname := randSeq(8)

	/*	Create sample provider for notifications */
	p := Provider{Name: pname,
		DisplayName: "Test provider",
		Description: "test is in progress",
		WebURL:      "http://shrutiapp.com",
		IconURL:     "#",
		Voice:       "Brian",
	}

	err := sClient.RegisterProvider(p)
	assert.Nil(t, err)

	n := Notification{Title: "Test Notification Title",
		Url:          "http://foo.com/n",
		Key:          randSeq(5),
		ProviderName: pname,
		Priority:     PRIO_MED,
		Action:       ACT_POLL,
	}

	err = sClient.PushNotification(n)
	assert.Nil(t, err)

	n = Notification{Title: "Test Notification Title2",
		Url:          "http://foo.com/n2",
		Key:          randSeq(5),
		ProviderName: pname,
		Priority:     PRIO_MED,
		Action:       ACT_POLL,
	}

	err = sClient.PushNotification(n)
	assert.Nil(t, err)

	nt, err := sClient.GetUnheardNotifications()
	assert.Nil(t, err)
	if assert.NotNil(t, nt) {
		assert.NotEqual(t, 0, len(nt))
	}

	/* test error on duplicate key */

	err = sClient.PushNotification(n)
	assert.NotNil(t, err)

	err = sClient.DeleteProvider(pname)
	assert.Nil(t, err)
}
