FROM ghcr.io/astral-sh/uv:alpine

WORKDIR /app
ENV UV_NO_DEV=1

COPY .python-version pyproject.toml uv.lock ./
RUN uv sync --locked

RUN head -c16 /dev/urandom > secret_key.bin
RUN echo "h0h0h0{0h_my_0h_my_e9e324f3}" > /flag.txt

COPY ./static ./static
COPY ./templates ./templates
COPY app.py python_adapter.py ./

CMD ["uv", "run", "flask", "--app", "app.py", "run", "--host", "0.0.0.0"]