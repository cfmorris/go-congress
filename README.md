# go-congress

This is a work in progress

## Configuration

`exampleConfig.yaml` should be updated to contain your information and renamed `config.yaml`

A Propublica api key can be requested at https://www.propublica.org/datastore/api/propublica-congress-api

This application uses a neo4j community edition database https://neo4j.com/download-center/

## Retreiving Information

go-congress currently only retrieves:
- Congress Members and their congress number.

TODO:
- retrieve subcomittees and members.
- retreive bills
- retrieve votes
- retrieve statements about votes

## Presenting information

Routes served currently present no information and just echo back the endpoints.

TODO:
- present raw data for congressmembers to an endpoint
- decide on a frontend framework to present that data.

## Version Information

go      1.18.5
neo4j   4.4.7
