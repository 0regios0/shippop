package api

func CreateUserEndPoint(c echo.Context) error {
	data := struct {
		User_Name    string    `json:"user_name"`
		Email   	 string    `json:"email"`
		Passsword    string    `json:"password"`
	}{}

	var dataVal []data

	if err := c.Bind(&data); err != nil {
		return echo.ErrBadRequest
	}

	if data.Email == "" || data.Password == "" {
		return ("Invalid insert value")
	}

	if err := dbPop.Ctx().Table("user").Where("email = ?", data.Email).Scan(&dataVal).Error; err != nil {
		return ("Found an error, please try again.")
	}
	if len(dataVal) > 0 {
		return c.JSON(http.StatusBadRequest, m.Result{Message: "Duplicate user"})
	}

	if err := dbPop.Ctx().Table("user").Create(&data).Error; err != nil {
		return ("insert error, please try again.")
	}

	return ("CREATE success")
}