package model

import (
	"time"
)

// User models the graphql layer and entity
type User struct {
	ID               string     `json:"id" gorm:"primaryKey;type:int;column:Id"`
	Reputation       *string    `json:"reputation" gorm:"type:int;column:Reputation"`
	CreationDate     *time.Time `json:"creationDate" gorm:"type:DATETIME;column:CreationDate"`
	DisplayName      *string    `json:"displayName" gorm:"type:VARCHAR(50);column:DisplayName"`
	LastAccessDate   *time.Time `json:"lastAccessDate" gorm:"type:DATETIME;column:LastAccessDate"`
	Views            *int       `json:"views" gorm:"type:int;column:Views"`
	WebsiteURL       *string    `json:"websiteUrl" gorm:"type:VARCHAR(256);column:WebsiteUrl"`
	Location         *string    `json:"location" gorm:"type:VARCHAR(256);column:Location"`
	AboutMe          *string    `json:"aboutMe" gorm:"type:TEXT;column:AboutMe"`
	Age              *int       `json:"age" gorm:"type:int;column:Age"`
	UpVotes          *int       `json:"upVotes" gorm:"type:int;column:UpVotes"`
	DownVotes        *int       `json:"downVotes" gorm:"type:int;column:DownVotes"`
	EmailHash        *string    `json:"emailHash" gorm:"type:VARCHAR(256);column:EmailHash"`
	AccountID        *int       `json:"accountId" gorm:"type:int;column:AccountId"`
	LastAccessedDate *time.Time `json:"lastAccessedDate" gorm:"type:DATETIME;column:LastAccessedDate"`
	ProfileImageURL  *string    `json:"profileImageUrl" gorm:"type:VARCHAR(255);column:ProfileImageUrl"`
}

// Post models the graphql layer and entity
type Post struct {
	ID                    string     `json:"id" gorm:"primaryKey;type:int;column:Id"`
	PostType              *int       `json:"postType" gorm:"type:int;column:PostType"`
	AcceptedAnswerID      *int       `json:"acceptedAnswerId" gorm:"type:int;column:AcceptedAnswerId"`
	ParentID              *int       `json:"parentId" gorm:"type:int;column:ParentId"`
	Score                 *int       `json:"score" gorm:"type:int;column:Score"`
	ViewCount             *int       `json:"viewCount" gorm:"type:int;column:ViewCount"`
	Body                  *string    `json:"body" gorm:"type:MEDIUMTEXT;column:Body"`
	CreationDate          *time.Time `json:"creationDate" gorm:"type:DATETIME;column:CreationDate"`
	OwnerUserID           *int       `json:"ownerUserId" gorm:"type:int;column:OwnerUserId"`
	LastEditorDisplayName *string    `json:"lastEditorDisplayName" gorm:"type:VARCHAR(50);column:LastEditorDisplayName"`
	LastActivityDate      *time.Time `json:"lastActivityDate" gorm:"type:DATETIME;column:LastActivityDate"`
	ClosedDate            *time.Time `json:"closedDate" gorm:"type:DATETIME;column:ClosedDate"`
	Title                 *string    `json:"title" gorm:"type:VARCHAR(256);column:Title"`
	Tags                  *string    `json:"tags" gorm:"type:VARCHAR(256);column:Tags"`
	AnswerCount           *int       `json:"answerCount" gorm:"type:int;column:AnswerCount"`
	CommentCount          *int       `json:"commentCount" gorm:"type:int;column:CommentCount"`
	ContentLicense        *string    `json:"contentLicense" gorm:"type:VARCHAR(20);column:ContentLicense"`
}

// Badge models the graphql layer and entity
type Badge struct {
	ID         string     `json:"id" gorm:"primaryKey;type:int;column:Id"`
	UserID     *int       `json:"userId" gorm:"type:int;column:UserId"`
	Name       *string    `json:"name" gorm:"type:VARCHAR(50);column:Name"`
	Date       *time.Time `json:"date" gorm:"type:DATETIME;column:Date"`
	Class      *int       `json:"class" gorm:"type:int;column:Class"`
	BadgeClass *int       `json:"badgeClass" gorm:"type:int;column:BadgeClass"`
	TagBased   *int       `json:"TagBased" gorm:"type:int;column:TagBased"`
}

// Comment models the graphql layer and entity
type Comment struct {
	ID              string     `json:"id" gorm:"primaryKey;type:int;column:Id"`
	PostID          *int       `json:"postId" gorm:"type:int;column:PostId"`
	Score           *int       `json:"score" gorm:"type:int;column:Score"`
	Text            *string    `json:"text" gorm:"type:TEXT;column:Text"`
	CreationDate    *time.Time `json:"creationDate" gorm:"type:DATETIME;column:CreationDate"`
	UserDisplayName *string    `json:"userDisplayName" gorm:"type:VARCHAR(50);column:UserDisplayName"`
	UserID          *int       `json:"userId" gorm:"type:int;column:UserId"`
	ContentLicense  *string    `json:"contentLicense" gorm:"type:VARCHAR(20);column:ContentLicense"`
}

// PostHistory models the graphql layer and entity
type PostHistory struct {
	ID                string     `json:"id" gorm:"primaryKey;type:int;column:Id"`
	PostHistoryTypeID *int       `json:"postHistoryTypeId" gorm:"type:int;column:PostHistoryTypeId"`
	PostID            *int       `json:"postId" gorm:"type:int;column:PostID"`
	RevisionGUID      *string    `json:"revisionGUID" gorm:"type:VARCHAR(36);column:RevisionGUID"`
	CreationDate      *time.Time `json:"creationDate" gorm:"type:DATETIME;column:CreationDate"`
	UserID            *int       `json:"userId" gorm:"type:int;column:UserId"`
	UserDisplayName   *string    `json:"userDisplayName" gorm:"type:VARCHAR(36);column:UserDisplayName"`
	Comment           *string    `json:"comment" gorm:"type:TEXT;column:Comment"`
	CloseReasonTypes  *int       `json:"closeReasonTypes" gorm:"type:int;column:CloseReasonTypes"`
	PostNoticeID      *int       `json:"postNoticeId" gorm:"type:int;column:PostNoticeId"`
	Text              *string    `json:"text" gorm:"type:TEXT;column:Text"`
	ContentLicense    *string    `json:"contentLicense" gorm:"type:VARCHAR(20);column:ContentLicense"`
}

// TableName overrides the table name
func (PostHistory) TableName() string {
	return "post_history"
}

// Vote models the graphql layer and entity
type Vote struct {
	ID           string     `json:"id" gorm:"primaryKey;type:int;column:Id"`
	PostID       *int       `json:"postId" gorm:"type:int;column:PostId"`
	VoteTypeID   *int       `json:"voteTypeId" gorm:"type:int;column:VoteTypeId"`
	UserID       *int       `json:"userId" gorm:"type:int;column:UserId"`
	CreationDate *time.Time `json:"creationDate" gorm:"type:DATETIME;column:CreationDate"`
	BountyAmount *int       `json:"bountyAmount" gorm:"type:int;column:BountyAmount"`
}
