FROM python:3.9-slim

WORKDIR /python-claculator

COPY . .

RUN pip install --upgrade pip && pip install -r requirements.txt

CMD ["python", "app.py"]
