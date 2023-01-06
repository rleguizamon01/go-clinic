package medicines

import (
	"common"
	"gorm.io/gorm"
	"log"
	"time"
)

type Medicine struct {
	ID          uint           `gorm:"primarykey" json:"id" form:"id"`
	Name        string         `json:"name" form:"name"`
	Description string         `json:"description" form:"description"`
	CreatedAt   *time.Time     `json:"created_at" form:"created_at"`
	UpdatedAt   *time.Time     `json:"updated_at" form:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index" json:"deleted_at" form:"deleted_at"`
}

func (medicine *Medicine) AfterFind(tx *gorm.DB) (err error) {
	err = common.GetWebSocket().Broadcast([]byte("Se listó el medicamento: " + medicine.Name))

	if err != nil {
		log.Panic(err.Error())
	}

	return nil
}

func (medicine *Medicine) AfterCreate(tx *gorm.DB) (err error) {
	err = common.GetWebSocket().Broadcast([]byte("Se creó el medicamento: " + medicine.Name))

	if err != nil {
		log.Panic(err.Error())
	}

	return nil
}

func (medicine *Medicine) AfterUpdate(tx *gorm.DB) (err error) {
	err = common.GetWebSocket().Broadcast([]byte("Se actualizó el medicamento: " + medicine.Name))

	if err != nil {
		log.Panic(err.Error())
	}

	return nil
}

func (medicine *Medicine) AfterDelete(tx *gorm.DB) (err error) {
	err = common.GetWebSocket().Broadcast([]byte("Se eliminó el medicamento: " + medicine.Name))

	if err != nil {
		log.Panic(err.Error())
	}

	return nil
}
