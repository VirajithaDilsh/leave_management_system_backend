package models

import "time"

type LeaveRequest struct {
	ID         uint      `json:"id" gorm:"primaryKey"`
	CreatedAt  time.Time `json:"createdAt"`
	UpdatedAt  time.Time `json:"updatedAt"`
	EmployeeID uint      `json:"employeeId"`
	Employee   User      `json:"employee" gorm:"foreignKey:EmployeeID"`
	LeaveType  string    `json:"leaveType"`
	StartDate  string    `json:"startDate"`
	EndDate    string    `json:"endDate"`
	Reason     string    `json:"reason"`
	Status     string    `json:"status" gorm:"default:'pending'"`
}