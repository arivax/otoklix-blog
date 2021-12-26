package blogpost

import (
	repo "go-otoklix-blog/infra/repository/sqlite"
)

const tableName = "blog_post"

// BlogPost model
type BlogPost struct {
	Id           string
	Title        string
	Content      string
	Published_at string
	Created_at   string
	Updated_at   string
}

var (
	qrySql = "SELECT CAST(id AS TEXT) AS id, CAST(title AS TEXT) AS title, CAST(content AS TEXT) AS content, CAST(published_at AS TEXT) AS published_at, CAST(created_at AS TEXT) AS created_at, CAST(updated_at AS TEXT) AS updated_at FROM " + tableName + " p "

	filterAll = " AND UPPER(CONCAT(CAST(id AS TEXT), ' ', CAST(title AS TEXT), ' ', CAST(content AS TEXT), ' ', CAST(published_at AS TEXT), ' ', CAST(created_at AS TEXT), ' ', CAST(updated_at AS TEXT))) "

	Repo = BlogPost{}
)

func (o *BlogPost) FindAll() ([]BlogPost, error) {
	result := []BlogPost{}

	db, err := repo.Connect()
	if err != nil {
		return result, err
	}
	sql := qrySql + " WHERE TRUE ORDER BY published_at"
	rows, err1 := db.Query(sql)
	if err1 != nil { // Exit if the SQL doesn't work for some reason
		return result, err1
	}
	defer rows.Close() // Make sure to cleanup when the program exits

	for rows.Next() {
		err := rows.Scan(&o.Id, &o.Title, &o.Content, &o.Published_at, &o.Created_at, &o.Updated_at)
		if err != nil { // Exit if we get an error
			return result, err
		}
		result = append(result, *o)
	}
	return result, nil
}

func (o *BlogPost) FindById(id string) (BlogPost, error) {
	db, err := repo.Connect()
	if err != nil {
		return BlogPost{}, err
	}
	err = db.QueryRow(qrySql+" WHERE id=?", id).Scan(&o.Id, &o.Title, &o.Content, &o.Published_at, &o.Created_at, &o.Updated_at)
	if err != nil { // Exit if the SQL doesn't work for some reason
		return *o, err
	}
	return *o, nil
}

func (o *BlogPost) CreateOne(data *BlogPostReq) (BlogPost, error) {
	db, err := repo.Connect()
	if err != nil {
		return *o, err
	}
	_, err = db.Exec("INSERT INTO "+tableName+" (title, content, published_at, created_at, updated_at) VALUES (?, ?, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP, CURRENT_TIMESTAMP)", data.Title, data.Content)
	if err != nil { // Exit if we get an error
		return *o, err
	}
	err = db.QueryRow(qrySql+" WHERE title=?", data.Title).Scan(&o.Id, &o.Title, &o.Content, &o.Published_at, &o.Created_at, &o.Updated_at)
	if err != nil { // Exit if the SQL doesn't work for some reason
		return *o, err
	}
	return *o, nil
}

func (o *BlogPost) UpdateOne(data *BlogPostReq, id string) (BlogPost, error) {
	db, err := repo.Connect()
	if err != nil {
		return *o, err
	}
	_, err = db.Exec("UPDATE "+tableName+" SET title=?, content=?, updated_at=CURRENT_TIMESTAMP WHERE id=?", data.Title, data.Content, id)
	if err != nil { // Exit if we get an error
		return *o, err
	}
	err = db.QueryRow(qrySql+" WHERE id=?", id).Scan(&o.Id, &o.Title, &o.Content, &o.Published_at, &o.Created_at, &o.Updated_at)
	if err != nil { // Exit if the SQL doesn't work for some reason
		return *o, err
	}
	return *o, nil
}

func (o *BlogPost) DeleteOne(id string) (BlogPost, error) {
	db, err := repo.Connect()
	if err != nil {
		return *o, err
	}
	row, err1 := Repo.FindById(id)
	if err1 != nil {
		return row, err1
	}
	_, err = db.Exec("DELETE FROM "+tableName+" WHERE id=?", id)
	if err != nil { // Exit if we get an error
		return *o, err
	}
	return row, nil
}
