package repository

import (
	"fmt"
	"strings"

	"gorm.io/gorm"
)

// QueryCondition represents different types of query conditions
type QueryCondition struct {
	Field    string
	Operator string // "=", "!=", ">", "<", ">=", "<=", "LIKE", "IN", "NOT IN"
	Value    interface{}
}

// QueryOptions provides additional query options
type QueryOptions struct {
	Limit   int
	Offset  int
	OrderBy string
	Preload []string // For eager loading relationships
}

func CreateDB[T any](db *gorm.DB, data *T) (*T, error) {
	create := db.Create(data)
	if create.Error != nil {
		fmt.Printf("failed to create record: %v", create.Error)
		return nil, create.Error
	}

	return data, nil
}

// FindOneDB - Original function for backward compatibility
func FindOneDB[T any](db *gorm.DB, data *T, condition map[string]interface{}) (*T, error) {
	findOne := db.Where(condition).First(data)
	if findOne.Error != nil {
		fmt.Printf("failed to find record: %v", findOne.Error)
		return nil, findOne.Error
	}
	return data, nil
}

// FindOneWithConditions - Enhanced function with flexible query conditions
func FindOneWithConditions[T any](db *gorm.DB, data *T, conditions []QueryCondition, options *QueryOptions) (*T, error) {
	query := db.Model(new(T))

	// Apply preloads if specified
	if options != nil && len(options.Preload) > 0 {
		for _, preload := range options.Preload {
			query = query.Preload(preload)
		}
	}

	// Apply conditions
	for _, condition := range conditions {
		query = applyCondition(query, condition)
	}

	// Apply ordering if specified
	if options != nil && options.OrderBy != "" {
		query = query.Order(options.OrderBy)
	}

	result := query.First(data)
	if result.Error != nil {
		fmt.Printf("failed to find record: %v", result.Error)
		return nil, result.Error
	}

	return data, nil
}

// FindManyWithConditions - Find multiple records with flexible conditions
func FindManyWithConditions[T any](db *gorm.DB, data *[]T, conditions []QueryCondition, options *QueryOptions) (*[]T, error) {
	query := db.Model(new(T))

	// Apply preloads if specified
	if options != nil && len(options.Preload) > 0 {
		for _, preload := range options.Preload {
			query = query.Preload(preload)
		}
	}

	// Apply conditions
	for _, condition := range conditions {
		query = applyCondition(query, condition)
	}

	// Apply ordering if specified
	if options != nil && options.OrderBy != "" {
		query = query.Order(options.OrderBy)
	}

	// Apply limit and offset if specified
	if options != nil {
		if options.Limit > 0 {
			query = query.Limit(options.Limit)
		}
		if options.Offset > 0 {
			query = query.Offset(options.Offset)
		}
	}

	result := query.Find(data)
	if result.Error != nil {
		fmt.Printf("failed to find records: %v", result.Error)
		return nil, result.Error
	}

	return data, nil
}

// Dynamic variants (runtime-selected model or table)
// These helpers allow selecting the model/table at runtime instead of via the generic type T.

// FindOneDynamic - model is a struct (or pointer) representing the table; dest is a pointer to struct to scan the row into
func FindOneDynamic(db *gorm.DB, model any, dest any, conditions []QueryCondition, options *QueryOptions) error {
	query := db.Model(model)

	// Apply preloads if specified
	if options != nil && len(options.Preload) > 0 {
		for _, preload := range options.Preload {
			query = query.Preload(preload)
		}
	}

	// Apply conditions
	for _, condition := range conditions {
		query = applyCondition(query, condition)
	}

	// Apply ordering if specified
	if options != nil && options.OrderBy != "" {
		query = query.Order(options.OrderBy)
	}

	return query.First(dest).Error
}

// FindManyDynamic - model is a struct (or pointer) representing the table; dest is a pointer to a slice []T
func FindManyDynamic(db *gorm.DB, model any, dest any, conditions []QueryCondition, options *QueryOptions) error {
	query := db.Model(model)

	// Apply preloads if specified
	if options != nil && len(options.Preload) > 0 {
		for _, preload := range options.Preload {
			query = query.Preload(preload)
		}
	}

	// Apply conditions
	for _, condition := range conditions {
		query = applyCondition(query, condition)
	}

	// Apply ordering if specified
	if options != nil && options.OrderBy != "" {
		query = query.Order(options.OrderBy)
	}

	// Apply limit and offset if specified
	if options != nil {
		if options.Limit > 0 {
			query = query.Limit(options.Limit)
		}
		if options.Offset > 0 {
			query = query.Offset(options.Offset)
		}
	}

	return query.Find(dest).Error
}

// UpdateDynamic - update rows for a runtime model
func UpdateDynamic(db *gorm.DB, model any, conditions []QueryCondition, updates map[string]interface{}) error {
	query := db.Model(model)

	// Apply conditions
	for _, condition := range conditions {
		query = applyCondition(query, condition)
	}

	return query.Updates(updates).Error
}

// DeleteDynamic - delete rows for a runtime model
func DeleteDynamic(db *gorm.DB, model any, conditions []QueryCondition) error {
	query := db.Model(model)

	// Apply conditions
	for _, condition := range conditions {
		query = applyCondition(query, condition)
	}

	return query.Delete(model).Error
}

// Table-based variants (use a table name string)

// FindOneTable - query by table name and scan into dest
func FindOneTable(db *gorm.DB, table string, dest any, conditions []QueryCondition, options *QueryOptions) error {
	query := db.Table(table)

	// Apply preloads if specified (effective only if model relations are known)
	if options != nil && len(options.Preload) > 0 {
		for _, preload := range options.Preload {
			query = query.Preload(preload)
		}
	}

	// Apply conditions
	for _, condition := range conditions {
		query = applyCondition(query, condition)
	}

	// Apply ordering if specified
	if options != nil && options.OrderBy != "" {
		query = query.Order(options.OrderBy)
	}

	return query.First(dest).Error
}

// FindManyTable - query many rows by table name
func FindManyTable(db *gorm.DB, table string, dest any, conditions []QueryCondition, options *QueryOptions) error {
	query := db.Table(table)

	// Apply preloads if specified (effective only if model relations are known)
	if options != nil && len(options.Preload) > 0 {
		for _, preload := range options.Preload {
			query = query.Preload(preload)
		}
	}

	// Apply conditions
	for _, condition := range conditions {
		query = applyCondition(query, condition)
	}

	// Apply ordering if specified
	if options != nil && options.OrderBy != "" {
		query = query.Order(options.OrderBy)
	}

	// Apply limit and offset if specified
	if options != nil {
		if options.Limit > 0 {
			query = query.Limit(options.Limit)
		}
		if options.Offset > 0 {
			query = query.Offset(options.Offset)
		}
	}

	return query.Find(dest).Error
}

// UpdateWithConditions - Update records with flexible conditions
func UpdateWithConditions[T any](db *gorm.DB, data *T, conditions []QueryCondition, updates map[string]interface{}) error {
	query := db.Model(new(T))

	// Apply conditions
	for _, condition := range conditions {
		query = applyCondition(query, condition)
	}

	result := query.Updates(updates)
	if result.Error != nil {
		fmt.Printf("failed to update records: %v", result.Error)
		return result.Error
	}

	return nil
}

// DeleteWithConditions - Delete records with flexible conditions
func DeleteWithConditions[T any](db *gorm.DB, data *T, conditions []QueryCondition) error {
	query := db.Model(new(T))

	// Apply conditions
	for _, condition := range conditions {
		query = applyCondition(query, condition)
	}

	result := query.Delete(data)
	if result.Error != nil {
		fmt.Printf("failed to delete records: %v", result.Error)
		return result.Error
	}

	return nil
}

// Helper function to apply conditions to query
func applyCondition(query *gorm.DB, condition QueryCondition) *gorm.DB {
	switch strings.ToUpper(condition.Operator) {
	case "=", "":
		return query.Where(fmt.Sprintf("%s = ?", condition.Field), condition.Value)
	case "!=":
		return query.Where(fmt.Sprintf("%s != ?", condition.Field), condition.Value)
	case ">":
		return query.Where(fmt.Sprintf("%s > ?", condition.Field), condition.Value)
	case "<":
		return query.Where(fmt.Sprintf("%s < ?", condition.Field), condition.Value)
	case ">=":
		return query.Where(fmt.Sprintf("%s >= ?", condition.Field), condition.Value)
	case "<=":
		return query.Where(fmt.Sprintf("%s <= ?", condition.Field), condition.Value)
	case "LIKE":
		return query.Where(fmt.Sprintf("%s LIKE ?", condition.Field), condition.Value)
	case "IN":
		return query.Where(fmt.Sprintf("%s IN ?", condition.Field), condition.Value)
	case "NOT IN":
		return query.Where(fmt.Sprintf("%s NOT IN ?", condition.Field), condition.Value)
	case "IS NULL":
		return query.Where(fmt.Sprintf("%s IS NULL", condition.Field))
	case "IS NOT NULL":
		return query.Where(fmt.Sprintf("%s IS NOT NULL", condition.Field))
	default:
		// Default to equals if operator is not recognized
		return query.Where(fmt.Sprintf("%s = ?", condition.Field), condition.Value)
	}
}

// GetAllRecords returns all records from the specified table
// Example: 
// var users []User
// users, err := repository.GetAllRecords[User](db, nil)
func GetAllRecords[T any](db *gorm.DB, options *QueryOptions) ([]T, error) {
    var records []T
    _, err := FindManyWithConditions(db, &records, nil, options)
    if err != nil {
        return nil, err
    }
    return records, nil
}

// Convenience functions for common operations

// FindByEmail - Common function to find user by email
func FindByEmail[T any](db *gorm.DB, data *T, email string) (*T, error) {
	conditions := []QueryCondition{
		{Field: "email", Operator: "=", Value: email},
	}
	return FindOneWithConditions(db, data, conditions, nil)
}

// FindByID - Common function to find record by ID
func FindByID[T any](db *gorm.DB, data *T, id uint) (*T, error) {
	conditions := []QueryCondition{
		{Field: "id", Operator: "=", Value: id},
	}
	return FindOneWithConditions(db, data, conditions, nil)
}

// FindActiveRecords - Find records where status is true
func FindActiveRecords[T any](db *gorm.DB, data *[]T, options *QueryOptions) (*[]T, error) {
	conditions := []QueryCondition{
		{Field: "status", Operator: "=", Value: true},
	}
	return FindManyWithConditions(db, data, conditions, options)
}
