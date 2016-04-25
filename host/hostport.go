package host

type HOST_PORT struct {
	Host string
	Port int
}


func GetHostPort(){
	r.POST("/host_port", func(c *gin.Context) {
		var a HOST_PORT
		err := c.BindJSON(&a)
		if err != nil {
			c.JSON(400, gin.H{"message": "Bad Request"})
			return
		}

		flag := false
		resp, err := client.ListTasks()
		if err != nil {
			c.JSON(400, gin.H{"message": err})
			return
		}
		for _, t := range resp.Tasks {
			if t.Host == a.Host {
				for _, p := range t.Ports {
					//fmt.Println("DEBUG port ", p)
					if p == a.Port {
						resp, err := client.GetApp(t.AppID)
						if err != nil {
							c.JSON(400, gin.H{"message": "Bad Request"})
							return
						}
						c.JSON(200, resp)
						flag = true
					}
				}
			}
		}
		if !flag {
			c.JSON(400, gin.H{"message": "No Record Found"})
			return
		}
	}
}