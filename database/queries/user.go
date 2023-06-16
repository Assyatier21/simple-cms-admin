package database

const (
	INSERT_USER = `INSERT INTO user (phone, name, role, password) 
                        VALUES (?, ?, ?, ?)`
	GET_USER = `SELECT phone, name, role, password 
                        FROM user WHERE phone = ?`
)
