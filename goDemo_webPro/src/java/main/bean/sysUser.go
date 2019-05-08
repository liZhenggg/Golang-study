package bean

import (
	"time"
)

type sysUser struct {
	nickname    string
	name_CN     string
	name_EN     string
	age         int
	gender      int
	idcard      string
	email       string
	nationality int
	hobby       []int
	birthday    time.Time
}
