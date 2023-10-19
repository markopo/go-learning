# BOOKINGS

### run:
``` go run cmd/web/*.go ```

### go version:
    1.17

### dependencies:
    github.com/alexedwards/scs/v2 v2.5.0 // indirect
    github.com/go-chi/chi/v5 v5.0.7 // indirect
    github.com/justinas/nosurf v1.1.1 // indirect
    github.com/jackc/pgx/v4

### migrations
    - soda migrations - https://gobuffalo.io/documentation/database/pop/

    soda migrate
    
    i.e.  soda generate fizz CreateFkForRoomRestrictions    

    soda reset 