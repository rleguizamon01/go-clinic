package main

import (
	"fmt"
)

func getMedicines() ([]Medicine, error) {
	var medicines []Medicine
	rows, err := db.Query("select * from medicines")
	if err != nil {
		return nil, fmt.Errorf("getMedicines query error: %v", err)
	}

	defer rows.Close()

	for rows.Next() {
		var medicine Medicine
		err := rows.Scan(&medicine.ID, &medicine.Name, &medicine.Description)
		if err != nil {
			return nil, fmt.Errorf("getMedicine foreach error: %v", err)
		}
		medicines = append(medicines, medicine)
	}

	err = rows.Err()
	if err != nil {
		return nil, fmt.Errorf("getMedicine rows error: %v", err)
	}

	return medicines, nil
}

func getMedicine(id int) (Medicine, error) {
	row := db.QueryRow("select * from medicines where id = ?", id)

	var medicine Medicine
	err := row.Scan(&medicine.ID, &medicine.Name, &medicine.Description)
	if err != nil {
		return medicine, err
	}

	return medicine, nil
}

func createMedicine(medicine Medicine) (Medicine, error) {
	res, err := db.Exec(
		"insert into medicines (name, description) values(?, ?)",
		medicine.Name,
		medicine.Description,
	)

	if err != nil {
		fmt.Println(err)
		return medicine, err
	}

	id, err := res.LastInsertId()

	if err != nil {
		fmt.Println(err)
		return medicine, err
	}

	medicine.ID = int(id)
	return medicine, nil
}
