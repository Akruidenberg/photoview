// Code generated by github.com/99designs/gqlgen, DO NOT EDIT.

package models

import (
	"fmt"
	"io"
	"strconv"
	"time"
)

type AuthorizeResult struct {
	Success bool    `json:"success"`
	Status  string  `json:"status"`
	Token   *string `json:"token"`
}

type MediaDownload struct {
	Title    string    `json:"title"`
	MediaURL *MediaURL `json:"mediaUrl"`
}

type Notification struct {
	Key      string           `json:"key"`
	Type     NotificationType `json:"type"`
	Header   string           `json:"header"`
	Content  string           `json:"content"`
	Progress *float64         `json:"progress"`
	Positive bool             `json:"positive"`
	Negative bool             `json:"negative"`
	// Time in milliseconds before the notification will close
	Timeout *int `json:"timeout"`
}

type Ordering struct {
	OrderBy        *string         `json:"order_by"`
	OrderDirection *OrderDirection `json:"order_direction"`
}

type Pagination struct {
	Limit  *int `json:"limit"`
	Offset *int `json:"offset"`
}

type ScannerResult struct {
	Finished bool     `json:"finished"`
	Success  bool     `json:"success"`
	Progress *float64 `json:"progress"`
	Message  *string  `json:"message"`
}

type SearchResult struct {
	Query  string   `json:"query"`
	Albums []*Album `json:"albums"`
	Media  []*Media `json:"media"`
}

// Credentials used to identify and authenticate a share token
type ShareTokenCredentials struct {
	Token    string  `json:"token"`
	Password *string `json:"password"`
}

type TimelineGroup struct {
	Album      *Album    `json:"album"`
	Media      []*Media  `json:"media"`
	MediaTotal int       `json:"mediaTotal"`
	Date       time.Time `json:"date"`
}

type LanguageTranslation string

const (
	LanguageTranslationEnglish LanguageTranslation = "English"
	LanguageTranslationFrench  LanguageTranslation = "French"
	LanguageTranslationItalian LanguageTranslation = "Italian"
	LanguageTranslationSwedish LanguageTranslation = "Swedish"
	LanguageTranslationDanish  LanguageTranslation = "Danish"
	LanguageTranslationSpanish LanguageTranslation = "Spanish"
	LanguageTranslationPolish  LanguageTranslation = "Polish"
	LanguageTranslationGerman  LanguageTranslation = "German"
	LanguageTranslationRussian LanguageTranslation = "Russian"
)

var AllLanguageTranslation = []LanguageTranslation{
	LanguageTranslationEnglish,
	LanguageTranslationFrench,
	LanguageTranslationItalian,
	LanguageTranslationSwedish,
	LanguageTranslationDanish,
	LanguageTranslationSpanish,
	LanguageTranslationPolish,
	LanguageTranslationGerman,
	LanguageTranslationRussian,
}

func (e LanguageTranslation) IsValid() bool {
	switch e {
	case LanguageTranslationEnglish, LanguageTranslationFrench, LanguageTranslationItalian, LanguageTranslationSwedish, LanguageTranslationDanish, LanguageTranslationSpanish, LanguageTranslationPolish, LanguageTranslationGerman, LanguageTranslationRussian:
		return true
	}
	return false
}

func (e LanguageTranslation) String() string {
	return string(e)
}

func (e *LanguageTranslation) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = LanguageTranslation(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid LanguageTranslation", str)
	}
	return nil
}

func (e LanguageTranslation) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type NotificationType string

const (
	NotificationTypeMessage  NotificationType = "Message"
	NotificationTypeProgress NotificationType = "Progress"
	// Close a notification with a given key
	NotificationTypeClose NotificationType = "Close"
)

var AllNotificationType = []NotificationType{
	NotificationTypeMessage,
	NotificationTypeProgress,
	NotificationTypeClose,
}

func (e NotificationType) IsValid() bool {
	switch e {
	case NotificationTypeMessage, NotificationTypeProgress, NotificationTypeClose:
		return true
	}
	return false
}

func (e NotificationType) String() string {
	return string(e)
}

func (e *NotificationType) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = NotificationType(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid NotificationType", str)
	}
	return nil
}

func (e NotificationType) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}

type OrderDirection string

const (
	OrderDirectionAsc  OrderDirection = "ASC"
	OrderDirectionDesc OrderDirection = "DESC"
)

var AllOrderDirection = []OrderDirection{
	OrderDirectionAsc,
	OrderDirectionDesc,
}

func (e OrderDirection) IsValid() bool {
	switch e {
	case OrderDirectionAsc, OrderDirectionDesc:
		return true
	}
	return false
}

func (e OrderDirection) String() string {
	return string(e)
}

func (e *OrderDirection) UnmarshalGQL(v interface{}) error {
	str, ok := v.(string)
	if !ok {
		return fmt.Errorf("enums must be strings")
	}

	*e = OrderDirection(str)
	if !e.IsValid() {
		return fmt.Errorf("%s is not a valid OrderDirection", str)
	}
	return nil
}

func (e OrderDirection) MarshalGQL(w io.Writer) {
	fmt.Fprint(w, strconv.Quote(e.String()))
}
