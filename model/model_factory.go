package model

import "encoding/base64"

// PostEdge creates an Edge for the post
func (p *Post) PostEdge() *PostEdge {
	cursor := base64.StdEncoding.EncodeToString([]byte(p.ID))
	return &PostEdge{Cursor: cursor, Node: p}
}
