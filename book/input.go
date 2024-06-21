package book

type BookInput struct {
	Id       int    `json:"id_book" validate:"required,gt=0"`
	Title    string `json:"title" validate:"required,min=5"`
	SubTitle string `json:"sub_title" validate:"required"`
	Price    int    `json:"price" validate:"required,number"`
}
