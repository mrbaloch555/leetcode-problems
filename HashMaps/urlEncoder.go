package main

import (
	"fmt"
	"math/rand"
	"time"
)

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

type Codec struct {
	urls map[string]string
}

func Constructor() Codec {
	rand.Seed(time.Now().UnixNano())
	return Codec{
		urls: map[string]string{},
	}
}

// Encodes a URL to a shortened URL.
func (this *Codec) encode(longUrl string) string {
	shortUrl := fmt.Sprintf("http://tinyurl.com/%s", randomString(5))
	this.urls[shortUrl] = longUrl
	return shortUrl
}

// Decodes a shortened URL to its original URL.
func (this *Codec) decode(shortUrl string) string {
	return this.urls[shortUrl]
}

func randomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		b[i] = charset[rand.Intn(len(charset))]
	}
	return string(b)
}

func main() {
	obj := Constructor()
	s := obj.encode("https://leetcode.com/problems/design-tinyurl")
	fmt.Println("SSS", obj.decode(s))
}
