package main

import (
	"testing"
	"time"
)

func TestDecode(t *testing.T) {
	post, err := decode("post.json")
	if err != nil {
		t.Error(err)
	}
	if post.Id != 1 {
		t.Error("Wrong id, was Expecting 1 but got", post.Id)
	}
	if post.Content != "Hello World!" {
		t.Error("Wrong content, was Expecting  but got", post.Content)
	}
}

func TestEncode(t *testing.T) {
	// 時間のかかるテストなどをスキップさせるのに使える
	t.Skip("Skipping encoding for now")
}

func TestLongRunning(t *testing.T) {
	// -shortが指定された場合
	if testing.Short() {
		t.Skip("Skipping long-running test in short mode")
	}
	time.Sleep(10 * time.Second)
}
