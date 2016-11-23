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


func (c *Controller) Login(user string, pass string) (err error) {
	login := unifiLogin{
		User: user,
		Pass: pass,
	}
	json, _ := json.Marshal(&login)
	rsp, err := c.httpClient.Post(c.url + `/api/login`,"application/json",bytes.NewBuffer(json))
	if rsp.StatusCode == http.StatusOK {
	 	return err
	}
	body, err := ioutil.ReadAll(rsp.Body)
	return fmt.Errorf("req: %+v, body %s", rsp, body)
}

func (c *Controller) AuthorizeGuest(mac string, time time.Duration) (err error) {
	authorize := unifiAuthorize {
		Cmd: "authorize-guest",
		Mac: mac,
		Minutes: fmt.Sprintf("%1f",time.Minutes()),
	}
	json, _ := json.Marshal(&authorize)
	rsp, err := c.httpClient.Post(c.url + `/api/s/` + c.site + `/cmd/stamgr`,"application/json",bytes.NewBuffer(json))
	if rsp.StatusCode == http.StatusOK {
		return err
	}
	body, err := ioutil.ReadAll(rsp.Body)
	return fmt.Errorf("req: %+v, body %s", rsp, body)
}

func (c *Controller) GetClients() ([]UnifiClient, error){
	rsp, err := c.httpClient.Post(
		c.url + `/api/s/` + c.site + `/stat/guest?within=24r`,
		"application/json",
		bytes.NewBuffer([]byte("{}")),
	)
	_ = err
	body, _ := ioutil.ReadAll(rsp.Body)
	var result  UnifiClientResult
	json.Unmarshal(body, &result)
	return result.Data, err
}
