#!/bin/bash

exec gunicorn --bind ${HOST:-0.0.0.0}:${PORT:-8080} app:app --timeout 90
