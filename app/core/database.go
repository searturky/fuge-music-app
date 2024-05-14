package core

import (
	modelV1 "fuge/app/models/v1"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB(c *ConfYaml) *gorm.DB {
	var err error
	DB, err = gorm.Open(postgres.Open(c.DBDSN), &gorm.Config{})
	sqlDB, _ := DB.DB()
	sqlDB.SetMaxIdleConns(30)
	if err != nil {
		panic(err)
	}
	DB.AutoMigrate(modelV1.Models...)
	initMockData()
	return DB
}

func CloseDB() {
	sqlDB, err := DB.DB()
	if err != nil {
		panic(err)
	}
	sqlDB.Close()
}

func initMockData() {
	initStore()
	initUser()
	initRoom()
	initCategory()
	initService()
}

func initStore() {
	store := &modelV1.Store{
		Name: "赋格音乐中心（武侯校区）",
	}
	// create if not exists
	if err := DB.FirstOrCreate(store, modelV1.Store{Name: store.Name}).Error; err != nil {
		panic(err)
	}
}

func initRoom() {
	store := &modelV1.Store{
		Name: "赋格音乐中心（武侯校区）",
	}
	err := DB.Find(store, "name = ?", "赋格音乐中心（武侯校区）").Error
	if err != nil {
		panic(err)
	}
	room := &modelV1.Room{
		StoreID: store.ID,
		Name:    "钢琴教室1",
	}
	if err := DB.FirstOrCreate(room, &modelV1.Room{
		Name: "钢琴教室1",
	}).Error; err != nil {
		panic(err)
	}
}

func initUser() {
	store := &modelV1.Store{
		Name: "赋格音乐中心（武侯校区）",
	}
	err := DB.Find(store, "name = ?", "赋格音乐中心（武侯校区）").Error
	if err != nil {
		panic(err)
	}
	user := &modelV1.User{
		Nickname: "教师1",
		Stores:   []modelV1.Store{*store},
	}
	if err := DB.FirstOrCreate(user, &modelV1.User{
		Nickname: "教师1",
	}).Error; err != nil {
		panic(err)
	}
}

func initCategory() {
	store := &modelV1.Store{
		Name: "赋格音乐中心（武侯校区）",
	}
	err := DB.Find(store, "name = ?", "赋格音乐中心（武侯校区）").Error
	if err != nil {
		panic(err)
	}
	category := &modelV1.ServiceCategory{
		StoreID: store.ID,
		Name:    "键盘乐器",
	}
	if err := DB.FirstOrCreate(category, &modelV1.ServiceCategory{
		Name: "键盘乐器",
	}).Error; err != nil {
		panic(err)
	}
}

func initService() {
	store := &modelV1.Store{
		Name: "赋格音乐中心（武侯校区）",
	}
	err := DB.Find(store, "name = ?", "赋格音乐中心（武侯校区）").Error
	if err != nil {
		panic(err)
	}
	category := &modelV1.ServiceCategory{
		Name: "键盘乐器",
	}
	err = DB.Find(category, "name = ?", "键盘乐器").Error
	if err != nil {
		panic(err)
	}
	user := &modelV1.User{}
	if err := DB.First(user, &modelV1.User{
		Nickname: "教师1",
	}).Error; err != nil {
		panic(err)
	}
	room := &modelV1.Room{}
	if err := DB.First(room, &modelV1.Room{
		Name: "钢琴教室1",
	}).Error; err != nil {
		panic(err)
	}
	// today
	// today := time.Now()
	// service1Start := time.Date(today.Year(), today.Month(), today.Day(), 9, 0, 0, 0, time.Local)
	// service1End := time.Date(today.Year(), today.Month(), today.Day(), 21, 0, 0, 0, time.Local)
	service1 := &modelV1.Service{
		StoreID:        store.ID,
		CategoryID:     category.ID,
		Name:           "古典钢琴",
		Price:          100.00,
		DailyStartTime: "09:00",
		DailyEndTime:   "21:00",
		TimePeriod:     60,
		Users:          []modelV1.User{*user},
		Rooms:          []modelV1.Room{*room},
	}
	if err := DB.FirstOrCreate(service1, &modelV1.Service{
		Name: "古典钢琴",
	}).Error; err != nil {
		panic(err)
	}

	service2 := &modelV1.Service{
		StoreID:        store.ID,
		CategoryID:     category.ID,
		Name:           "流行钢琴",
		Price:          120,
		DailyStartTime: "10:00",
		DailyEndTime:   "20:00",
		TimePeriod:     60,
		Users:          []modelV1.User{*user},
		Rooms:          []modelV1.Room{*room},
	}

	if err := DB.FirstOrCreate(service2, &modelV1.Service{
		Name: "流行钢琴",
	}).Error; err != nil {
		panic(err)
	}

	service3 := &modelV1.Service{
		StoreID:        store.ID,
		CategoryID:     category.ID,
		Name:           "爵士钢琴",
		Price:          150,
		DailyStartTime: "10:00",
		DailyEndTime:   "20:00",
		TimePeriod:     60,
		Users:          []modelV1.User{*user},
		Rooms:          []modelV1.Room{*room},
	}
	if err := DB.FirstOrCreate(service3, &modelV1.Service{
		Name: "爵士钢琴",
	}).Error; err != nil {
		panic(err)
	}

}
