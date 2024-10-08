package repository

import (
	"time"

	"github.com/nihalnclt/bookings-go/internal/models"
)

type DatabaseRepo interface {
	Authenticate(email, testPassword string) (int, string, error)

	SearchAvailabilityForAllRooms(start, end time.Time) ([]models.Room, error)
	SearchAvailabilityByDatesByRoomID(start, end time.Time, roomID int) (bool, error)
	GetRoomById(id int) (models.Room, error)
	InsertReservation(res models.Reservation) (int, error)
	InsertRoomRestriction(res models.RoomRestriction) error
	AllNewReservations() ([]models.Reservation, error)
	AllReservations() ([]models.Reservation, error)
	AllRooms() ([]models.Room, error)
	GetRestrictionsForRoomByDate(roomID int, start, end time.Time) ([]models.RoomRestriction, error)
	InsertBlockForRoom(id int, startDate time.Time) error
	DeleteBlockByID(id int) error
}
