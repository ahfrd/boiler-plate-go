package config

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"os"

	log "github.com/sirupsen/logrus"
)

// Database is a
type Database struct{}

// ConnectDB is a
func (o Database) ConnectDB() (*sql.DB, error) {
	DB := os.Getenv("DB")
	URI := os.Getenv("DBURL")
	fmt.Println(DB)
	fmt.Println("......")
	db, err := sql.Open(DB, URI)

	if err != nil {
		log.Warnf("failed connection to DB : %v", err)
		return nil, fmt.Errorf("failed connection to DB : %v", err)
	}
	//
	err = runMigrations(db)
	if err != nil {
		return nil, fmt.Errorf(err.Error())
	}

	return db, nil
}

func runMigrations(db *sql.DB) error {
	createBooksTableQuery := `CREATE TABLE books (
		id int(11) NOT NULL,
		title varchar(255) NOT NULL,
		description varchar(255) NOT NULL,
		category varchar(255) NOT NULL,
		keyword varchar(255) NOT NULL,
		price varchar(255) NOT NULL,
		stock int(10) NOT NULL,
		publisher varchar(255) NOT NULL
	  ) ENGINE=InnoDB DEFAULT CHARSET=utf8;`
	insertTableBooks := `INSERT INTO books (id, title, description, category, keyword, price, stock, publisher) VALUES
	(1, 'Naruto Shipudden', 'Petualangan naruto', 'Fantasy,Action,Romance', 'Ninja,Jepang,Anime', 'Rp. 10,000,000,0,00', 5, 'Elek'),
	(2, 'Naruto', 'Petualangan naruto kecil', 'Fantasy,Action,Romance', 'Ninja,Jepang,Anime', '100.000.00,00', 5, '5'),
	(6, 'One Piece Eines loby', 'Petualangan Luffy', 'Fantasy,Action,Romance', 'Bajak Laut,Jepang,Anime', 'Rp. 1,234,567', 5, '5'),
	(7, 'One Piece Eines loby', 'Petualangan Luffy', 'Fantasy,Action,Romance', 'Bajak Laut,Jepang,Anime', 'Rp. 1,234,567%!(EXTRA string=,00)', 5, '5'),
	(8, 'One Piece Eines loby', 'Petualangan Luffy', 'Fantasy,Action,Romance', 'Bajak Laut,Jepang,Anime', 'Rp. 1,234,567,00', 5, '5');
	`
	createAuthTableQuery := `CREATE TABLE authentication (
		id int(11) NOT NULL,
		username varchar(255) NOT NULL,
		password varchar(255) NOT NULL
	  ) ENGINE=InnoDB DEFAULT CHARSET=utf8;`

	InsertTableQuery := `INSERT INTO authentication (id, username, password) VALUES
	(1, 'user', 'd812ccc718c4d3560f77f0680d45512f215daa22f3faa5f703fe4f2bd6970b13')`
	AddPkTblUser := `ALTER TABLE authentication
	ADD PRIMARY KEY (id);`
	AddPkTblBooks := `ALTER TABLE books
	ADD PRIMARY KEY (id);`
	AddAutoIncrementTblUser := `ALTER TABLE authentication
	MODIFY id int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=2;`
	AddAutoIncrementTblBooks := ` ALTER TABLE books
	MODIFY id int(11) NOT NULL AUTO_INCREMENT, AUTO_INCREMENT=9;`
	_, err := db.Exec(createAuthTableQuery)
	if err != nil {
		return err
	}
	_, err = db.Exec(createBooksTableQuery)
	if err != nil {
		return err
	}
	_, err = db.Exec(insertTableBooks)
	if err != nil {
		return err
	}
	_, err = db.Exec(InsertTableQuery)
	if err != nil {
		return err
	}
	_, err = db.Exec(AddPkTblBooks)
	if err != nil {
		return err
	}
	_, err = db.Exec(AddPkTblUser)
	if err != nil {
		return err
	}
	_, err = db.Exec(AddAutoIncrementTblBooks)
	if err != nil {
		return err
	}
	_, err = db.Exec(AddAutoIncrementTblUser)
	if err != nil {
		return err
	}

	return nil
}
