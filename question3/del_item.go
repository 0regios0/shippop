package api

func EditItemPointEndPoint(c echo.Context) error {
	data := struct {
		Item_ID		   	string    	`json:"item_id"`
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
		if err := dbSale.Ctx().Exec("DELETE FROM item WHERE item_id = ?", data.Id).Error; err != nil {
			return ("Found an error, please try again.")
		}
	}

	return ("Delete success")
}