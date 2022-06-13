func CreateItemEndPoint(c echo.Context) error {
	data := struct {
		Item_ID		   	string    	`json:"item_id"`
		Item_Name   	string    	`json:"item_name"`
		Price	    	string    	`json:"price"`
		Item_Detail   	string    	`json:"item_detail"`
		Item_Status		string		`json:"item_status"`
		Owner_Status	bool		`json:"owner_status"`
	}{}

	var dataVal []data

	if err := c.Bind(&data); err != nil {
		return echo.ErrBadRequest
	}

	if err := dbPop.Ctx().Table("item").Where("item_id = ?", data.item_id).Scan(&dataVal).Error; err != nil {
		return ("Found an error, please try again.")
	}

	if len(dataVal) > 0 && data.Owner_Status = true {
		sql := `UPDATE item SET item_name = ?, price = ?, item_detail = ?, item_status = ? WHERE item_id = ?;`

		if err := dbSale.Ctx().Exec(sql, data.Item_Name, data.Price, data.Item_Detail, data.Item_Status, data.Item_ID).Error; err != nil {
			return ("Found an error, please try again.")
		}
	}

	return ("Update success")
}