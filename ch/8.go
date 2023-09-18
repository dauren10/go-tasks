package main

import "errors"

func main() {

}

func performDatabaseOperation() error {
	tx := db.Begin()
	defer func() {
		if err != nil {
			tx.Rollback()
		} else {
			tx.Commit()
		}
	}()

	// Операции с базой данных

	return nil
}

func divide(a, b float64) (result float64, err error) {
	if b == 0 {
		err = errors.New("division by zero")
		return
	}
	result = a / b
	return
}

type MyError struct {
	Message string
}

func (e *MyError) Error() string {
	return e.Message
}
