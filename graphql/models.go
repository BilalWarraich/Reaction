package main

type ReactionTypes struct {
	ID     string  `json:"id"`
	Reacts   string  `json:"reacts"`
}

type Reactions struct {
	ID        string `json:"id"`
	ReactType string `json:"reactType"`
	PostID    string `json:"postId"`
	UserID    string `json:"userId"`
}