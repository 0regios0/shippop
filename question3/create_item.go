package api

func CreateItemEndPoint(c echo.Context) error {
	data := struct {
		Item_Name   	string    	`json:"item_name"`
		Item_Detail   	string    	`json:"item_detail"`
		Price	    	string    	`json:"price"`
		Discount    	string    	`json:"discount"`
		User_Id    		string    	`json:"User_id"`
		Item_Status		string		`json:"item_status"`
		Item_Point		int			`json:"item_point"`
	}{}

	var dataVal []data

	if err := c.Bind(&data); err != nil {
		return echo.ErrBadRequest
	}

	if data.Item_Name == "" || data.User_Id == "" {
		return ("Invalid insert value")
	}

	if err := dbPop.Ctx().Table("user").Where("user_id = ?", data.User_Id).Scan(&dataVal).Error; err != nil {
		return ("Found an error, please try again.")
	}
	if len(dataVal) > 0 {
		if err := dbPop.Ctx().Table("item").Create(&data).Error; err != nil {
			return ("insert error, please try again.")
		}
	}

	return ("CREATE success")
}