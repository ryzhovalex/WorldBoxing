# quco.go
Query Collection protocol. Go flavour.

## Examples
### Get users with name "Mark" and set their name to "Johnny"
```quco
GET
name.IN=("Mark")
THEN SET
name="Johnny"
```

### Create user
```quco
CREATE
name="Fiona"
age=25
country="Virginia"
item=(
     type="gun"
     price=105
)
```

### Delete users with age less than 18 and without permission item
```quco
GET user
age.LT=18
item.type.NE="permission"
THEN delete
```

## Known issues
* All containers (enclosed in `()`) must be placed in single line, or the interpreter will treat them as separate instructions, and thus will throw an error.
