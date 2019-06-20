package store

import (
	"code.byted.org/stream/stream/logger"
)

func GetG() ([]*G, error){
	db, err := GetDB()
	if err != nil {
		logger.Error("get db failed:", err.Error())
		return nil, err
	}
	var g []*G
	rows, err := db.Query(`select * from [G]`)
	if err != nil {
		logger.Error("Query failed:", err.Error())
		return nil, err
	}
	defer rows.Close()

	//遍历每一行
	for rows.Next() {
		var row = new(G)
		rows.Scan(&row.GID, &row.GName, &row.GSex, &row.GAge, &row.GNum, &row.GPhone, &row.GStatus)
		g = append(g, row)
	}
	logger.Info("%v", g)
	return g, nil
}

