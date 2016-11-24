package unifi

import (
	"net/http"
    "net/http/cookiejar"
	"encoding/json"
	"bytes"
	"fmt"
	"io/ioutil"
	"time"
)

type Controller struct {
	cookieJar *cookiejar.Jar
	httpClient  *http.Client
	site string
	url string
}

func New(url string, site ...string) (o *Controller, err error) {
	var ctrl Controller
	ctrl.cookieJar, _ = cookiejar.New(nil)
	ctrl.url=url
	ctrl.httpClient = &http.Client{
		Jar: ctrl.cookieJar,

	}
	if len(site) > 0 {
		ctrl.site=site[0]
	} else {
		ctrl.site="default"
	}
	return &ctrl, err
}


func (c *Controller) SetTransport(t *http.Transport) {
	c.httpClient.Transport = t
}

// Login, required before any other function will work, unifi uses cookies that are saved in memory
func (c *Controller) Login(user string, pass string) (err error) {
	login := unifiLogin{
		User: user,
		Pass: pass,
	}
	json, err := json.Marshal(&login)
	if err != nil {return  err}
	rsp, err := c.httpClient.Post(c.url + `/api/login`,"application/json",bytes.NewBuffer(json))
	if rsp.StatusCode == http.StatusOK {
	 	return err
	}
	body, err := ioutil.ReadAll(rsp.Body)
	return fmt.Errorf("req: %+v, body %s", rsp, body)
}
// authorize guest access for a given mac
func (c *Controller) AuthorizeGuest(mac string, time time.Duration) (err error) {
	authorize := unifiAuthorize {
		Cmd: "authorize-guest",
		Mac: mac,
		Minutes: int(time.Minutes()),
	}
	json, err := json.Marshal(&authorize)
	if err != nil { return  err}
	rsp, err := c.httpClient.Post(c.url + `/api/s/` + c.site + `/cmd/stamgr`,"application/json",bytes.NewBuffer(json))
	if err != nil { return err }
	if rsp.StatusCode == http.StatusOK { return err }
	body, err := ioutil.ReadAll(rsp.Body)
	return fmt.Errorf("req: %+v, body %s", rsp, body)
}
// get list of all clients within last 24 hours
func (c *Controller) GetClients() ([]UnifiClient, error){
	rsp, err := c.httpClient.Post(
		c.url + `/api/s/` + c.site + `/stat/guest?within=24r`,
		"application/json",
		bytes.NewBuffer([]byte("{}")),
	)
	if err != nil {return []UnifiClient{}, err}
	body, err := ioutil.ReadAll(rsp.Body)
	if err != nil {return []UnifiClient{}, err}
	var result  UnifiClientResult
	err = json.Unmarshal(body, &result)
	if err != nil {return []UnifiClient{}, err}
	return result.Data, err
}
