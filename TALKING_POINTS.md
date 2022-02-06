# Talking Points

Just some notes of the service and things to talk about potentially in the future.

### Things to improve upon
1. pagination: 
    returning if the next page has values 
2. sql builder: 
    allowing for easier building of sql queries and adding more complexity around them e.g 
    for building out the list api - allow the user to set what to filter by, etc
3. gzipping:
    there are a couple of easy packages that we can wrap around the http.Handler to allow us to 
    gzip the response data
4.  modulising internal services:
    as the service grows it may be worth modulising services under app - allow us to control what is
    talking to what in the db. Does Ratings need to talk to system?. Better control on communications 
    down the chain.
5. Tracing and metrics:
    We have split out the key functionality into two core packages app and store. We use interfaces to 
    tell us what methods are in this package. This allows us to easily write tracing and metric wrappers
    around them and collect this data.
6. Auth:
    In the ServeHTTP we have a nice todo to say we can add auth here. We can add more handler wraps to add more
    complexity to auth. E.g requireSession, requireToken, requireApiKey - all depending on what kind of auth 
    we want to use.
7. Logging: 
    A very basic package was setup. Should create a more robust package that writes to a file. Also create a logger wrapper
    that creates a new entry for each request and specific data attached to it which can be passed through out the application.
