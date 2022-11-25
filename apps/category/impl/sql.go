package impl

const (
	CreateSql        = "insert into category (create_at,update_at,name,key_picture) values(?,?,?,?)"
	QuerySql         = "select * from category where name=?"
	QuerySqlNotParam = "select * from category"
)
