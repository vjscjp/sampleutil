package app


type APP struct {
	Appid string
}


func GetApp(){
	r.POST("/app", func(c *gin.Context) {
		var a APP
		err := c.BindJSON(&a)
		if err != nil {
			c.JSON(400, gin.H{"message": "Bad Request"})
			return
		}

		appid := strings.Replace(a.Appid, "/", "", -1)
		resp, err := client.GetApp(appid)
		if err != nil {
			fmt.Println("Invalid Details : " + err.Error())
			return
		}

		c.JSON(200, resp)
	})
}