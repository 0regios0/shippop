package api

func GetItemEndPoint(c echo.Context) error {
	data := struct {
		Item_Id   		int    		`json:"item_id"`
		Item_Name   	string    	`json:"item_name"`
		Item_Detail   	string    	`json:"item_detail"`
		Price	    	string    	`json:"price"`
		Discount    	string    	`json:"discount"`
		User_Id    		string    	`json:"User_id"`
		Item_Status		string		`json:"item_status"`
		Item_Point		int			`json:"item_point"`
	}{}

	var dataVal []data
	
	sql := `select * from item`
	
	if err := dbSale.Ctx().Raw(sql).Scan(&dataVal).Error; err != nil {
		errr += 1
	}

	return (dataVal)

}