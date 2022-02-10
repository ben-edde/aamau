# aamau bakery

front-end: nuxt

back-end: gin

db: mysql

## workflow

* user submit order info
* backend check db to determine whether accepting the order
* a receipt is returned to user through file/email

## DB schema

```
Table Ingredient
id, name, weight, amount

Table Cake
id, name, timeNeeded, price

Table Recipe
id, cakeId, igId, intAmtRequired

Table Order
id, orderDate, deyDate, uId, cakeId, amount, totalPrice

Table User
id, name, contact, email, address
```

## Front-end features

* displaying cake info (ingredient, price, time needed for production)
* a form receiving order info (uname, contact, email, address, cake, amount)
* (valid order) return a preview of receipt with delivery date and total price
* (valid order) return final receipt to user through file (TBC)
* (invalid order) acknowledge user with reasons

## Back-end features

* provides cake info to front-end
* receives order info and determine is that a valid order
* (valid order) return preview of receipt
* (valid order) write order to DB and return final receipt (with order number)
* (TBC) send email to user to notify delivery