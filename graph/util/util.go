package util

import (
	b64 "encoding/base64"
	"strconv"

	"github.com/snimmagadda1/graphql-api/graph/model"
)

// GetCursor gets the specified cursor
func GetCursor(post model.Post, field string) string {
	sEnc := ""
	switch field {
	case "Score":
		sEnc = b64.StdEncoding.EncodeToString([]byte(strconv.Itoa(*post.Score)))
	case "CreationDate":
		sEnc = b64.StdEncoding.EncodeToString([]byte(post.CreationDate.String()))
	case "LastActivityDate":
		sEnc = b64.StdEncoding.EncodeToString([]byte(post.CreationDate.String()))
	default:
		sEnc = b64.StdEncoding.EncodeToString([]byte(post.ID))
	}

	return sEnc
}
