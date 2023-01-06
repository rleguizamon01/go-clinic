package medicines

import (
	"common"
)

func GetMedicines() ([]Medicine, error) {
	db := common.GetDB()
	var medicines []Medicine

	result := db.Find(&medicines)

	if result.Error != nil {
		return nil, result.Error
	}

	return medicines, nil
}

func GetMedicine(id string) (Medicine, error) {
	db := common.GetDB()
	var medicine Medicine

	result := db.First(&medicine, id)

	if result.Error != nil {
		return medicine, result.Error
	}

	return medicine, nil
}

func CreateMedicine(medicine Medicine) (Medicine, error) {
	db := common.GetDB()

	result := db.Create(&medicine)

	if result.Error != nil {
		return medicine, result.Error
	}

	return medicine, nil
}

func UpdateMedicine(medicine Medicine) (Medicine, error) {
	db := common.GetDB()

	result := db.Model(&Medicine{ID: medicine.ID}).Updates(&medicine)

	if result.Error != nil {
		return medicine, result.Error
	}

	return medicine, nil
}

func DeleteMedicine(id string) error {
	db := common.GetDB()

	result := db.Delete(&Medicine{}, id)

	if result.Error != nil {
		return result.Error
	}

	return nil
}
