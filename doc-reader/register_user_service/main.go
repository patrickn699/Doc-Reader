package main

import (
    "database/sql"
    "fmt"
    "log"

    _ "github.com/lib/pq" // PostgreSQL driver
	"encoding/json"
	
    "net/http"
)

// Database connection function
func getDBConnection() (*sql.DB, error) {
    connStr := "host=localhost port=5432 user=admin password=admindoc dbname=doc_users sslmode=disable"
	return sql.Open("postgres", connStr)
}

// Register user function
func registerUser(username, email, password string) (int, error) {
    db, err := getDBConnection()
    if err != nil {
        return 0, fmt.Errorf("error connecting to database: %v", err)
    }
    defer db.Close()

    var userID int
    query := `
        INSERT INTO users (username, email, password)
        VALUES ($1, $2, $3)
        RETURNING id
    `
    err = db.QueryRow(query, username, email, password).Scan(&userID)
    if err != nil {
        return 0, fmt.Errorf("error inserting user: %v", err)
    }

    return userID, nil
}

func validateUser(username, password string) (bool, error) {
	db, err := getDBConnection()
	if err != nil {
		return false, fmt.Errorf("error connecting to database: %v", err)
	}
	defer db.Close()

	var exists bool
	query := `
		SELECT EXISTS (
			SELECT 1 FROM users WHERE username = $1 and password = $2
		)
	`
	err = db.QueryRow(query, username, password).Scan(&exists)
	if err != nil {
		return false, fmt.Errorf("error checking user existence: %v", err)
	}

	return exists, nil
}

// HTTP handler for registering a user
func registerUserHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodPost {
        http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
        return
    }

    // Parse JSON request body
    var reqBody struct {
        Username string `json:"username"`
        Email    string `json:"email"`
        Password string `json:"password"`
    }
    err := json.NewDecoder(r.Body).Decode(&reqBody)
    if err != nil {
        http.Error(w, "Invalid request body", http.StatusBadRequest)
        return
    }

    // Call the registerUser function
    userID, err := registerUser(reqBody.Username, reqBody.Email, reqBody.Password)
    if err != nil {
        http.Error(w, fmt.Sprintf("Failed to register user: %v", err), http.StatusInternalServerError)
        return
    }

    // Respond with success
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(map[string]interface{}{
        "message": "User registered successfully",
        "user_id": userID,
    })
}

func validateUserHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Invalid request method", http.StatusMethodNotAllowed)
		return
	}

	// Parse JSON request body
	var reqBody struct {
		Username string `json:"username"`
		Password string `json:"password"` // Password is not used in validation, but included for completeness
		
	}
	err := json.NewDecoder(r.Body).Decode(&reqBody)
	if err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	exists, err := validateUser(reqBody.Username, reqBody.Password)
	if err != nil {
		http.Error(w, fmt.Sprintf("Error validating user: %v", err), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]bool{
		"exists": exists,
	})
}

func main() {
    http.HandleFunc("/register", registerUserHandler)
	http.HandleFunc("/validate", validateUserHandler)
    fmt.Println("Go service running on port 8080")
    log.Fatal(http.ListenAndServe(":8080", nil))
}