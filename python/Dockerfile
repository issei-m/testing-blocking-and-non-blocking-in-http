FROM debian:stable-slim
WORKDIR /opt/app
RUN <<EOF
set -e
apt-get update
apt-get install -y --no-install-recommends pipx procps
EOF
ENV PATH=/root/.local/bin:$PATH
RUN pipx install poetry
COPY ./pyproject.toml ./poetry.lock ./
RUN poetry install --no-root
COPY . .
RUN poetry install
EXPOSE 8000
CMD ["poetry", "run", "uvicorn", "app:app", "--host", "0.0.0.0", "--port", "8000"]
