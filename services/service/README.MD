# AUTHENTICATION

## REGISTER
### BODY REGISTER
- METHOD: POST
- URL ``` http://localhost:9000/api/v1/auth/register ```
```bash
{
    "username": "username", <string>
    "password": "password" <string>
}
```
### RESPONSE
```bash
{
    "_id": "66e2dfbf44cf08927a58c14f", <string>
    "username": "username", <string>
}
```
## LOGIN
### BODY LOGIN
- METHOD: POST
- URL ``` http://localhost:9000/api/v1/auth/login ```
```bash
{
    "username": "username", <string>
    "password": "password" <string>
}
```
### RESPONSE
```bash
{
    "token": "token", <string>
    "status": "success" <string>
}
```
### EXAMPLE DECODE TOKEN
```bash
{
  "_id": "66e1208887efe36606197967", <string>
  "username": "test", <string>
  "exp": 1726173052 <float64>
}
```