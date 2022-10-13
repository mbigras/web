FROM python:3.10.6

WORKDIR /app

COPY requirements.txt ./

RUN pip install -r requirements.txt

COPY app.py entrypoint.sh ./
COPY templates ./templates/

ENTRYPOINT ["/app/entrypoint.sh"]
