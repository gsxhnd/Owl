package sqlite

import (
	"errors"
	"fmt"
	"strings"
	"time"

	sq "github.com/Masterminds/squirrel"
	"github.com/gsxhnd/owl/server/db/database"
	"github.com/gsxhnd/owl/server/model"
)

func (db *sqliteDB) CreateFolder(folders []model.Folder) error {
	tx, err := db.conn.Begin()
	defer db.txRollback(tx, err)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}

	stmt, err := tx.Prepare(`INSERT INTO folder 
	(code,title,publish_date,director,produce_company,publish_company,series) 
	VALUES (?,?,?,?,?,?,?);`)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}
	defer stmt.Close()

	for _, v := range folders {
		_, err = stmt.Exec(v.Code, v.Title, v.PublishDate, v.Director, v.ProduceCompany, v.PublishCompany, v.Series)
		if err != nil {
			db.logger.Errorf(err.Error())
			return err
		}
	}

	err = tx.Commit()
	return err
}

func (db *sqliteDB) DeleteFolders(ids []uint) error {
	tx, err := db.conn.Begin()
	defer db.txRollback(tx, err)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}

	stmt, err := tx.Prepare(`DELETE FROM movie WHERE id IN (?` + strings.Repeat(`,?`, len(ids)-1) + `);`)
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

func (db *sqliteDB) UpdateFolder(movie *model.Folder) error {
	tx, err := db.conn.Begin()
	defer db.txRollback(tx, err)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}

	stmt, err := tx.Prepare(`UPDATE movie SET 
	code=?, title=?,cover=?,publish_date=?,director=?,produce_company=?,publish_company=?,series=?, updated_at=? 
	WHERE id=?;`)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(
		movie.Code,
		movie.Title,
		movie.Cover,
		movie.PublishDate,
		movie.Director,
		movie.ProduceCompany,
		movie.PublishCompany,
		movie.Series,
		time.Now(),
		movie.Id,
	)
	if err != nil {
		db.logger.Errorf(err.Error())
		return err
	}

	err = tx.Commit()
	return err
}

func (db *sqliteDB) GetFolders(p *database.Pagination, filter ...string) ([]model.Movie, error) {
	query := sq.Select("*").From("movie")
	if p != nil {
		query = query.Limit(p.Limit).Offset(p.Offset)
	}

	if len(filter)%2 != 0 {
		return nil, errors.New("123")
	}

	for i := 0; i < len(filter); i += 2 {
		key := filter[i]
		value := filter[i+1]
		query = query.Where(fmt.Sprintf("%s like ?", key), fmt.Sprint("%", value, "%"))
	}

	sql, args, err := query.ToSql()
	if err != nil {
		return nil, err
	}

	rows, err := db.conn.Query(sql, args...)
	if err != nil {
		db.logger.Errorf(err.Error())
		return nil, err
	}

	var dataList []model.Folder
	for rows.Next() {
		var data = model.Folder{}
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
