package database

const (
	InsertUser = `INSERT INTO user (phone, name, role, password) 
                        VALUES (?, ?, ?, ?)`
	GetUser = `SELECT phone, name, role, password 
                        FROM user WHERE phone = ?`
)
