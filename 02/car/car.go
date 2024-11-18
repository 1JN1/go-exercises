package car

import (
	"fmt"
	"time"
)

// Car 汽车
// Brand:品牌名 Country:国家 Price:价格 Parameters:参数列表 ProductionDate:生产日期
type Car struct {
	Brand string

	Country string

	Price float64

	Parameters map[string][]string

	ProductionDate time.Time
}

// Info 输出汽车信息字符串
func (c Car) Info() string {

	timeFormat := c.ProductionDate.Format("2006年1月2日")

	paramStr := "{\n"

	for k, values := range c.Parameters {
		paramStr += "\t\t\t" + k + ":"
		if len(values) > 1 {
			paramStr += "["

			for idx, value := range values {
				if idx != len(values)-1 {
					paramStr += value + "、"
				} else {
					paramStr += value
				}
			}

			paramStr += "]"
		} else if len(values) == 1 {
			paramStr += values[0]
		}

		paramStr += "\n"
	}

	paramStr += "\t\t }"

	return fmt.Sprintf(
		"品牌：%s\n"+
			"国家：%s\n"+
			"价格：%.5f\n"+
			"参数列表：%s\n"+
			"生产时间：%s",
		c.Brand, c.Country, c.Price, paramStr, timeFormat)
}
