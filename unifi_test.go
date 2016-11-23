package unifi

import (
	. "github.com/smartystreets/goconvey/convey"
	"testing"
	"net/http"
	"crypto/tls"
	"time"
	"os"
)

var unifiAddr="http://localhost:8443"
var unifiUser="admin"
var unifiPass="admin"
var unifiSite="default"
var unifiAuthorizeMac="11:22:33:44:55:66"

func init() {
	if len(os.Getenv("UNIFI_ADDR")) > 0 {
		unifiAddr = os.Getenv("UNIFI_ADDR")
	}
	if len(os.Getenv("UNIFI_USER")) > 0 {
		unifiUser = os.Getenv("UNIFI_USER")
	}
	if len(os.Getenv("UNIFI_PASS")) > 0 {
		unifiPass = os.Getenv("UNIFI_PASS")
	}
	if len(os.Getenv("UNIFI_SITE")) > 0 {
		unifiSite = os.Getenv("UNIFI_SITE")
	}
	if len(os.Getenv("UNIFI_TESTMAC")) > 0 {
		unifiAuthorizeMac = os.Getenv("UNIFI_TESTMAC")
	}

}


func TestUnlock(t *testing.T) {
	ctrl, err := New(unifiAddr)
	// test env, likely with bad/default cert
	ctrl.SetTransport(&http.Transport{
        TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
    })

	Convey("Init ok, testing on host " + unifiAddr, t, func() {
		So(err, ShouldEqual, nil)
	})
	err = ctrl.Login(unifiUser, unifiPass + "badpass")
	Convey("Logging in with bad password", t, func() {
		So(err, ShouldNotEqual, nil)
	})
	err = ctrl.Login(unifiUser, unifiPass)
	Convey("Logging in", t, func() {
	 	So(err, ShouldEqual, nil)
	})
	err = ctrl.AuthorizeGuest(unifiAuthorizeMac,time.Duration(time.Minute * 10))
	Convey("authorizing mac " + unifiAuthorizeMac, t, func() {
	 	So(err, ShouldEqual, nil)
	})
	clients, err := ctrl.GetClients()
	found := false
	for _, client := range clients {
		if client.Mac == unifiAuthorizeMac {
			found = true
		}
	}
	Convey("mac got authorized", t, func() {
	 	So(err, ShouldEqual, nil)
		So(found, ShouldEqual, true)
	})

}
