package main

import (
	"fmt"
	"log"
	"time"
)

type Document struct {
	Key   string
	Title string
}

type Feed struct {
}

type CachingFeed struct {
	docs  map[string]Document
	fetch int
	*Feed
}

func NewCachingFeed(f *Feed) *CachingFeed {
	return &CachingFeed{
		docs:  make(map[string]Document),
		fetch: 0,
		Feed:  f,
	}
}

func (cf *CachingFeed) Count() int {
	return cf.fetch
}

func (cf *CachingFeed) Fetch(key string) (Document, error) {
	if doc, ok := cf.docs[key]; ok {
		return doc, nil
	}

	doc, err := cf.Feed.Fetch(key)
	if err != nil {
		return Document{}, err
	}

	cf.docs[key] = doc
	cf.fetch++
	return doc, nil
}

// ==================================================

// FetchCounter is the behavior we depend on for our process function.
type FetchCounter interface {
	Fetch(key string) (Document, error)
	Count() int
}

func process(fc FetchCounter) {
	keys := []string{"a", "a", "a", "b", "b", "b"}

	for _, key := range keys {
		doc, err := fc.Fetch(key)
		if err != nil {
			log.Printf("Could not fetch %s : %v", key, err)
			return
		}

		fmt.Printf("%s : %v\n", key, doc)
	}

	fmt.Printf("There are %d documents\n", fc.Count())
}

func (f *Feed) Fetch(key string) (Document, error) {
	time.Sleep(time.Second)

	doc := Document{
		Key:   key,
		Title: "Title for " + key,
	}
	return doc, nil
}

func (f *Feed) Count() int {
	return 0 // Feed не считает вызовы
}

func main() {
	fmt.Println("Using Feed directly")
	process(&Feed{})

	fmt.Println("Using CachingFeed")
	c := NewCachingFeed(&Feed{})
	process(c)
}
