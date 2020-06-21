# numbers-app

The numbers-app is application that show the result from numbersapi.com trivia. Made using go and react.

![Alt text](image1.png?raw=true "image1")
![Alt text](image2.png?raw=true "image1")

## Specification

### Backend

- golang
- echo-framework

### Frontend

- react
- axios

### Deployment

- skaffold
- helm
- distroless images
- precommit

## How to use

To start the application you can use this command :

```shell
skaffold dev --profile { environment }
```

It will start the application and show all the logs including build, deployment, app logs, etc.

## Environment

There was 3 environment available.

- development
- staging
- production

## How to use

The development environment has only 1 pod will be deployed. To deploy the development profile you can use this command :

```shell
skaffold dev --profile development
```

The staging environment has 2 replicas active. To deploy the staging profile you can use this command :

```shell
skaffold dev --profile staging
```

The production environment has minimum of 2 replicas and maximum of 4 replicas active, with autoscale enabled and averageUtilization configured to 50% of cpu usage. To deploy the production profile you can use this command :

```shell
skaffold dev --profile production
```

## About skaffold

You can modify the values of replicas, cpu averageUtilization, create new profile, rename the profile in a single file skaffold.yaml.

While skaffold dev running, you can modify the code or manifest and the skaffold will instantly re-run the cycle.

And when you stop the process, it will automatically cleanup everything.

## Backend API

Test the API using this command :

```shell
curl -s -XPOST -u kita:bisa localhost:1323/numbersapi/23 | jq .
```

### API responses

200 (OK)

```json
"{\n \"text\": \"23 is the number of times Julius Caesar was stabbed.\",\n \"number\": 23,\n \"found\": true,\n \"type\": \"trivia\"\n}"
```

401 (Unauthorized)

```json
{"message":"Unauthorized"}
```

404 (Not fFund)

```json
{"message":"Not Found"}
```

405 (Method Not Allowed)

```json
{"message":"Method Not Allowed"}
```

## Source

- <https://skaffold.dev/>
- <https://github.com/GoogleContainerTools/distroless>
- <https://www.npmjs.com/package/axios>
- <https://echo.labstack.com/>
- <https://helm.sh/>
- <https://pre-commit.com/>
