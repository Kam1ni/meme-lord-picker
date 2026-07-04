package memelord

import "time"

type Comment struct {
	Author string    `json:"author"`
	Text   string    `json:"text"`
	Date   time.Time `json:"date"`
}

type Meme struct {
	ID           int       `json:"id"`
	Title        string    `json:"title"`
	Tags         []string  `json:"tags"`
	UploadDate   time.Time `json:"upload_date"`
	Author       string    `json:"author"`
	MediaType    MediaType `json:"media_type"`
	ImageUrl     string    `json:"image_url"`
	ThubmnailUrl string    `json:"thumbnail_url"`
	Comments     []Comment `json:"comments"`
}

type MediaType string

const NO_MEDIA_TYPE MediaType = ""
const IMAGE MediaType = "image"
const VIDOE MediaType = "video"

type Ordering string

const NO_ORDERING Ordering = ""
const ORDERING_UPLOAD_DATE_NONE Ordering = ""
const ORDERING_UPLOAD_DATE_DESCENDING Ordering = "-upload_date"
const ORDERING_UPLOAD_DATE_ASCENDING Ordering = "upload_date"
const ORDERING_TITLE_ASCENDING Ordering = "title"
const ORDERING_TITLE_DESCENDING Ordering = "-title"

type Query struct {
	Username  string
	Title     string
	Tags      []string
	DateFrom  *time.Time
	DateTo    *time.Time
	MediaType MediaType
	Ordering  Ordering
	Page      int
	PageSize  int
}

type ErrorResponse struct {
	Error string `json:"error"`
}

type MemesResponse struct {
	Count    int    `json:"count"`
	Page     int    `json:"page"`
	PageSize int    `json:"page_size"`
	Next     *string `json:"next"`
	Previous *string `json:"previous"`
	Results  []Meme `json:"results"`
}
