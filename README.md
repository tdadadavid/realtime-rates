# Realtime Exchange Rates

## Description

This project contains two features:

* This project is for getting the exchange rates for currency pairs `{"currency-pair": "USD-GBP"}` using two different APIs and returning the first response gotten from either.
* Reading from a json file and performing 3 basic data operations on it
   1. Grouping by currency.
   2. Filtering by Salary constraint in USD
   3. Sorting according to salary.

****

## Technologies

  1. Language: Golang
  2. Framework: Fiber
  3. Secret-Manager: AWS Secrets Manager
  4. Containerization Tool: Docker

## SETUP

### API Documentaion [Postman]

```bash
  https://documenter.getpostman.com/view/28806235/2sA2rGtyA6
```

### Clone project

```git
  git clone git@github.com:tdadadavid/realtime-rates.git
```

### Install dependencies

```markdown
  go mod install 
```

### Setup via Docker

```bash
  docker run -p $PORT:3000 dockerrundavid/realtime-rates -d
```

### Create environment file & fill them correctly

```bash
  cp .env.example .env
```

### Start project [dev mode]

```bash
  make start_dev
```

### Start project [prod]

```bash
  make start_prod
```

****

### Run Tests

```bash
  make test "/path/to/test/file"
```

****

## GUIDELINES

****

* When creating *new branch* if you're fixing a bug or implementing a feature follow this pattern
  * fixing: *`fix/<name-of-fix>`*
  * feature: *`feat/<name-of-feature>`*
* Never commit directly into the *`main`* branch.
