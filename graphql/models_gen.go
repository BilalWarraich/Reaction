// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package main


type ReactionTypesInput struct {
	Reacts string `json:"reacts"`
}


type ReactionsInput struct {
	ReactType string `json:"reactType"`
	PostID    string `json:"postId"`
	UserID    string `json:"userId"`
}
