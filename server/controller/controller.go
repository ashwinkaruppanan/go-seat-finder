package controller

import (
	"context"
	"net/http"

	"example.com/go-seat-finder/db"
	"example.com/go-seat-finder/model"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type departments struct {
	CA []string `json:"CA"`
	CS []string `json:"CS"`
	IT []string `json:"IT"`
	DS []string `json:"DS"`
}

var departmentsCollection = db.OpenCollection(db.Client, "departments")
var allotmentCollection = db.OpenCollection(db.Client, "allotment")

func Allot() gin.HandlerFunc {
	return func(c *gin.Context) {
		var allotDetails model.AllotRequest

		// exam_halls := []int{602, 604, 703, 704, 804, 805, 902, 903}

		if err := c.ShouldBindJSON(&allotDetails); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}

		var record departments

		id := "646b212afb8c1d396d836fb2"
		_id, err := primitive.ObjectIDFromHex(id)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
			return
		}

		departmentsCollection.FindOne(context.TODO(), bson.M{"_id": _id}).Decode(&record)

		var total_students []string
		count := 0

		for _, dep := range allotDetails.Departments {
			switch dep {
			case "CA":
				count += len(record.CA)
				total_students = append(total_students, record.CA...)
			case "CS":
				count += len(record.CS)
				total_students = append(total_students, record.CS...)
			case "DS":
				count += len(record.DS)
				total_students = append(total_students, record.DS...)
			case "IT":
				count += len(record.IT)
				total_students = append(total_students, record.IT...)
			}
		}

		if len(allotDetails.Rooms)*30 < count {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "insufficient rooms"})
			return
		}

		allotSeat := []any{}

		room_flag := 0
		i := 1
		for _, stu := range total_students {

			allotSeat = append(allotSeat, model.AllotSeat{
				ID:         primitive.NewObjectID(),
				RollNumber: stu,
				RoomNumber: allotDetails.Rooms[room_flag],
				SeatNumber: i,
				Date:       allotDetails.Date,
			})

			if i < 30 {
				i++
			} else {
				i = 1
				room_flag += 1
			}
		}

		allotmentCollection.InsertMany(context.TODO(), allotSeat)

		c.JSON(http.StatusOK, gin.H{"success": "allotment successful"})

	}
}

func FindSeat() gin.HandlerFunc {
	return func(c *gin.Context) {
		var reqInformation *model.FindSeatRequest

		if err := c.ShouldBindJSON(&reqInformation); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "bad request"})
			return
		}

		var resInformation *model.AllotSeat

		err := allotmentCollection.FindOne(context.TODO(), bson.M{"roll_number": reqInformation.RollNumber}).Decode(&resInformation)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "no results found"})
			return
		}

		c.JSON(http.StatusOK, resInformation)
	}
}
