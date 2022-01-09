package handler

import (
	"arm_go/calendar"
	"arm_go/db"
	"arm_go/model"
	"fmt"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func ReadCalendar(c *gin.Context) {

	dt := time.Now()

	tanggal := c.DefaultQuery("tanggal", dt.Format("02-01-2006"))
	limit, _ := strconv.ParseInt(c.DefaultQuery("limit", "10"), 10, 0)
	last, _ := strconv.ParseInt(c.DefaultQuery("last", "0"), 10, 0)

	database, err := db.Setup()

	if err != nil {

		fmt.Println("error detail: ", err)

		c.JSON(500, gin.H{
			"status":  500,
			"message": "Internal Server Error",
			"data":    "database connection problem #001",
		})

		return
	}

	repo := calendar.SetupRepository(database)
	service := calendar.SetupService(repo)
	readResult, err := service.Read(tanggal, int(limit), int(last))

	if err != nil {
		fmt.Println("=> Error Details: ", err)

		c.JSON(500, gin.H{
			"status":  500,
			"message": "Internal Server Error",
		})
		return
	}

	if readResult == nil {
		c.JSON(404, gin.H{
			"status":  404,
			"message": "Data Not Found",
			"data":    nil,
		})

		return
	}

	c.JSON(200, gin.H{
		"status":  200,
		"message": "Success",
		"data":    readResult,
	})
}

func CreateCalendar(c *gin.Context) {

	database, err := db.Setup()

	if err != nil {

		fmt.Println("error detail: ", err)

		c.JSON(500, gin.H{
			"status":  500,
			"message": "Internal Server Error",
			"data":    "database connection problem #001",
		})

		return
	}

	repo := calendar.SetupRepository(database)
	service := calendar.SetupService(repo)

	var json model.CalendarRequest

	err = c.ShouldBindJSON(&json)

	if err != nil {

		fmt.Println("error detail: ", err)

		c.JSON(500, gin.H{
			"status":  400,
			"message": "Bad Request",
			"data":    "incomplete data #003",
		})

		return
	}

	_, jadwalId, err := service.Create(json)

	if err != nil {

		fmt.Println("error detail: ", err)

		c.JSON(500, gin.H{
			"status":  500,
			"message": "Internal Server Error",
			"data":    "database writing problem #002",
		})

		return
	}

	c.JSON(200, gin.H{
		"status":  200,
		"message": "Success",
		"data":    jadwalId,
	})
}

func ReadJadwal(c *gin.Context) {

	database, err := db.Setup()

	if err != nil {

		fmt.Println("error detail: ", err)

		c.JSON(500, gin.H{
			"status":  500,
			"message": "Internal Server Error",
			"data":    "database connection problem #001",
		})

		return
	}

	jadwalId := c.Param("id")

	repo := calendar.SetupRepository(database)
	service := calendar.SetupService(repo)

	jadwal, err := service.ReadJadwal(jadwalId)

	if err != nil {

		fmt.Println("==============================")
		fmt.Println("detail error: ", err)
		fmt.Println("==============================")

		c.JSON(404, gin.H{
			"status":  404,
			"message": "Not Found",
			"data":    err,
		})

		return
	}

	c.JSON(200, gin.H{
		"status":  200,
		"message": "Success",
		"data":    jadwal,
	})
}

func UpdateCalendar(c *gin.Context) {

	database, err := db.Setup()
	id := c.Param("id")

	if err != nil {

		fmt.Println("===\nerror detail: ", err)

		c.JSON(500, gin.H{
			"status":  500,
			"message": "Internal Server Error",
			"data":    "database connection problem",
		})

		return
	}

	repo := calendar.SetupRepository(database)
	service := calendar.SetupService(repo)

	var json model.CalendarUpdate

	err = c.ShouldBindJSON(&json)

	if err != nil {

		fmt.Println("===\nerror detail: ", err)

		c.JSON(400, gin.H{
			"status":  400,
			"message": "Bad Request",
			"data":    "incomplete data request",
		})

		return
	}

	isUpdated, err := service.Update(id, json)

	if err != nil {

		fmt.Println("===\nerror detail: ", err)

		c.JSON(400, gin.H{
			"status":  400,
			"message": "Bad Request",
			"data":    "incomplete data request",
		})

		return
	}

	c.JSON(200, gin.H{
		"status":  200,
		"message": "success",
		"data":    isUpdated,
	})
}

func DeleteCalendar(c *gin.Context) {

	database, err := db.Setup()
	dataType := c.DefaultQuery("type", "date")
	id := c.Param("id")

	if err != nil {

		fmt.Println("===\nerror detail: ", err)

		c.JSON(500, gin.H{
			"status":  500,
			"message": "Internal Server Error",
			"data":    "database connection problem",
		})

		return
	}

	repo := calendar.SetupRepository(database)
	service := calendar.SetupService(repo)

	isDeleted, err := service.Delete(dataType, id)

	if err != nil {

		fmt.Println("===\nerror detail: ", err)

		c.JSON(500, gin.H{
			"status":  500,
			"message": "Internal Server Error",
			"data":    "database connection problem",
		})

		return
	}

	c.JSON(200, gin.H{
		"status":  200,
		"message": "Success",
		"data":    isDeleted,
	})

}
