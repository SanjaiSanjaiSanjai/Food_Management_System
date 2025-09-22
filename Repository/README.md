# Database Repository Functions

This package provides enhanced database operations with flexible query building capabilities using GORM.

## Features

- **Flexible Query Conditions**: Support for various operators (=, !=, >, <, >=, <=, LIKE, IN, NOT IN, IS NULL, IS NOT NULL)
- **Query Options**: Pagination, ordering, and eager loading
- **Type Safety**: Generic functions that work with any struct type
- **Convenience Functions**: Common operations like FindByEmail, FindByID
- **Backward Compatibility**: Original functions are preserved

## Core Types

### QueryCondition
```go
type QueryCondition struct {
    Field    string      // Database field name
    Operator string      // Comparison operator
    Value    interface{} // Value to compare against
}
```

### QueryOptions
```go
type QueryOptions struct {
    Limit   int      // Maximum number of records to return
    Offset  int      // Number of records to skip
    OrderBy string   // ORDER BY clause (e.g., "created_at DESC")
    Preload []string // Relationships to eager load
}
```

## Available Functions

### Basic Operations

#### FindOneDB (Original - Backward Compatible)
```go
func FindOneDB[T any](db *gorm.DB, data *T, condition map[string]interface{}) (*T, error)
```

#### FindOneWithConditions (Enhanced)
```go
func FindOneWithConditions[T any](db *gorm.DB, data *T, conditions []QueryCondition, options *QueryOptions) (*T, error)
```

#### FindManyWithConditions
```go
func FindManyWithConditions[T any](db *gorm.DB, data *[]T, conditions []QueryCondition, options *QueryOptions) (*[]T, error)
```

#### UpdateWithConditions
```go
func UpdateWithConditions[T any](db *gorm.DB, data *T, conditions []QueryCondition, updates map[string]interface{}) error
```

#### DeleteWithConditions
```go
func DeleteWithConditions[T any](db *gorm.DB, data *T, conditions []QueryCondition) error
```

### Convenience Functions

#### FindByEmail
```go
func FindByEmail[T any](db *gorm.DB, data *T, email string) (*T, error)
```

#### FindByID
```go
func FindByID[T any](db *gorm.DB, data *T, id uint) (*T, error)
```

#### FindActiveRecords
```go
func FindActiveRecords[T any](db *gorm.DB, data *[]T, options *QueryOptions) (*[]T, error)
```

## Usage Examples

### Simple Find by Email
```go
var user schema.User
result, err := repository.FindByEmail(db.DB, &user, "user@example.com")
```

### Complex Query with Multiple Conditions
```go
var user schema.User

conditions := []repository.QueryCondition{
    {Field: "email", Operator: "=", Value: "user@example.com"},
    {Field: "status", Operator: "=", Value: true},
    {Field: "is_verified", Operator: "=", Value: true},
}

options := &repository.QueryOptions{
    Preload: []string{"Role", "User_Addresses"},
    OrderBy: "created_at DESC",
}

result, err := repository.FindOneWithConditions(db.DB, &user, conditions, options)
```

### Search with LIKE Operator
```go
var users []schema.User

conditions := []repository.QueryCondition{
    {Field: "username", Operator: "LIKE", Value: "%john%"},
    {Field: "status", Operator: "=", Value: true},
}

options := &repository.QueryOptions{
    Limit:   10,
    OrderBy: "username ASC",
}

results, err := repository.FindManyWithConditions(db.DB, &users, conditions, options)
```

### Find by Multiple IDs
```go
var users []schema.User
userIDs := []uint{1, 2, 3, 4, 5}

conditions := []repository.QueryCondition{
    {Field: "id", Operator: "IN", Value: userIDs},
}

results, err := repository.FindManyWithConditions(db.DB, &users, conditions, nil)
```

### Date Range Query
```go
var users []schema.User

conditions := []repository.QueryCondition{
    {Field: "created_at", Operator: ">=", Value: "2024-01-01"},
    {Field: "created_at", Operator: "<=", Value: "2024-12-31"},
}

options := &repository.QueryOptions{
    OrderBy: "created_at DESC",
    Limit:   50,
}

results, err := repository.FindManyWithConditions(db.DB, &users, conditions, options)
```

### Update Records
```go
var user schema.User

conditions := []repository.QueryCondition{
    {Field: "id", Operator: "=", Value: 1},
}

updates := map[string]interface{}{
    "status":      true,
    "is_verified": true,
}

err := repository.UpdateWithConditions(db.DB, &user, conditions, updates)
```

### Pagination
```go
var users []schema.User
page := 2
limit := 10
offset := (page - 1) * limit

options := &repository.QueryOptions{
    Limit:   limit,
    Offset:  offset,
    OrderBy: "created_at DESC",
}

results, err := repository.FindActiveRecords(db.DB, &users, options)
```

## Supported Operators

- `=` or `""` (empty): Equal to
- `!=`: Not equal to
- `>`: Greater than
- `<`: Less than
- `>=`: Greater than or equal to
- `<=`: Less than or equal to
- `LIKE`: Pattern matching (use % for wildcards)
- `IN`: Value in list
- `NOT IN`: Value not in list
- `IS NULL`: Field is null
- `IS NOT NULL`: Field is not null

## Error Handling

All functions return errors that should be handled appropriately:

```go
result, err := repository.FindByEmail(db.DB, &user, email)
if err != nil {
    // Handle error (user not found, database error, etc.)
    return err
}
// Use result
```

## Migration from Old Functions

The original `FindOneDB` function is still available for backward compatibility, but it's recommended to migrate to the new functions for better flexibility:

### Before:
```go
result, err := repository.FindOneDB(db.DB, &user, map[string]interface{}{
    "email": user.Email,
})
```

### After:
```go
// Simple way
result, err := repository.FindByEmail(db.DB, &user, user.Email)

// Or with more control
conditions := []repository.QueryCondition{
    {Field: "email", Operator: "=", Value: user.Email},
}
result, err := repository.FindOneWithConditions(db.DB, &user, conditions, nil)
```
