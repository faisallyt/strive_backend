# Strive Go Backend

Go backend is designed in keeping microservices in consideration, hence it contains packages which can be converted to microservices as needed.

## Setup Dev Environment

Go `version 1.22.4` needs to be installed on the system of use. Run the following commands to setup this repo.

```bash
# Clone from github
git clone github.com/Strive-Gaming-Hub/strive_go
cd strive_go

# install dependencies
go mod tidy # install go deps
go install github.com/swaggo/swag/cmd/swag@latest
```

### Environment Variables

Copy the contents of `.env.example` to a new file `.env` that is to be made in the outmost dir. Or you can run the following commands in project dir.

```bash
cp .env.example .env
```

Now fill all the variables in .env according to you environment.
