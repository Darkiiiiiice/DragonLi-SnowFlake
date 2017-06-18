package service

import (
	"fmt"
	"io/mariomang/github/dao"
	"io/mariomang/github/domain"
)

func GenrateIDService(request *domain.RequestDomain) string {
	maxID, step := dao.GetMaxIDAndStepByWorkID("AAAABBBBCCCCDDDDEEEEFFFF")
	fmt.Println(maxID, step)
	return "HelloWorld"
}
