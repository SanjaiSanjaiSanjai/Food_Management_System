package enum

// create enum in db query
const ROLE string = "CREATE TYPE role_enum AS ENUM ('User', 'Deliver_Agent', 'Owner')"

// check enum exists or not query
const EXISTS string = "SELECT EXISTS (SELECT 1 FROM pg_type WHERE typname = 'role_enum')"
