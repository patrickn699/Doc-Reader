from flask import Flask, render_template, request, redirect, url_for

app = Flask(__name__)

# Redirect root to login
@app.route("/")
def index():
    return redirect(url_for("login"))

# Login route
@app.route("/login", methods=["GET", "POST"])
def login():
    if request.method == "POST":
        username = request.form["username"]
        password = request.form["password"]
        if username == "admin" and password == "password":
            return redirect(url_for("home"))
        return "Invalid credentials", 401
    return render_template("login.html")

# Home route
@app.route("/home")
def home():
    return render_template("home.html")

# # Upload route
# @app.route("/upload", methods=["GET", "POST"])
# def upload():
#     if request.method == "POST":
#         file = request.files["file"]
#         if file:
#             # Here you would handle the file upload
#             return redirect(url_for("home"))
#     return render_template("upload.html")

# Register route
@app.route("/register", methods=["GET", "POST"])
def register():
    if request.method == "POST":
        username = request.form["username"]
        password = request.form["password"]
        # Here you would handle user registration logic
        return redirect(url_for("login"))
    return render_template("register.html")

if __name__ == "__main__":
    app.run(debug=True)
