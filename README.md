<h1>🧹 Clean Arch Medium Product</h1>
<h3>🧑🏼‍💻 Established by</h3>
<ul>
    <li><strong>Ruangyot Nanchiang</strong></li>
</ul>

<h3>📃 Introduction</h3>
<ul>
    <li>This project is to build a REST API with Golang and GO Fiber using the clean architecture of Uncle Bob.</li>
</ul>

<h3>📝 Requirements</h3>
<ul>
    <li>⚡ GO v1.18+</li>
    <li>🐘 PostgreSQL</li>
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
docker exec -it manga-store-db-<your-stage> bash
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

<h3>🔩 Initialize The Project (Just an example)</h3>
<ul>

```zsh
📂 app/
├─ 📄 main.go
📂 assets/
├─ 📂 logs/
📂 configs/
├─ 📄 configs.go
📂 modules/
├─ 📂 servers/
│  ├─ 📄 server.go
│  ├─ 📄 handler.go
├─ 📂 entities/
│  ├─ 📄 users.go
│  ├─ 📄 response.go
├─ 📂 users/
│  ├─ 📂 controllers/
│  │  ├─ 📄 users_controllers.go
│  ├─ 📂 usecases/
│  │  ├─ 📄 users_usecases.go
│  ├─ 📂 repositories/
│  │  ├─ 📄 users_repositories.go
📂 pkg/
├─ 📂 databases/
│  ├─ 📂 migrations/
│  ├─ 📄 postgresql.go
├─ 📂 middlewares/
├─ 📂 utils/
│  ├─📄 connection_url_builder.go
📂 tests/
├─ 📂 users/
│  ├─ 📄 users_test.go
📄 .env
📂 tests/
├─ 📂 users/
│  ├─ 📄 users_test.go
📄 .env
```
</ul>

<h3>🚀 Start a Project</h3>
<ul>

```zsh
cd app/
go build main.go&&./main.exe
```
</ul>