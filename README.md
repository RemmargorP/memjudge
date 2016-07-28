**MemJudge**
========

Online Judge for Competitive Programming Contests. And Memes

----------


##**Structure**##
1. Front End:
  * JavaScript:
    * Handle user actions
    * Send requests
2. Back end:
  * Golang:
    * Handles HTTP Requests from users:
      * Looks up DB
      * Adds new Solutions to Queue
      * ...
  * Python:
    * Testing Solutions (?)
3. Database:
  1. MongoDB:
    * Users Info
    * Solutions
    * Testing Results
    * Testing Queue
    * Contests


----------


##**Database Structure**##

###MongoDB Collections###
1. users:
  * Example:
  ```json
  {
      '_id': {
        'login': ...,
        'passwordHash': ...,
        'nickname': ...,
        'email': ...,
        'name': ...,
        'surname': ...,
        'address': ..., #and other (index, ...)
      'solutions': [ids],
      'contests': [ids],
      ... #TODO (Polygon auth info)
      },
      ...
  }
  ```

2. contests:
  * Example:
  ```json
  {
    'id': {
      'owners': [Rights],
      'authors': [user ids],
      'registeredUsers': [user ids],
      'problems': [ids],
      'startDate': date,
      'endDate': date,
      'private': true/false,
      'scoring': [len(problems), ranking system],
      # TODO Ranking System
    },
    ...
  }
  ```
3. problems:
  * Example:
  ```json
  {
    'id': {
      'owners': [Rights],
      'authors': [user ids],
      'name': ...,
      'statement': ...,
      'tests': [Tests],
      'inputType': stdin/file name (if not interactive),
      'outputType': stdout/file name (if not interactive),
      'checker': checker file, # TODO specification
      'type': standard/interactive,
      
      'polygonId': ..., #TODO
    },
    ...
  
  }
  ```
4. solutions:
  * Example:
  ```json
  {
    'id': {
      'owner': user id,
      'problem': problem id,
      'testingTesult': TestingResult,
      'timestamp': date,
      
    }
  }
  ```
5. testing_queue

####Utility Classes####
```json
'Rights':
{
  'owner': user id,
  'read': true/false,
  'write': true/false,
  'setRights': true/false
}
```
```json
'Test':
{
  'problem': problem id,
  'input': ...,
  'inputShort': first 1024 symbols of input,
  'output': ...,
  'outputShort': first 1024 symbols of output,
  'timeConsumed': float,
  'memoryUsed': bytes,
  'reason': does not exist | reason of verdict;
}
# TODO add limits
```
```json
'TestingResult' :
{
  'status': (Queued/Compiling/Compilation Error/Running/(OK|Accepted)/Time Limit Exceeded/Memory Limit Exceeded/Runtime Error/Wrong Answer/Security Violation/...)
  'tests': [Test],
  'invokerId': id,
}
```

----------


##**TODO**##
### What about... ###
1. Move Testing Queue to another DB (like **MemCached**, **Redis**, etc.)
2. Move main DB from **MongoDB** to **MySQL**, ...
3. Decide which language to use for testing solutions **(Python / Golang / C++ / Java)**
4. Add **Polygon** Integration

----------
