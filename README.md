
## FizzBuzz implementation in Rest Calls

Use *make help* for listing all available commands.    

# Launch server on localhost:8080
> make runServer -serverPort=8080  
without make :  
> go run ./ -port 8080  
you can hen use http://localhost:8080/fizzbuzz for testing fizzbuzz
and http://localhost:8080/fizzbuzz/stat for displaying the more used entry.

# generate binary fizzbuzz on local folder
> make build  
the executable fizzbuzz could be launch with option -port XXXX  
# REST API
/fizzbuzz   
    POST  
        input
            content-type : application/x-www-form-urlencoded  
            args  
                int1 integer  
                int2 integer  
                limit integer  
                string1 string  
                string2 string  
        output
            status :
                200 : OK
            content-type : text/plain; charset=utf-8  
            body:
                fizzbuzz result 
            
/fizzbuzz/stat  
    GET  
        output  
            content-type : application/json; charset=utf-8  
            json : 
                {
                "Entry":{
                        "Int1":int ,
                        "Int2":int,
                        "Limit":int,
                        "String1":string,
                        "String2":string},
                "NbCalls":int}


# Run unit tests
make runTests

# Docker Image
build the image zlounes/fizzbuzz:1.0 : make buildImage
run integrations tests on the image : make runIT
env required : 
 - SERVER_PORT int

