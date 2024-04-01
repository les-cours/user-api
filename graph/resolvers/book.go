package resolvers

import (
	"context"
	"fmt"
	"github.com/les-cours/user-api/graph/models"
	pb "github.com/les-cours/user-api/protobuf/book"
	"github.com/les-cours/user-api/transform"
	"log"
)

func (r *queryResolver) GetBook(ctx context.Context, in string) (*models.Book, error) {
	log.Println("API: GetBook")

	var res, err = r.BookClient.GetBook(ctx, &pb.BookId{
		Id: in,
	})

	if err != nil {
		return nil, err
	}

	var book = transform.PbBookToGraphBook(res.Book)

	return book, nil
}

func (r *mutationResolver) CreateBook(ctx context.Context, in *models.CreateBookInput) (*models.Output, error) {
	book, err := r.BookClient.CreateBook(ctx, &pb.CreateBookRequest{
		Title:       in.Title,
		Author:      in.Author,
		Description: *in.Description,
		Pages:       int64(*in.Pages),
		IsPublished: *in.IsPublished,
	})
	if err != nil {
		return nil, err
	}

	message := "book " + book.Id + " created!"
	return &models.Output{
		Success: true,
		Message: &message,
	}, nil
}

func (r *mutationResolver) UpdateBook(ctx context.Context, in *models.UpdateBookInput) (*models.Output, error) {

	//v := reflect.ValueOf(*in)
	//t := reflect.TypeOf(*in)
	//values := make([]interface{}, v.NumField())
	//fields := make([]interface{}, t.NumField())
	//
	//for i := 0; i < v.NumField(); i++ {
	//	values[i] = v.Field(i).Interface()
	//	fields[i] = t.Field(i).Name
	//
	//}
	//var book pb.Book{}
	//for i, _ := range values {
	//
	//	if values[i] != nil {
	//
	//	}
	//	fmt.Println(fields[i])
	//	fmt.Println(values[i])
	//}

	var book pb.Book
	if in.Title != nil {
		book.Title = *in.Title
	}
	if in.Author != nil {
		book.Author = *in.Author
	}
	if in.Description != nil {
		book.Description = *in.Description
	}
	if in.Pages != nil {
		book.Pages = int64(*in.Pages)
	}
	if in.IsPublished != nil {
		book.IsPublished = *in.IsPublished
	}

	bookID, err := r.BookClient.UpdateBook(ctx, &book)
	if err != nil {
		return nil, err
	}
	message := "updated success" + bookID.Id
	return &models.Output{
		Success: true,
		Message: &message,
	}, nil
}
func (r *mutationResolver) DeleteBook(ctx context.Context, in string) (*models.Output, error) {
	book, err := r.BookClient.DeleteBook(ctx, &pb.BookId{Id: in})

	message := "delete " + book.Id
	if err != nil {
		return nil, err
	}
	return &models.Output{
		Success: true,
		Message: &message,
	}, nil
}
func (r *queryResolver) GetBooks(ctx context.Context, in *models.PaginationInput) ([]*models.Book, error) {

	fmt.Println("get books starting ...")
	booksRes, err := r.BookClient.GetBooks(ctx, &pb.BookPagination{
		Ids: in.Ids,
		Pagination: &pb.Pagination{
			CurrentPage: int64(in.Pagination.CurrentPage),
			PerPage:     int64(in.Pagination.PerPage),
		},
	})

	if err != nil {
		return nil, err

	}

	var returnedBook []*models.Book
	for _, book := range booksRes.Books {
		returnedBook = append(returnedBook, &models.Book{
			ID:          book.Id,
			Title:       book.Title,
			Author:      book.Author,
			Description: &book.Description,
			IsPublished: &book.IsPublished,
		})
	}

	if err != nil {
		fmt.Println(err)
		return nil, err
	}
	return returnedBook, nil
}
