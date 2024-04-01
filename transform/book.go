package transform

import (
	"github.com/les-cours/user-api/graph/models"
	pb "github.com/les-cours/user-api/protobuf/book"
)

func PbBookToGraphBook(pbBook *pb.Book) *models.Book {
	return &models.Book{
		ID:          pbBook.Id,
		Title:       pbBook.Title,
		Author:      pbBook.Author,
		Description: &pbBook.Description,
	}
}
