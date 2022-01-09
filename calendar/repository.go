package calendar

import (
	"arm_go/model"
	"fmt"

	"gorm.io/gorm"
)

type Repository interface {
	Read(tanggal string, limit int, last int) ([]model.Calendar, error)
	ReadJadwal(jadwalId string) (model.Calendar, error)
	Create(calendar model.Calendar, detailJadwal model.DetailJadwal) (model.Calendar, uint, error)
	Update(id string, calendar model.CalendarUpdateDb) (bool, error)
	Delete(dataType string, id string) (bool, error)
}

type repository struct {
	db *gorm.DB
}

func SetupRepository(db *gorm.DB) *repository {
	return &repository{db}
}

func (repo repository) Read(tanggal string, limit int, last int) ([]model.Calendar, error) {

	var calendarArray []model.Calendar
	var detailJadwal []model.DetailJadwal
	var calendar []model.Calendar
	var err error

	if len(tanggal) == 7 {
		err = repo.db.Limit(32).Where("tanggal LIKE ?", "%"+tanggal).Find(&calendarArray).Error
	}

	if len(tanggal) == 10 {
		err = repo.db.Limit(32).Where("tanggal = ?", tanggal).Find(&calendarArray).Error
	}

	for _, v := range calendarArray {
		err = repo.db.Limit(limit).Where("calendar_refer = ?", v.ID).Where("id > ?", last).Find(&detailJadwal).Error
		calendar = append(calendar, model.Calendar{
			Tanggal:      v.Tanggal,
			DetailJadwal: detailJadwal,
		})
	}

	return calendar, err
}

func (repo repository) ReadJadwal(jadwalId string) (model.Calendar, error) {

	var calendar model.Calendar
	var detailJadwal model.DetailJadwal
	var detailJadwalData []model.DetailJadwal

	err := repo.db.First(&detailJadwal, jadwalId).Error

	if err != nil {

		fmt.Println("=================================")
		fmt.Println(err)
		fmt.Println("=================================")

		return calendar, err
	}

	detailJadwalData = append(detailJadwalData, detailJadwal)

	err = repo.db.First(&calendar, detailJadwal.CalendarRefer).Error

	if err != nil {

		fmt.Println("=================================")
		fmt.Println(err)
		fmt.Println("=================================")

		return calendar, err
	}

	calendar.DetailJadwal = detailJadwalData

	fmt.Println(calendar)
	fmt.Println("=================================")

	return calendar, err
}

func (repo repository) Create(calendar model.Calendar, detailJadwal model.DetailJadwal) (model.Calendar, uint, error) {

	var findCalendar model.Calendar

	err := repo.db.Where("tanggal = ? ", &calendar.Tanggal).First(&findCalendar).Error

	if err != nil {
		err = repo.db.Create(&calendar).Error

		if err != nil {
			fmt.Println("=> Error Detail: ", err)
		}
	}

	createCalendar := repo.db.Where("lokasi = ?", detailJadwal.Lokasi).Where("alamat = ?", detailJadwal.Alamat).Where("calendar_refer = ?", findCalendar.ID).First(&detailJadwal)
	err = createCalendar.Error

	if err != nil && err.Error() == "record not found" {
		err = nil
		err = repo.db.Model(&findCalendar).Association("DetailJadwal").Append(&detailJadwal)
	}

	jadwalId := detailJadwal.ID

	return calendar, jadwalId, err
}

func (repo repository) Update(id string, calendar model.CalendarUpdateDb) (bool, error) {

	var getCalendar model.DetailJadwal

	err := repo.db.Where("id = ? ", id).First(&getCalendar).Error

	if err != nil {
		return false, err
	}

	fmt.Println(getCalendar.Alamat)

	if calendar.Alamat != "" {
		getCalendar.Alamat = calendar.Alamat
	}

	if calendar.Kota != "" {
		getCalendar.Kota = calendar.Kota
	}

	if calendar.Lokasi != "" {
		getCalendar.Lokasi = calendar.Lokasi
	}

	if calendar.PemangkuManaqib != "" {
		getCalendar.PemangkuManaqib = calendar.PemangkuManaqib
	}

	if calendar.Waktu != "" {
		getCalendar.Waktu = calendar.Waktu
	}

	if calendar.Pengisi != "" {
		getCalendar.Pengisi = calendar.Pengisi
	}

	err = repo.db.Save(&getCalendar).Error

	if err != nil {
		return false, err
	}

	return true, nil
}

func (repo repository) Delete(dataType string, id string) (bool, error) {

	var deletedJadwal model.DetailJadwal
	var deletedCalendar model.Calendar

	fmt.Println("===", id)

	if dataType == "event" {
		err := repo.db.Delete(&deletedJadwal, id).Error
		if err != nil {
			return false, err
		}
	}

	if dataType == "date" {

		fmt.Println("=== 2", id)
		err := repo.db.Delete(&deletedCalendar, id).Error
		if err != nil {
			return false, err
		}
	}

	return true, nil

}
