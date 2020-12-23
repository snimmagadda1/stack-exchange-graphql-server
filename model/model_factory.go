package model

import "encoding/base64"

// PostEdge creates an Edge for the post
func (p *Post) PostEdge() *PostEdge {
	cursor := base64.StdEncoding.EncodeToString([]byte(p.ID))
	return &PostEdge{Cursor: cursor, Node: p}
}

// UserEdge creates an Edge for the user
func (u *User) UserEdge() *UserEdge {
	cursor := base64.StdEncoding.EncodeToString([]byte(u.ID))
	return &UserEdge{Cursor: cursor, Node: u}
}

// CommentEdge creates an Edge for the comment
func (c *Comment) CommentEdge() *CommentEdge {
	cursor := base64.StdEncoding.EncodeToString([]byte(c.ID))
	return &CommentEdge{Cursor: cursor, Node: c}
}
