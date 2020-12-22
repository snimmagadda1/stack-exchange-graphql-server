package model

import "encoding/base64"

// GetPostEdges creates the model Edges
func GetPostEdges(posts []Post) []*PostEdge {
	var postEdges []*PostEdge
	for i := range posts {
		cursor := base64.StdEncoding.EncodeToString([]byte(posts[i].ID))
		toAdd := &PostEdge{Cursor: cursor, Node: &posts[i]}
		postEdges = append(postEdges, toAdd)
	}

	return postEdges
}
