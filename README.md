### About
parser_sample is a sample to task solution that has to scrap company info from rusprofile.ru and provide swagger documentation. 

### How to run

`make run-srv-linux`


### Resources

GET `/gRPCServer.Parser/GetCompany` - returns json formatted company information:

`{
    "name": "",
    "ceo": "",
    "inn": "",
    "kpp": "",
    "error": ""
}`

### Swagger
Swagger gui is accesible on `/swaggerui/`
