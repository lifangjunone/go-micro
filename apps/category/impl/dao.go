package impl

import (
	"database/sql"
	"github.com/lifangjunone/go-micro/apps/category"
)

func (i *impl) save(ins *category.Category) error {
	stmt, err := i.db.Prepare(CreateSql)
	defer stmt.Close()
	if err != nil {
		return err
	}
	stmt.Exec(ins.CreateAt, ins.UpdateAt, ins.Data.Name, ins.Data.KeyPicture)
	return nil
}

func (i *impl) query(query *category.QueryCategoryRequest) (*category.CategorySet, error) {
	rows := &sql.Rows{}
	var err error
	if query.Keyword == "" {
		rows, err = i.db.Query(QuerySqlNotParam)
	} else {
		rows, err = i.db.Query(QuerySql, query.Keyword)
	}
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	data := category.NewCategorySet()
	for rows.Next() {
		item := category.NewDefaultCategory()
		err = rows.Scan(&item.Id, &item.CreateAt, &item.UpdateAt, &item.Data.Name, &item.Data.KeyPicture)
		if err != nil {
			i.log.Error(err.Error())
			continue
		}
		data.Add(item)
	}
	return data, err
}
