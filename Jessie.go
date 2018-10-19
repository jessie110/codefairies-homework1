//写三个函数，给定一个[]Score，找出:找出某个特定班级的，找出名字长度超过10的

package main

import (
	"fmt"
)

type score struct {
	name  string
	class string
	score int
}

const (
	certainScore   = 59
	certainClass   = "two"
	certainNamelen = 10
)

var schoolReports = []score{
	{"tom", "one", 40},
	{"jerry", "one", 50},
	{"john", "two", 85},
	{"jessie.xie", "two", 100},
	{"teddddddddddd", "three", 60},
}

func main() {
	a := getCertainScoreReports(schoolReports, certainScore)
	b := getCertainClassReports(schoolReports, certainClass)
	c := getCertainNameReports(schoolReports, certainNamelen)
	fmt.Println(a, b, c)
}

//分数大于59
func getCertainScoreReports(reports []score, condition int) []score {
	var returnReports []score
	for _, report := range reports {
		if report.score > condition {
			returnReports = append(returnReports, report)
		}
	}
	return returnReports
}

//某个特定班级的
func getCertainClassReports(reports []score, condition string) []score {
	var returnReports []score
	for _, report := range reports {
		if report.class == condition {
			returnReports = append(returnReports, report)
		}
	}
	return returnReports
}

//找出名字长度超过10的
func getCertainNameReports(reports []score, condition int) []score {
	var returnReports []score
	for _, report := range reports {
		if len(report.name) > condition {
			returnReports = append(returnReports, report)
		}
	}
	return returnReports
}









// func getCertainReport(reports []score, selector interface{}, condition interface{}) []score {
// 	var returnReports []score
// 	for _, report := range reports {
// 		temp := func(selector interface{}, condition interface{}) interface{} {
// 			return selector=condition
// 		}
// 		if temp {
// 			returnReports = append(returnReports, report)
// 		}
// 	}
// 	return returnReports
// }

// func getScoreSelector(reports []score) interface{} {
// 	var returnSelector interface{}
// 	for _, report := range reports {
// 		returnSelector = report.score
// 	}
// 	return returnSelector
// }
