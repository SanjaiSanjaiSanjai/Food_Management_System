package repository

import (
	schema "Food_Delivery_Management/Schema"
	"gorm.io/gorm"
)

// This file contains examples of how to use the enhanced database functions

// Example 1: Simple find by email
func ExampleFindByEmail(db *gorm.DB, email string) (*schema.User, error) {
	var user schema.User
	return FindByEmail(db, &user, email)
}

// Example 2: Find user with multiple conditions
func ExampleFindUserWithConditions(db *gorm.DB, email string) (*schema.User, error) {
	var user schema.User
	
	conditions := []QueryCondition{
		{Field: "email", Operator: "=", Value: email},
		{Field: "status", Operator: "=", Value: true},
		{Field: "is_verified", Operator: "=", Value: true},
	}
	
	options := &QueryOptions{
		Preload: []string{"Role", "User_Addresses"},
	}
	
	return FindOneWithConditions(db, &user, conditions, options)
}

// Example 3: Find users with LIKE operator (search)
func ExampleSearchUsersByUsername(db *gorm.DB, searchTerm string) (*[]schema.User, error) {
	var users []schema.User
	
	conditions := []QueryCondition{
		{Field: "username", Operator: "LIKE", Value: "%" + searchTerm + "%"},
		{Field: "status", Operator: "=", Value: true},
	}
	
	options := &QueryOptions{
		Limit:   10,
		OrderBy: "created_at DESC",
	}
	
	return FindManyWithConditions(db, &users, conditions, options)
}

// Example 4: Find users created in a date range
func ExampleFindUsersByDateRange(db *gorm.DB, startDate, endDate string) (*[]schema.User, error) {
	var users []schema.User
	
	conditions := []QueryCondition{
		{Field: "created_at", Operator: ">=", Value: startDate},
		{Field: "created_at", Operator: "<=", Value: endDate},
		{Field: "status", Operator: "=", Value: true},
	}
	
	options := &QueryOptions{
		OrderBy: "created_at DESC",
		Limit:   50,
	}
	
	return FindManyWithConditions(db, &users, conditions, options)
}

// Example 5: Find users by multiple IDs
func ExampleFindUsersByIDs(db *gorm.DB, userIDs []uint) (*[]schema.User, error) {
	var users []schema.User
	
	conditions := []QueryCondition{
		{Field: "id", Operator: "IN", Value: userIDs},
		{Field: "status", Operator: "=", Value: true},
	}
	
	return FindManyWithConditions(db, &users, conditions, nil)
}

// Example 6: Update user status
func ExampleUpdateUserStatus(db *gorm.DB, userID uint, status bool) error {
	var user schema.User
	
	conditions := []QueryCondition{
		{Field: "id", Operator: "=", Value: userID},
	}
	
	updates := map[string]interface{}{
		"status": status,
	}
	
	return UpdateWithConditions(db, &user, conditions, updates)
}

// Example 7: Soft delete user (set status to false)
func ExampleSoftDeleteUser(db *gorm.DB, userID uint) error {
	var user schema.User
	
	conditions := []QueryCondition{
		{Field: "id", Operator: "=", Value: userID},
	}
	
	updates := map[string]interface{}{
		"status": false,
	}
	
	return UpdateWithConditions(db, &user, conditions, updates)
}

// Example 8: Find active users with pagination
func ExampleFindActiveUsersWithPagination(db *gorm.DB, page, limit int) (*[]schema.User, error) {
	var users []schema.User
	
	offset := (page - 1) * limit
	
	options := &QueryOptions{
		Limit:   limit,
		Offset:  offset,
		OrderBy: "created_at DESC",
		Preload: []string{"Role"},
	}
	
	return FindActiveRecords(db, &users, options)
}

// Example 9: Complex query with multiple operators
func ExampleComplexUserQuery(db *gorm.DB, minID uint, excludeEmails []string) (*[]schema.User, error) {
	var users []schema.User
	
	conditions := []QueryCondition{
		{Field: "id", Operator: ">=", Value: minID},
		{Field: "email", Operator: "NOT IN", Value: excludeEmails},
		{Field: "is_verified", Operator: "=", Value: true},
		{Field: "status", Operator: "=", Value: true},
	}
	
	options := &QueryOptions{
		OrderBy: "username ASC",
		Limit:   100,
		Preload: []string{"Role", "User_Addresses"},
	}
	
	return FindManyWithConditions(db, &users, conditions, options)
}

// Example 10: Find users with null/empty fields
func ExampleFindUsersWithMissingData(db *gorm.DB) (*[]schema.User, error) {
	var users []schema.User
	
	conditions := []QueryCondition{
		{Field: "username", Operator: "IS NULL", Value: nil},
		// OR you could use: {Field: "username", Operator: "=", Value: ""},
	}
	
	return FindManyWithConditions(db, &users, conditions, nil)
}
