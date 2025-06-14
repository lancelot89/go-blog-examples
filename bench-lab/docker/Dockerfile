FROM golang:1.24.3
RUN apt-get update && apt-get install -y graphviz
WORKDIR /app
COPY . .
ENTRYPOINT ["bash", "bench-lab/scripts/run_bench.sh"]
