package main

import (
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestIndexHandler(t *testing.T) {
	assert := assert.New(t)

	res := httptest.NewRecorder()
	req := httptest.NewRequest("GET", "/", nil)

	mux := NewHttpHandler()
	mux.ServeHTTP(res, req)

	assert.Equal(http.StatusOK, res.Code)
	assert.Equal("Hello World", res.Body.String())
}

func TestBarHandler(t *testing.T) {
	t.Run("WithoutName", func(t *testing.T) {
		assert := assert.New(t)

		res := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/bar", nil)

		mux := NewHttpHandler()
		mux.ServeHTTP(res, req)

		assert.Equal(http.StatusOK, res.Code)
		assert.Equal("Hello World!", res.Body.String())
	})

	t.Run("WithName", func(t *testing.T) {
		assert := assert.New(t)

		res := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/bar?name=Neo", nil)

		mux := NewHttpHandler()
		mux.ServeHTTP(res, req)

		assert.Equal(http.StatusOK, res.Code)
		assert.Equal("Hello Neo!", res.Body.String())
	})
}

func TestFooHandler(t *testing.T) {
	t.Run("WithoutJson", func(t *testing.T) {
		assert := assert.New(t)

		res := httptest.NewRecorder()
		req := httptest.NewRequest("GET", "/foo", nil)

		mux := NewHttpHandler()
		mux.ServeHTTP(res, req)

		assert.Equal(http.StatusBadRequest, res.Code)
	})

	t.Run("WithJson", func(t *testing.T) {
		assert := assert.New(t)
		require := require.New(t)

		res := httptest.NewRecorder()
		req := httptest.NewRequest("POST", "/foo",
			strings.NewReader(`{"first_name":"Neo", "last_name":"Lee", "email":"Neo@Neonet.com"}`))

		mux := NewHttpHandler()
		mux.ServeHTTP(res, req)

		assert.Equal(http.StatusCreated, res.Code)

		user := new(User)
		err := json.NewDecoder(res.Body).Decode(user)
		require.Nil(err)
		assert.Equal("Neo", user.FirstName)
		assert.Equal("Lee", user.LastName)
	})
}
