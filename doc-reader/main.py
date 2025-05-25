from flask import Flask, render_template, request, redirect, url_for
from psycopg2 import sql
import requests

app = Flask(__name__)

# Redirect root to login page
@app.route("/")
def index():
    return redirect(url_for("login"))

# Login route
@app.route("/login", methods=["GET", "POST"])
def login():
    if request.method == "POST":
        username = request.form["username"]
        password = request.form["password"]
        # Here you would handle user authentication logic
        user = validate_user(username, password)
        if "error" in user:
            return render_template("login.html", error=user["error"])
        else:
            # Assuming user validation is successful
            return redirect(url_for("upload_file"))
    return render_template("login.html")


def validate_user(username, password):
    # This function would typically check the database for user credentials.
    try:
        go_service_url = "http://localhost:8080/validate_user"
        payload = {
            "username": username,
            "password": password
        }
        response = requests.post(go_service_url, json=payload)
        if response.status_code == 200:
            return response.json()  # Assuming the Go service returns JSON
        else:
            return {"error": "Invalid credentials"}
    except Exception as e:
        return {"error": str(e)}


# Register User
def register_user(username, email, password):
    try:
        # Define the Go service URL
        go_service_url = "http://localhost:8080/register"

        # Prepare the payload
        payload = {
            "username": username,
            "email": email,
            "password": password
        }

        # Make a POST request to the Go service
        response = requests.post(go_service_url, json=payload)

        # Check if the request was successful
        if response.status_code == 200:
            return response.json()  # Return the JSON response from the Go service
        else:
            # Handle errors
            return {"error": f"Failed to register user: {response.text}"}
    except Exception as e:
        return {"error": str(e)}

        

# Register route
@app.route("/register", methods=["GET", "POST"])
def register():
    if request.method == "POST":
        username = request.form["username"]
        email = request.form["email"]
        password = request.form["password"]
        confirm_password = request.form["confirm_password"]

        # Here you would handle user registration logic
        if password != confirm_password:
            return "Passwords do not match", 400
        elif register_user(username, email, password):
            return redirect(url_for("login"))

        return redirect(url_for("login"))
    return render_template("register.html")

# File upload route
@app.route("/upload", methods=["GET", "POST"])
def upload_file():
    if request.method == "POST":
        file = request.files["file"]
        if file:
            # Here you would handle file upload logic
            # For example, save the file to a specific directory
            file.save(f"./uploaded-files/{file.filename}")
            return redirect(url_for("index"))
    return render_template("index.html")

if __name__ == "__main__":
    app.run(debug=True)
