package v1

import "arm_go/handler"

func calendarRoute() {

	calendar.GET("/", handler.ReadCalendar)
	calendar.GET("/:id", handler.ReadJadwal)
	calendar.POST("/", handler.CreateCalendar)
	calendar.PUT("/:id", handler.UpdateCalendar)
	calendar.DELETE("/:id", handler.DeleteCalendar)
}
