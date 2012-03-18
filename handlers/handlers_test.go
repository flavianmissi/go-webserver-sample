package handlers_test

import (
	"webserver/handlers"
	. "launchpad.net/gocheck"
	"testing"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
)

func Test(t *testing.T) { TestingT(t)}

type WebServerSuite struct{}
var _ = Suite(&WebServerSuite{})

func (s *WebServerSuite) TestShouldRequestFoo(c *C) {
	request, err := http.NewRequest("GET", "/foo", nil)
	recorder := httptest.NewRecorder()

	c.Check(err, IsNil)
	handlers.FooHandler(recorder, request)

	data, err := ioutil.ReadAll(recorder.Body)
	c.Assert(string(data), Equals, "This is Foo handler!")
}

func (s *WebServerSuite) TestShouldRequestBar(c *C) {
	request, err := http.NewRequest("GET", "/bar", nil)
	recorder := httptest.NewRecorder()

	c.Check(err, IsNil)
	handlers.BarHandler(recorder, request)

	data, err := ioutil.ReadAll(recorder.Body)
	c.Assert(string(data), Equals, "This is Bar handler!!")
}
