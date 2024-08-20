# Celeritas skeleton application

## Usage

### Build celeritas cli
```shell
git clone https://github.com/john-wraa/celeritas
cd celeritas
make build
cp dist/celeritas.exe ../.
cd ..
``` 

### Create new app
```shell
./celeritas.exe new testapp
cd testapp
``` 

### Run docker images
```shell
docker compose up -d
``` 

### Trying things out 1
```shell
make build
make start
curl http://localhost:4000`
``` 

### Populate variables for database and cache 
- Populate `./config/database.yml`
- Populate sections in `.env`
    - `database_*`
    - `cache` 

### Trying things out 2
```shell
make restart
../celeritas.exe help
../celeritas.exe down
curl http://localhost:4000`
../celeritas.exe up
curl http://localhost:4000`
../celeritas.exe make auth
``` 

### Replace some old imports as a workaround for a bug
- Replace `myapp/data` with `testapp4/data` in imports of `handlers/auth-handlers.go` and `middleware/remember.go`
- Try `go mod tidy`, if errors try one or both of the following
  - if error is `ambiguous import: found package cloud.google.com/go/compute/metadata in multiple modules:`
    - then run `go get cloud.google.com/go/compute/metadata` and then `go mod tidy` again
  - else close all but one goland process and run `go clean -modcache` and then `go mod tidy`
- Add `RememberToken`, `Users` and `Tokens` in `Models` struct in `./data/models.go`
- Add auth routes in `routes.go`
- Run `make build`
- Run `make restart`
- Check `http://localhost:4000/users/login`
  - Note, user table is empty and neither `GitHub` nor `Google` variables are set in `.env` file so actual login won't work
