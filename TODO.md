# Matillion Go Tech Test

Build a ratings web service that supports the retrieval and creation of ratings for Star Wars films.

## Requirements

The service should be built to the best of your ability and should mimic a production facing application (e.g. include adequate tests, logging, standardised HTTP status codes, documentation etc).

Film data from https://swapi.dev/api/films should be used in combination with your own data to build a simple ratings system.

Any supporting services should use Docker where/if appropriate.

The web service should include the ability to:
- Get a list of films that includes the Title, Episode ID, Director, Producer and Release Date.

- Create and store ratings for a given film. A rating should contain an Author, Score/Rating out of 5 and Created At. 

- Get a list of ratings for a given film that includes the Title, Score/Rating, Author, Created At. There should be an optional Max Score/Rating filter that can be applied (e.g. Find all ratings that are a maximum of 2/5).


NOTE:

Feel free to use either the file structure provided or feel free to create your own Golang module from scratch. This is your project to do as you see fit to solve the brief.

Note that you can choose to open an IDE in offline or in-browser online mode. We recommend offline mode.

In offline mode you can clone the project locally and later push the changes online to submit your solution.