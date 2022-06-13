package api

const pkgName = "API"

func initDataStore() error {
	dbPop = core.NewDatabase(pkgName, "shippop")
	if err := dbPop.Connect(); err != nil {
		log.Errorln(pkgName, err, "Connect to shippop error")
		return err
	}
	return nil
}

func InitApiRouter(g *echo.Group) error {
	if err := initDataStore(); err != nil {
		log.Errorln(pkgName, err, "connect database error")
	}

	user := g.Group("/users")
	user.POST("", CreateUserEndPoint)

	item := g.Group("/item")
	item.GET("", GetItemEndPoint)
	item.POST("/create", CreateItemEndPoint)
	item.PUT("/edit", EditItemEndPoint)
	item.PUT("/editpoint", EditItemPointEndPoint)
	item.DELETE("/delete", DeleteItemEndPoint)

	return nil
}