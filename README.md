<h1>๐งน Clean Architecture Project</h1>
<h3>๐ง๐ผโ๐ป Established by</h3>
<ul>
    <li><strong>Ruangyot Nanchiang</strong></li>
</ul>

<h3>๐ Introduction</h3>
<ul>
    <li>This project is to build a REST API with Golang and GO Fiber using the clean architecture of Uncle Bob.</li>
</ul>

<h3>๐ Requirements</h3>
<ul>
    <li>โก GO v1.18+</li>
    <li>๐ PostgreSQL</li>
</ul>

<h3>Install the postgreSQL on Docker</h3>
<ul>
<li>

<strong>Pull Image -> <a href="https://hub.docker.com/_/postgres" target="_blank">PostgreSQL Docker Image</a></strong>

```
docker pull postgres:alpine
```
</li>

<li>

<strong>Run the container</strong>

```
docker run --name clean-arch-db -e POSTGRES_PASSWORD=123456 -p 1122:5432 -d postgres:alpine
```
</li>
<li>

<strong>Config the postgres</strong>

```
docker exec -it clean-arch-db bash
```
```
psql -U postgres
```
```
create database clean_arch_db;
```

```
\c clean_arch_db;
```

```
CREATE TABLE "users"(
    "id" SERIAL PRIMARY KEY,
    "username" VARCHAR(255) UNIQUE NOT NULL,
    "password" VARCHAR(255) NOT NULL
);
```
</li>
<li>

<strong>Check the database that created or not</strong>

```
\l
```
```
      Name      |  Owner   | Encoding |  Collate   |   Ctype    |   Access privileges
----------------+----------+----------+------------+------------+-----------------------
 clean_arch_db | postgres | UTF8     | en_US.utf8 | en_US.utf8 |
```
</li>
</ul>

<h3>๐ฉ Initialize The Project (Just an example)</h3>
<ul>

```zsh
๐ app/
โโ ๐ main.go
๐ assets/
โโ ๐ logs/
๐ configs/
โโ ๐ configs.go
๐ modules/
โโ ๐ servers/
โ  โโ ๐ server.go
โ  โโ ๐ handler.go
โโ ๐ entities/
โ  โโ ๐ users.go
โ  โโ ๐ response.go
โโ ๐ users/
โ  โโ ๐ controllers/
โ  โ  โโ ๐ users_controllers.go
โ  โโ ๐ usecases/
โ  โ  โโ ๐ users_usecases.go
โ  โโ ๐ repositories/
โ  โ  โโ ๐ users_repositories.go
๐ pkg/
โโ ๐ databases/
โ  โโ ๐ migrations/
โ  โโ ๐ postgresql.go
โโ ๐ middlewares/
โโ ๐ utils/
โ  โโ๐ connection_url_builder.go
๐ tests/
โโ ๐ users/
โ  โโ ๐ users_test.go
๐ .env
๐ tests/
โโ ๐ users/
โ  โโ ๐ users_test.go
๐ .env
```
</ul>

<h3>๐ Start a Project</h3>
<ul>

```zsh
cd app/
go build main.go&&./main.exe
```
</ul>