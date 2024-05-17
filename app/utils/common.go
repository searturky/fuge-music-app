package utils

import (
	"crypto/rand"
	"encoding/hex"
	"fuge/app/core"
	"strconv"
	"strings"

	"github.com/google/uuid"
)

func GetRandomName(prefix string) string {
	return prefix + "_" + uuid.New().String()[:8]
}

func GetRandomString(length int) string {
	bytes := make([]byte, length)
	_, err := rand.Read(bytes)
	if err != nil {
		return ""
	}

	return hex.EncodeToString(bytes)
}

func GenRandomNickname() string {
	prefix := core.GetConf().DefaultUserPrefix
	randomName := GetRandomName(prefix)
	return randomName
}

func GetMinuteFromTimeStr(timeStr string) (int, error) {
	if len(timeStr) < 5 {
		return 0, nil
	}
	timeArr := strings.Split(timeStr, ":")
	if len(timeArr) != 2 {
		return 0, nil
	}
	hour, err := strconv.Atoi(timeArr[0])
	if err != nil {
		return 0, err
	}
	minute, err := strconv.Atoi(timeArr[1])
	if err != nil {
		return 0, err
	}
	return hour*60 + minute, nil
}

type BitMap struct {
	bits []byte
	vmax uint
}

func NewBitMap(maxVal ...uint) *BitMap {
	var max uint = 86400
	if len(maxVal) > 0 && maxVal[0] > 0 {
		max = maxVal[0]
	}
	bm := &BitMap{}
	bm.vmax = max
	sz := (max + 7) / 8
	bm.bits = make([]byte, sz)
	return bm
}

// 添加
func (bm *BitMap) Set(num uint) {
	if num > bm.vmax {
		bm.vmax += 1024
		if bm.vmax < num {
			bm.vmax = num
		}

		dd := int(num+7)/8 - len(bm.bits)
		if dd > 0 {
			tmpArr := make([]byte, dd)
			bm.bits = append(bm.bits, tmpArr...)
		}
	}

	//将1左移num%8后，然后和以前的数据做|，这样就替换成1了
	bm.bits[num/8] |= 1 << (num % 8)
}

// 删除
func (bm *BitMap) UnSet(num uint) {
	if num > bm.vmax {
		return
	}
	//&^:将1左移num%8后，然后进行与非运算，将运算符左边数据相异的位保留，相同位清零
	bm.bits[num/8] &^= 1 << (num % 8)
}

// 判断是否存在
func (bm *BitMap) Check(num uint) bool {
	if num > bm.vmax {
		return false
	}
	//&:与运算符，两个都是1，结果为1
	return bm.bits[num/8]&(1<<(num%8)) != 0
}
