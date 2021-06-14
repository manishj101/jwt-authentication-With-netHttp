# jwt-authentication-With-netHttp
using golang, this is the simple signin and signup project with jwt-authentication

in this project we are generating, JWT token and validating it in further APIs.

we just using net/http package for making API server

**We can call the API from command line as below**

  **Call Post API with test-json.json payload**

	  . curl -d @test-json.json -H "Content-Type: application/json"  http://localhost:8000/getToken
	
	  . curl -d '{"username":"manish","password":"password","email":"maniya0412@gmail.com"}' -H 'Content-Type: application/json' http://localhost:8000/getToken
	
  **Call Get API without headers**

	  . curl -v http://localhost:8000/check 

  **Call Get with header**

	  . curl -v -H "Authorization: bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJ1c2VybmFtZSI6Im1hbmlzaCIsInVzZXJ0eXBlIjoiYWRtaW4iLCJlbWFpbCI6Im1hbml5YTA0MTJAZ21haWwuY29tIiwiZXhwIjoxNjIzNjQwNDc5fQ.nn0qw5SzIftuEJKw3wT1mM0ZpquO4pqlrGgfwjF-yWc" http://localhost:8000/check




