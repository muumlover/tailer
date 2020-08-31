package conf

type Connect struct {
	Name  string
	Type  string
	Param interface{}
}

//Save 增加链接
func (c *Configuration) AddConnect(connect *Connect) {
	c.Connects = append(c.Connects, connect)
	c.Save()
}

//Delete 删除链接
func (c *Configuration) Delete(connect *Connect) {
	for index, item := range c.Connects {
		if item == connect {
			c.Connects = append(c.Connects[:index], c.Connects[index+1:]...)
		}
	}
	c.Save()
}
