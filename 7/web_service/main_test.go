package main_test

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"

	. "github.com/Joju-Matsumoto/go_web_app_practice/7/web_service"
)

var _ = Describe("Testing with Ginkgo", func() {
	var mux *http.ServeMux
	var post *FakePost
	var writer *httptest.ResponseRecorder

	BeforeEach(func() {
		post = &FakePost{}
		mux = http.NewServeMux()
		mux.Handle("/posts/", HandleRequest(post))
		writer = httptest.NewRecorder()
	})

	Context("Get a post using an id", func() {
		It("should get a post", func() {
			request, _ := http.NewRequest("GET", "/posts/1", nil)
			mux.ServeHTTP(writer, request)

			Expect(writer.Code).To(Equal(200))

			var post Post
			json.Unmarshal(writer.Body.Bytes(), &post)

			Expect(post.ID).To(Equal(uint(1)))
		})
	})

	Context("Get  an error if post id is not an integer", func() {
		It("should get a HTTP 500 response", func() {
			request, _ := http.NewRequest("GET", "/posts/hello", nil)
			mux.ServeHTTP(writer, request)

			Expect(writer.Code).To(Equal(500))
		})
	})
})
