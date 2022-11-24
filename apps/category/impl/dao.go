package impl

import (
	"github.com/lifangjunone/go-micro/apps/category"
)

func (i *impl) save(ins *category.Category) error {
	stmt, err := i.db.Prepare(CreateSql)
	defer stmt.Close()
	if err != nil {
		return err
	}
	stmt.Exec(ins.Data.Name, ins.Data.KeyPicture)
	return nil
}

func (i *impl) query(query *category.QueryCategoryRequest) (*category.CategorySet, error) {
	rows, err := i.db.Query(QuerySql, query.Keyword)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	data := category.NewCategorySet()
	for rows.Next() {
		item := category.NewDefaultCategory()
		rows.Scan(item.Id, item.Data.Name, item.Data.KeyPicture)
		data.Add(item)
	}
	return data, err
}
