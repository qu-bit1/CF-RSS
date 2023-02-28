package models

// Package models contains all the shared models for the application.

// BlogEntry represents a sample blog on Codeforces.
type BlogEntry struct {
	Id                      int
	OriginalLocale          string   `json:"originalLocale"`
	CreationTimeSeconds     int64    `json:"creationTimeSeconds"`
	AuthorHandle            string   `json:"authorHandle"`
	Title                   string   `json:"title"`
	Content                 string   `json:"content"`
	Locale                  string   `json:"locale"`
	ModificationTimeSeconds int64    `bson:"modificationTimeSeconds" json:"modificationTimeSeconds"`
	AllowViewHistory        bool     `bson:"allowViewHistory" json:"allowViewHistory"`
	Tags                    []string `bson:"tags" json:"tags"`
	Rating                  int      `bson:"rating" json:"rating"`
}

// Comment represents a sample comment on a Codeforces blog.
type Comment struct {
	Id                  int    `bson:"id" json:"id"`
	CreationTimeSeconds int64  `bson:"creationTimeSeconds" json:"creationTimeSeconds"`
	CommentatorHandle   string `bson:"commentatorHandle" json:"commentatorHandle"`
	Locale              string `bson:"locale" json:"locale"`
	Text                string `bson:"text" json:"text"`
	ParentCommentId     int    `bson:"parentCommentId" json:"parentCommentId"`
	Rating              int    `bson:"rating" json:"rating"`
}

// RecentAction represents an activity on Codeforces blog/comment.
type RecentAction struct {
	TimeSeconds int64      `bson:"timeSeconds" json:"timeSeconds"`
	BlogEntry   *BlogEntry `bson:"blogEntry,omitempty" json:"blogEntry,omitempty"`
	Comment     *Comment   `bson:"comment,omitempty" json:"comment,omitempty"`
}
