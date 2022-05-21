package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Post struct {
	Id     int    `json:"id"`
	UserId int    `json:"userId"`
	Title  string `json:"title"`
	Body   string `json:"body"`
}

func main() {
	post, err := getPosts()
	if err != nil {
		log.Fatal(err)
	}

	for _, p := range post {
		if p.Id == 8 {
			post, err := savePost(p.UserId, "test title", "test content")
			if err != nil {
				log.Fatal(err)
			}
			log.Printf("Save new post: %v", post)
		}
	}
}

func getPost(id int) (*Post, error) {
	url := fmt.Sprintf("https://jsonplaceholder.typicode.com/posts/%d", id)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	post := &Post{}

	err = json.NewDecoder(resp.Body).Decode(post)
	if err != nil {
		return nil, err
	}

	return post, nil
}

func getPosts() ([]*Post, error) {
	posts := []*Post{}
	url := "https://jsonplaceholder.typicode.com/posts/"

	resp, err := http.Get(url)
	if err != nil {
		return posts, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(&posts)
	if err != nil {
		return posts, err
	}

	return posts, nil
}

func savePost(userId int, title, body string) (*Post, error) {
	url := "https://jsonplaceholder.typicode.com/posts"

	post := &Post{
		UserId: userId,
		Title:  title,
		Body:   body,
	}

	content, err := json.Marshal(post)
	if err != nil {
		return nil, err
	}

	buf := bytes.NewBuffer(content)

	resp, err := http.Post(url, "application/json", buf)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	err = json.NewDecoder(resp.Body).Decode(post)

	return post, err
}
