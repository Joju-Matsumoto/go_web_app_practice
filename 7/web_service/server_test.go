package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	. "gopkg.in/check.v1"
)

type PostTestSuite struct {
	mux    *http.ServeMux
	post   *FakePost
	writer *httptest.ResponseRecorder
}

func init() {
	Suite(&PostTestSuite{})
}

func Test(t *testing.T) { TestingT(t) }

func (s *PostTestSuite) SetUpTest(c *C) {
	s.post = &FakePost{}
	s.mux = http.NewServeMux()
	s.mux.Handle("/posts/", handleRequest(s.post))
	s.writer = httptest.NewRecorder()
}

func (s *PostTestSuite) TearDownTest(c *C) {
	c.Log("Finished test - ", c.TestName())
}

func (s *PostTestSuite) SetUpSuite(c *C) {
	c.Log("Starting Post Test Suite")
}

func (s *PostTestSuite) TearDownSuite(c *C) {
	c.Log("Finishing Post Test Suite")
}

func (s *PostTestSuite) TestHandleGet(c *C) {
	request, _ := http.NewRequest("GET", "/posts/1", nil)
	s.mux.ServeHTTP(s.writer, request)
	c.Check(s.writer.Code, Equals, 200)
	var post Post
	json.Unmarshal(s.writer.Body.Bytes(), &post)
	c.Check(s.post.ID, Equals, uint(1))
}

func (s *PostTestSuite) TestHandlePut(c *C) {
	content := "Updated post"
	json := strings.NewReader(fmt.Sprintf(`{"content": "%s", "author": "joju"}`, content))
	request, _ := http.NewRequest("PUT", "/posts/1", json)
	s.mux.ServeHTTP(s.writer, request)

	c.Check(s.writer.Code, Equals, 200)
	c.Check(s.post.ID, Equals, uint(1))
	c.Check(s.post.Content, Equals, content)
}
