package repository

import (
	"asia-quest/config"
	"asia-quest/entity"
	"asia-quest/entity/request"
	"asia-quest/entity/response"
	"database/sql"
	"fmt"
	"strings"

	"github.com/labstack/gommon/log"
)

type booksRepository struct {
}

func NewBooksRepository() entity.BooksRepository {
	return &booksRepository{}
}

func (r *booksRepository) Create(params *request.CreateRequest) error {

	var err error
	var res sql.Result
	db, err := config.Database.ConnectDB(config.Database{})
	if err != nil {
		fmt.Println(err)
		return fmt.Errorf("%s", err)
	}
	defer db.Close()
	if err != nil {
		return fmt.Errorf("%s", err)
	}

	var categoryS []string
	for _, itemC := range params.Category {
		categoryS = append(categoryS, itemC.Data)
	}
	category := strings.Join(categoryS, ",")

	var keywordS []string
	for _, itemK := range params.Keyword {
		keywordS = append(keywordS, itemK.Data)
	}
	keyword := strings.Join(keywordS, ",")

	queryInsert := "INSERT INTO books (title,description,category,keyword,price,stock,publisher) values (?,?,?,?,?,?,?)"
	res, err = db.Exec(queryInsert, params.Title, params.Description, category, keyword, params.Price, params.Stock, params.Publisher)

	if err != nil {
		return fmt.Errorf("failed to insert error_general on books SQL : %v", err)
	}
	count, err := res.RowsAffected()
	if err != nil {
		return err
	}
	if count == 0 {
		return fmt.Errorf("terjadi kesalahan pada database")

	}
	fmt.Println(count)
	return nil
}
func (r *booksRepository) Read(params *request.ReadRequest) (*response.ReadQueryEntity, error) {
	var data response.ReadQueryEntity
	var id sql.NullString
	var title sql.NullString
	var description sql.NullString
	var category sql.NullString
	var keyword sql.NullString
	var price sql.NullString
	var stock sql.NullInt64
	var publisher sql.NullString

	db, err := config.Database.ConnectDB(config.Database{})
	if err != nil {
		return nil, err
	}
	var query string = fmt.Sprintf(`select id,title,description,category,keyword,price,stock,publisher from books where id = "%s"`, params.Id)
	db.QueryRow(query).Scan(
		&id,
		&title,
		&description,
		&category,
		&keyword,
		&price,
		&stock,
		&publisher,
	)
	fmt.Println(query)
	data.Id = id.String
	data.Title = title.String
	data.Description = description.String
	data.Category = category.String
	data.Keyword = keyword.String
	data.Price = price.String
	data.Stock = int(stock.Int64)
	data.Publisher = publisher.String
	defer db.Close()
	if err != nil && err != sql.ErrNoRows {
		return nil, fmt.Errorf("failed Select SQL for books : %v", err)
	}

	return &data, nil
}

func (r *booksRepository) Update(params *request.UpdateRequest) error {

	db, err := config.Database.ConnectDB(config.Database{})
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	var categoryBook []request.CategoryBook
	var categoryS []string
	for _, itemC := range categoryBook {
		categoryS = append(categoryS, itemC.Data)
	}

	category := strings.Join(categoryS, ", ")
	var keywordBook []request.KeywordBook
	var keywordS []string
	for _, itemK := range keywordBook {
		keywordS = append(keywordS, itemK.Data)
	}
	keyword := strings.Join(keywordS, ",")

	var queryUpdate string = fmt.Sprintf("UPDATE books set title = ?,description = ?,category = ?,keyword = ?,price = ?,stock = ?,publisher = ? where id = ?")
	fmt.Println(queryUpdate)
	res, err := db.Exec(queryUpdate, params.Title, params.Description, category, keyword, params.Price, params.Stock, params.Id)
	defer db.Close()
	if err != nil {
		return fmt.Errorf("failed to books SQL : %v", err)
	}

	if err != nil {
		return fmt.Errorf("failed to books SQL : %v", err)
	}

	counter, err := res.RowsAffected()
	if err != nil {
		return fmt.Errorf("failed to populate status updated : %v", err)
	}
	fmt.Println(counter)
	if counter == 0 {
		return fmt.Errorf("failed to update books SQL")
	}
	return nil
}
func (r *booksRepository) Delete(params *request.DeleteRequest) error {

	var err error
	var res sql.Result
	var prepare *sql.Stmt
	db, err := config.Database.ConnectDB(config.Database{})
	if err != nil {
		return fmt.Errorf("%s", err)
	}
	defer db.Close()
	if err != nil {
		return err
	}
	// var idS []string
	// for _, itemC := range params.Id {
	// 	idS = append(idS, itemC.Data)
	// }
	// idArr := strings.Join(idS, ",")
	// fmt.Println("******")
	// fmt.Println(idArr)
	for _, item := range params.Id {
		queryUpdate := "DELETE FROM books WHERE id = ?;"
		prepare, err = db.Prepare(queryUpdate)
		if err != nil {
			log.Warnf("failed to delete tbl books SQL : %v", err)
			return err
		}
		res, err = prepare.Exec(item.Data)
	}
	if err != nil {
		log.Warnf("failed to delete limit on tbl books SQL : %v", err)
		return err
	}
	count, err := res.RowsAffected()
	if err != nil {
		log.Warnf("failed to populate status delete : %v", err)
		return err
	}
	fmt.Println(count)
	return nil
}
