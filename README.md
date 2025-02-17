[![genDocsGPT](https://img.shields.io/badge/Docs%20generated%20by-genDocsGPT-blue)](https://github.com/marco-rosner/genDocsGPT)

## Lightweight Go Server

A performance comparison to find the lightweight Go server using Echo, Gin and Fiber servers. We will use MongoDB and Postgres as database and Docker compose as container enviorement. We will use some of the rules from [Rinha de beckend](https://github.com/zanfranceschi/rinha-de-backend-2023-q3) to get the results.

## Run server

You should go to the server folder that you would like to use, set the env var for the choosen database and run the server. Like:

```sh
cd echo && DB="mongo" MONGODB_URI="..." go run .
```

## Run with Docker compose

Make sure that you have [docker](https://docs.docker.com/get-docker/) and [docker compose](https://docs.docker.com/compose/install/) installed. Next, you should choose which server and database you would like to run (e.g. docker-compose-<server>-<database>.yml) and run the command like this:

```sh
docker-compose -f docker-compose-echo-postgres.yml up --remove-orphans
```

## Run benchmark

I am using Gatling to run the benchmark and to install it go to benchmark folder and run the `install-gatling` script. After that, move the `BenchmarkSimulation.scala` file to `deps/gatling/user-files/simulations` and the resources `people-payloads.tsv` and `search-terms.tsv` to `deps/gatling/user-files/recources`.

Go to the command line and run:

```sh
sh deps/gatling/bin/gatling.sh
```

## Benchmark Results

```
Download Benchmark Results From Github Actions Artifacts.
```