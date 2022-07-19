
# FizzBuzz implementation in Rest Calls

Use *make help* for listing all available commands.    

## Launch server on localhost:8080
> make runServer -serverPort=8080  

without make :  
> go run ./ -port 8080  

you can then use http://localhost:8080/fizzbuzz for testing fizzbuzz
and http://localhost:8080/fizzbuzz/stat for displaying the more used entry.

## generate binary fizzbuzz on local folder
> make build  

the executable fizzbuzz could be then be launched with option -port XXXX  
## REST API

```
-    /fizzbuzz   
      POST  
        input
          content-type : application/x-www-form-urlencoded  
            args  
              int1 integer > 0  
              int2 integer  > 0
              limit integer  > 0
              string1 string
              string2 string  
        output
          status :
              200 : OK
          content-type : text/plain; charset=utf-8  
          body :
            fizzbuzz result 
      GET
        output
          content-type : text/html
          body : 
            form to input fizzbuzz calculation     

-    /fizzbuzz/stat  
      GET  
        output  
          status:
            200 : OK
          content-type : application/json; charset=utf-8  
          json : 
            {
              "Entry":{
                "Int1":int ,
                "Int2":int,
                "Limit":int,
                "String1":string,
                "String2":string},
              "NbCalls":int
            }
```

## Run unit tests
>make runTests

## Docker Image
build the image zlounes/fizzbuzz:1.0 : 
>make buildImage

run integrations tests on the image : 
>make runIT

environment variable required for running image: 
 - SERVER_PORT int

