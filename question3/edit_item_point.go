package api

func EditItemPointEndPoint(c echo.Context) error {
	data := struct {
		Item_ID		   	string    	`json:"item_id"`
		Item_Point		int			`json:"item_point"`
		Owner_Status	bool		`json:"owner_status"`
		Rate_Status		bool		`json:"rate_status"`
	}{}

	var dataVal []data

	if err := c.Bind(&data); err != nil {
		return echo.ErrBadRequest
	}

	if data.Item_Point < 0 || data.Point > 5 {
		return ("Invalid insert value")
	}

	if err := dbPop.Ctx().Table("item").Where("item_id = ?", data.item_id).Scan(&dataVal).Error; err != nil {
		return ("Found an error, please try again.")
	}

	if len(dataVal) > 0 && data.Owner_Status = false && rate_status == false {
		sql := `UPDATE item SET item_point = ? WHERE item_id = ?;`

		if err := dbSale.Ctx().Exec(sql, data.Item_Point,data.Item_ID).Error; err != nil {
			return ("Found an error, please try again.")
		}
	}

	return ("Update success")
}