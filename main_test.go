package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"net/url"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestAdd(t *testing.T) {
	type args struct {
		a string
		b string
	}
	tests := map[string]struct {
		args
		want string
	}{
		"test_1": {args: args{a: "1", b: "2"}, want: "12"},
		"test_2": {args: args{a: "11", b: "22"}, want: "1122"},
		"test_3": {args: args{a: "111", b: "222"}, want: "111222"},
		"test_4": {args: args{a: "1111", b: "2222"}, want: "11112222"},
		"test_5": {args: args{a: "11111", b: "22222"}, want: "1111122222"},
		"test_6": {args: args{a: "111111", b: "222222"}, want: "111111222222"},
		"test_7": {args: args{a: "1111111", b: "2222222"}, want: "11111112222222"},
	}

	for key, tt := range tests {
		t.Run(key, func(t *testing.T) {
			if got := Add(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("Add() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAddHandler(t *testing.T) {

	type args struct {
		a string
		b string
	}
	tests := map[string]struct {
		args
		want string
	}{
		"test_1": {args: args{a: "1", b: "2"}, want: "12"},
		"test_2": {args: args{a: "11", b: "22"}, want: "1122"},
		"test_3": {args: args{a: "111", b: "222"}, want: "111222"},
		"test_4": {args: args{a: "1111", b: "2222"}, want: "11112222"},
		"test_5": {args: args{a: "11111", b: "22222"}, want: "1111122222"},
		"test_6": {args: args{a: "111111", b: "222222"}, want: "111111222222"},
		"test_7": {args: args{a: "1111111", b: "2222222"}, want: "11111112222222"},
	}
	for key, tt := range tests {
		t.Run(key, func(t *testing.T) {
			request := http.Request{
				Method: "GET",
				URL: &url.URL{
					Path:     "/add",
					RawQuery: fmt.Sprintf("a=%s&b=%s", tt.args.a, tt.args.b),
				},
			}

			rw := httptest.NewRecorder()

			AddHandler(rw, &request)

			resp := rw.Body.String()
			expResp := tt.want
			assert.Equal(t, expResp, resp)
		})
	}
}
