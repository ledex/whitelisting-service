# Whitelisting-Service
_This is still pretty much work in progress, please keep that in mind._

## About
This service can be used to whitelist user on a Minecraft-server via an Rest-API.

## Endpoints

### General parameters:
| Name | Type | Details |
|--------|--------------------|------------------------------------------------------------------------------------------------------|
| apiKey | query-string-param | The 'apiKey' is used to authenticate any user of the API. Therefore it is required at all endpoints. |

### GET `localhost:3000/members`  
  _Lists all whitelisted members of the server._
### POST `localhost:3000/members`  
  _Adds a new member to the whitelist of the server_.  
  **Parameters:**  
  | Name | Type | Details |
|------|------------------------|--------------------------------------------------------------------------|
| username | request-body-parameter | The 'user' is the user that will be added to the whitelist of the server. The request should look like this:<br>```{"username":"foo"}```|
### DELETE `localhost:3000/members/<identifier>`  
  _Deletes the given name or uuid from the whitelist of the server._  
  **Parameters:**
  | Name | Type | Details |
  |------------|------------|-------------------------------------------------------------------------------------------------------------------------------------------------------------------------|
  | identifier | path-param | The 'identifier' is used to identify the member that should be removed from the whitelist. It can either be the current username or the UUID of the user (recommended). |

## Common return identifier
* **I001**: User is already whitelisted
* **I002**: User has been removed from the whitelist
* **E001**: Error writing/reading the whitelist
* **E002**: Username is invalid
* **E-1**: Unknown error -> Check details
  
## Todos
- [x] Return errors as Jsons
- [x] Unify id and uuid
- [ ] Add cli-parameters for port and whitelist-file
- [ ] Add some common-practises
