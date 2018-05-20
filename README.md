# inventory

> A Vue.js + Go project

## Web Config Setup


## Web Build Setup

``` bash
# to web directory
cd web

# install dependencies
npm install

# serve with hot reload at localhost:8080
npm run dev

# build for production with minification
npm run build

# build for production and view the bundle analyzer report
npm run build --report

# run e2e tests
npm run e2e

# run all tests
npm test
```

For a detailed explanation on how things work, check out the [guide](http://vuejs-templates.github.io/webpack/) and [docs for vue-loader](http://vuejs.github.io/vue-loader).


## Go Configuration Setup
```bash

# configuration file
in root directory create a inventory.json

{
    "app": {
        "name": "inventory",
        "host": "localhost:3000",
        "post": "3000",
        "apiBase": "/api/v1",
        "timeout": {
            "read": 60,
            "write": 60
        }
    },
    "database": {
        "connection": {
            "host": "localhost:27017",
            "username": "sa",
            "password": "p@ssw0rd",
            "timeout": 60
        },
        "main": {
            "name": "inventory"
        },
        "master": {
            "name": "master"
        },
        "demo": {
            "name": "demo"
        }
    },
    "secretKey": "secret"
}

```

## Go Build Setup

```bash
# build go project
go build


# run go prohect
.\inventory (window)
./inventory (linux)

```

TEMPLATE launch.json for VS Code
{
    "version": "0.2.0",
    "configurations": [
    {
        "name": "Launch",
        "type": "go",
        "request": "launch",
        "mode": "debug",
        "remotePath": "",
        "port": 3000,
        "host": "localhost",
        "program": "${workspaceRoot}",
        "env": {},
        "args": ["--config-file", "inventory.json"],
        "showLog": true
    }]
}