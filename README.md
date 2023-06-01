# LinkUp

This project is a website that I am using to improve my skills. Currently, it uses the following technologies:

* Go
* PostgreSQL
* Docker / Compose
The list of technologies will only grow in the future.
The project supports:

* Cookies (they are encrypted)
* Middlewares (the structure of the middlewares supports expandability, and the middleware chain can be customized for specific needs. Currently, there are such middlewares as cookie verification, CORS substitution, Access verification (the id from the url is checked when going to another user's page, if the id matches the account of the user who is going, then he turns out to be on his page with full access, and if it is someone else's page, then read-only rights are set, with access to view user data and friends), and HTTP logging.)
* Registration (passwords are encrypted) url (/register)
* Authorization (/auth)
* Logout (/logout)
* Logging
* It is possible to view and edit user data (/home/data), search for users (/search), add them as friends (after searching on the /search page), and see them displayed on the /friends page. It is also possible to remove friends on the /friends page.
All endpoints except /register and /auth are blocked until the user logs in.

Current endpoints:
Main:
* /register
* /auth
* /home
* /home/id:
* /home/data
* /home/data/id:
* /friends
* /friends/id:
* /search

Additional:
* /friends/add
* /friends/delete

The server is started via the main file in the root of the project. The env file is required, it must contain:

Setting server
* HTTP_PORT=7878
* HTTP_HOST=127.0.0.1

Start logging or not (Bool)
* LOGGING_API=true

Outputs http request and response data
* LOGGING_HTTP=true

Secret Key (UUID) and name for cookies
* SECRET_KEY_1=9643560ceefe93fcfd5cb3a7de8a979dd04b5255ddd29e5b4d8033a62818d4ed
* SECRET_KEY_2=fc3e6701ec93ceb4e3de750568ef6438fe7891831a09b8aad3f273f6c8c89e5b
* COOKIE_NAME=name

DB Connection
* CONNECT_DB=postgres - Name database
* DATABASE_NAME=newRestApi
* DATABASE_PORT=5432
* DATABASE_HOST=127.0.0.1
* DATABASE_USERNAME=name
* DATABASE_PASSWORD=pass
* SSLMODE=disable

Time to Read and Write the server is Seconds
* READ_TIMEOUT=3
* WRITE_TIMEOUT=3
