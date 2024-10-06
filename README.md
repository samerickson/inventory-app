# Inventory App

## Running

By far the simplest way to run this application is using docker, but there are
also instructions on how to run it without docker in this section.

### Production

> ![WARNING]
> TODO: this section still needs to be completed.

### Development Server

Add a `.env` file. You can do so by copying the `.env.sample` file:

```sh
cp .env.sample .env
```
> ![NOTE]
> The sample will work as is so long as you are running with docker compose,
> but it is recommended to update these values.

**Run Docker compose:**

```sh
docker compose watch
```

#### Without Docker

Note: you will need to have access to an instance of postgres, and you will need
to accurately add a `.env` file that matches the `.env.sample` file but with the
correct details set for each of the variables.

**Back-end:**

First navigate to the back-end folder. Then walk through the following steps:

```sh
go run main.go
```

Running with [go-air](https://github.com/air-verse/air)

Install go-air:

```sh
go install github.com/air-verse/air@latest
```

Run the server with hot reloading:

```sh
air
```

**Front-end:**

First navigate to the front-end folder. Then execute the following:

```sh
npm i && npm run dev
```

## Running

You can run the application by entering:
