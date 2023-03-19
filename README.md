# Short Description API

This repository represents a simple API to get people's data from those whose Wikipedia page exists in English.
## How to run this project
###  Requirements
To run this project, you must have installed Docker and Docker Compose.

###  Run command

> docker-compose up -d

### API URL
> :warning: **Warning**: This API URL only responds to HTTP **GET** method.

The api URL is http://localhost:10080/api/v1/person/<name> where token `<name>` should be substituted by the interested person’s name.


## API output schema

### in case of person exists
You will receive a http status **200** and you will have an output like the following:

```json
{
  "person": {
    "name": "",
    "short_description":"", 
            ""
  }
}
```

