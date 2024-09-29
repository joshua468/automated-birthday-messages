package repository

import (
	"github.com/joshua468/automated-birthday-messenger/db"
	"github.com/joshua468/automated-birthday-messenger/models"
)

// CreateEmployee adds a new employee to the database
func CreateEmployee(employee models.Employee) error {
	return db.DB.Create(&employee).Error // Use the initialized db connection
}

// GetAllEmployees returns all employees from the database
func GetAllEmployees() ([]models.Employee, error) {
	var employees []models.Employee
	err := db.DB.Find(&employees).Error // Use the initialized db connection
	return employees, err
}
