FROM python:3.12-slim

WORKDIR /app

ENV PIP_DEFAULT_TIMEOUT=120 \
    PIP_RETRIES=10

RUN pip install --no-cache-dir fastapi uvicorn python-multipart

COPY main.py .
COPY index.html .
COPY images ./images

COPY flag.txt /flag.txt

EXPOSE 8000

CMD ["uvicorn", "main:app", "--host", "0.0.0.0", "--port", "8000"]
