package link

import (
	"log"

	"github.com/lamhai1401/hackernews/internal/user"
	db "github.com/lamhai1401/hackernews/mysql"
)

// #1
type Link struct {
	ID      string
	Title   string
	Address string
	User    *user.User
}

//#2
func (link Link) Save() int64 {
	//#3
	stmt, err := db.Db.Prepare("INSERT INTO Links(Title,Address) VALUES(?,?)")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	//#4
	res, err := stmt.Exec(link.Title, link.Address)
	if err != nil {
		log.Fatal(err)
	}

	//#5
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal("Error:", err.Error())
	}
	log.Print("Row inserted!")

	return id
}

func GetAll() []Link {
	stmt, err := db.Db.Prepare("select id, title, address from Links")
	if err != nil {
		log.Fatal(err)
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		log.Fatal(err)
	}

	defer rows.Close()
	var links []Link

	for rows.Next() {
		var link Link
		err := rows.Scan(&link.ID, &link.Title, &link.Address)
		if err != nil {
			log.Fatal(err)
		}
		links = append(links, link)
	}

	if err = rows.Err(); err != nil {
		log.Fatal(err)
	}

	return links
}
