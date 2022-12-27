package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func getMedicinesHandler(w http.ResponseWriter, _ *http.Request) {
	medicines, err := getMedicines()

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	medicineListBytes, err := json.Marshal(medicines)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	_, err = w.Write(medicineListBytes)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
}

func createMedicineHandler(w http.ResponseWriter, r *http.Request) {
	dec := json.NewDecoder(r.Body)
	dec.DisallowUnknownFields()

	var medicine Medicine
	err := dec.Decode(&medicine)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}

	newMedicine, err := createMedicine(medicine)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	newMedicineEncoded, err := json.Marshal(newMedicine)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
	_, err = w.Write(newMedicineEncoded)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Println(err)
		return
	}
}
