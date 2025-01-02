package sqlite

import (
	"strings"

	"github.com/gsxhnd/owl/server/db/database"
	"github.com/gsxhnd/owl/server/model"
)

func (db *sqliteDB) CreateFiles(files []model.File) error {
	tx, err := db.conn.Begin()
	defer db.txRollback(tx, err)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}

	stmt, err := tx.Prepare(`INSERT INTO file 
	(code,title,title_cn, cover, publish_date, created_at, updated_at) 
	VALUES (?,?,?,?,?,?,?);`)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}
	defer stmt.Close()

	// for _, v := range animes {
	// 	_, err = stmt.Exec(v.Code, v.Title, v.TitleCn, v.Cover, v.PublishDate, v.CreatedAt, v.UpdatedAt)
	// 	if err != nil {
	// 		db.logger.Errorf(err.Error())
	// 		return err
	// 	}
	// }

	err = tx.Commit()
	return err
}

func (db *sqliteDB) DeleteFiles(ids []uint) error {
	tx, err := db.conn.Begin()
	defer db.txRollback(tx, err)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}

	stmt, err := tx.Prepare(`DELETE FROM anime WHERE id IN (?` + strings.Repeat(`,?`, len(ids)-1) + `);`)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}
	defer stmt.Close()

	var args []interface{}
	for _, id := range ids {
		args = append(args, id)
	}

	_, err = stmt.Exec(args...)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}

	err = tx.Commit()
	return err
}

func (db *sqliteDB) UpdateFile(anime model.File) error {
	tx, err := db.conn.Begin()
	defer db.txRollback(tx, err)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}

	stmt, err := tx.Prepare(`UPDATE anime SET 
	title = ?, 
	cover = ?, 
	publish_date = ?, 
	updated_at = ? 
	WHERE id = ?;`)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(anime.Title, anime.Cover, anime.PublishDate, anime.UpdatedAt, anime.Id)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}

	err = tx.Commit()
	return err
}

func (db *sqliteDB) GetAnimes(p *database.Pagination) ([]model.Anime, error) {
	rows, err := db.conn.Query("SELECT * FROM anime;")
	if err != nil {
		db.logger.Errorf(err.Error())
		return nil, err
	}
	defer rows.Close()

	var dataList []model.File
	for rows.Next() {
		var data = model.File{}
		if err := rows.Scan(
			&data.Id,
			&data.CreatedAt,
			&data.UpdatedAt,
		); err != nil {
			db.logger.Errorf(err.Error())
			return nil, err
		}
		dataList = append(dataList, data)
	}
	return dataList, nil
}
