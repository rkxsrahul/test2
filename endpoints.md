# Example Requests to API endpoints => #

## 1. Healthz ##

```
Get /healthz
```

## 2. Create MarkAndFetch - To update the answer table which every option is marked by student. Then its fetch the next question with option and show to the user.##

```
POST /v1/markandfetch

Header:-
Authorization: Bearer <<test_token>>
Content-Type: application/json

Body:-
index       int       
pool        string    
marked      int       
next_pool   string    
next_index  int       
```

## 3. SubmitTest - when user submit the test its update the total time taken, Number of question attempted, marked and unattemped by user and  ##

```
GET /v1/submittest

Header:-
Authorization: Bearer <<test_token>>
Content-Type: application/json
```
## 4. IncrementSwitch - when user switch the browser its update the user_sessions table that how many time user switch the browser.##

```
PUT /v1/increment-switch

Header:-
Authorization: Bearer <<test_token>>
```

## 5. To disable or enable Mail Service##

```
PUT /v1/admin/mail/:value
Header:-
Authorization: Bearer <<admin token>>

Param data
value -> true if want to diable service and false if you want to enable service
```
