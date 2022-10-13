import flask
import requests

import os

app = flask.Flask(os.environ.get("APP", __name__))
app.config["JSON_SORT_KEYS"] = False
app.config["APP_URLS"] = []

if os.environ.get("APP1_URL"):
	app.config["APP_URLS"].append(os.environ.get("APP1_URL"))

if os.environ.get("APP2_URL"):
	app.config["APP_URLS"].append(os.environ.get("APP2_URL"))

if os.environ.get("APP3_URL"):
	app.config["APP_URLS"].append(os.environ.get("APP3_URL"))

@app.route("/")
def hello():
	apps = []
	for app_url in app.config["APP_URLS"]:
		apps.append(requests.get(app_url).json())
	return flask.render_template("index.html", apps=apps)
