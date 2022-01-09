package calendar

import (
	"arm_go/model"
	"encoding/json"
	"fmt"
)

type Service interface {
	Read(tanggal string, limit int, last int) ([]model.Calendar, error)
	ReadJadwal(jadwalId uint) (model.CalendarResponse, error)
	Create(calendar model.Calendar) (model.Calendar, uint, error)
	Update(id string, calendar model.CalendarUpdate) (bool, error)
	Delete(dataType string, id string) (bool, error)
}

type service struct {
	repository Repository
}

func SetupService(repo Repository) *service {
	return &service{repo}
}

func (s service) Read(tanggal string, limit int, last int) ([]model.CalendarResponse, error) {

	var dataPengisi []string
	var calendar []model.CalendarResponse

	data, err := s.repository.Read(tanggal, limit, last)

	for _, q := range data {
		var CalendarResponse []model.DetailJadwalResponse
		for _, v := range q.DetailJadwal {

			pengisiQuery := []byte(v.Pengisi)
			err = json.Unmarshal(pengisiQuery, &dataPengisi)

			CalendarResponse = append(CalendarResponse, model.DetailJadwalResponse{
				ID:              v.ID,
				Kota:            v.Kota,
				Lokasi:          v.Lokasi,
				Alamat:          v.Alamat,
				Waktu:           v.Waktu,
				PemangkuManaqib: v.PemangkuManaqib,
				Pengisi:         dataPengisi,
			})
		}

		calendar = append(calendar, model.CalendarResponse{
			Tanggal: q.Tanggal,
			Jadwal:  CalendarResponse,
		})
	}

	return calendar, err
}

func (s service) ReadJadwal(jadwalId string) (model.CalendarResponse, error) {

	calendar, err := s.repository.ReadJadwal(jadwalId)

	var jadwalData []model.DetailJadwalResponse
	var pengisiData []string

	fmt.Println("==============================")
	fmt.Println("detail error: ", calendar)
	fmt.Println(calendar)
	fmt.Println("==============================")

	if calendar.DetailJadwal == nil {
		return model.CalendarResponse{}, fmt.Errorf("data not found")
	}

	jadwalData = append(jadwalData, model.DetailJadwalResponse{
		ID:              calendar.DetailJadwal[0].ID,
		Kota:            calendar.DetailJadwal[0].Kota,
		Lokasi:          calendar.DetailJadwal[0].Lokasi,
		Alamat:          calendar.DetailJadwal[0].Alamat,
		Waktu:           calendar.DetailJadwal[0].Waktu,
		PemangkuManaqib: calendar.DetailJadwal[0].PemangkuManaqib,
		Pengisi:         pengisiData,
	})

	jadwal := model.CalendarResponse{
		Tanggal: calendar.Tanggal,
		Jadwal:  jadwalData,
	}

	return jadwal, err
}

func (s service) Create(calendarRequest model.CalendarRequest) (model.Calendar, uint, error) {

	pengisi, _ := json.Marshal(calendarRequest.Pengisi)

	calendar := model.Calendar{
		Tanggal: calendarRequest.Tanggal,
	}

	detailJadwal := model.DetailJadwal{
		Kota:            calendarRequest.Kota,
		Lokasi:          calendarRequest.Lokasi,
		Alamat:          calendarRequest.Alamat,
		Waktu:           calendarRequest.Waktu,
		PemangkuManaqib: calendarRequest.PemangkuManaqib,
		Pengisi:         string(pengisi),
	}

	result, jadwalId, err := s.repository.Create(calendar, detailJadwal)

	return result, jadwalId, err
}

func (s service) Update(id string, calendar model.CalendarUpdate) (bool, error) {

	var pengisiConverted string

	if calendar.Pengisi != nil {
		pengisiByte, err := json.Marshal(calendar.Pengisi)

		if err != nil {
			return false, err
		}

		pengisiConverted = string(pengisiByte)

		fmt.Println(pengisiConverted)
	}

	dataCalendar := model.CalendarUpdateDb{
		Kota:            calendar.Kota,
		Lokasi:          calendar.Lokasi,
		Alamat:          calendar.Alamat,
		Waktu:           calendar.Waktu,
		PemangkuManaqib: calendar.PemangkuManaqib,
		Pengisi:         pengisiConverted,
	}

	isUpdated, err := s.repository.Update(id, dataCalendar)

	return isUpdated, err
}

func (s service) Delete(dataType string, id string) (bool, error) {

	status, err := s.repository.Delete(dataType, id)

	return status, err
}
