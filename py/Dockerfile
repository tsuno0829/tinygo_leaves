FROM python:3.11

WORKDIR /app

COPY ./py /app

RUN pip install -r ./requirements.txt

CMD ["python", "py/main.py"]
