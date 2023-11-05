#!/bin/bash

# Define variables
DATABASE_NAME=ketoai
COLLECTION_NAME='ingredients'
JSON_FILE_PATH=./ingredients.json

# Command to import new data from a JSON file into the specified collection
../utils/mongoimport --db $DATABASE_NAME --collection $COLLECTION_NAME --file $JSON_FILE_PATH --jsonArray

echo "Data has been inserted into the database."
