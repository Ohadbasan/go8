package resource

import (
	"reflect"
	"strconv"
	"time"

	"github.com/jinzhu/copier"
	"github.com/jinzhu/now"
	"github.com/volatiletech/null/v8"

	"github.com/gmhafiz/go8/internal/model"
)

type BookRequest struct {
	BookID        string `json:"-"`
	Title         string `json:"title" validate:"required"`
	PublishedDate string `json:"published_date" validate:"required"`
	ImageURL      string `json:"image_url" validate:"url"`
	Description   string `json:"description" validate:"required"`
}

type BookResource struct {
	BookID        int64       `json:"book_id" deepcopier:"field:book_id" db:"id"`
	Title         string      `json:"title" deepcopier:"field:title" db:"title"`
	PublishedDate time.Time   `json:"published_date" deepcopier:"field:force" db:"published_date"`
	ImageURL      null.String `json:"image_url" deepcopier:"field:image_url" db:"image_url"`
	Description   null.String `json:"description" deepcopier:"field:description"`
}

type BookDB struct {
	BookID        int64       `db:"book_id"`
	Title         string      `db:"title"`
	PublishedDate time.Time   `db:"published_date"`
	ImageURL      null.String `db:"image_url"`
	Description   null.String `db:"description"`
	CreatedAt     null.Time   `db:"created_at"`
	UpdatedAt     null.Time   `db:"updated_at"`
	DeletedAt     null.Time   `db:"deleted_at"`
}

func ToBook(req *BookRequest) *model.Book {
	id, err := strconv.ParseInt(req.BookID, 10, 64)
	if err != nil {
		return nil
	}
	return &model.Book{
		BookID:        id,
		Title:         req.Title,
		PublishedDate: now.MustParse(req.PublishedDate),
		ImageURL: null.String{
			String: req.ImageURL,
			Valid:  true,
		},
		Description: null.String{
			String: req.Description,
			Valid:  true,
		},
	}
}

func Book(book *model.Book) (BookResource, error) {
	var resource BookResource

	err := copier.Copy(&resource, &book)
	if err != nil {
		return resource, err
	}

	return resource, nil
}

func Books(books []*model.Book) (interface{}, error) {
	var resource BookResource

	if len(books) == 0 {
		return make([]string, 0), nil
	}

	rt := reflect.TypeOf(books)
	if rt.Kind() == reflect.Slice {
		var resources []BookResource
		for _, book := range books {
			res, _ := Book(book)
			resources = append(resources, res)
		}
		return resources, nil
	}

	err := copier.Copy(&resource, books)
	if err != nil {
		return resource, err
	}

	return resource, nil
}
