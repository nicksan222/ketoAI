#!/bin/bash

# Define variables
DATABASE_NAME='your_db_name'
COLLECTION_NAME='your_collection_name'
JSON_FILE_PATH='/path/to/your/json_file.json'

# Command to import new data from a JSON file into the specified collection
mongoimport --db $DATABASE_NAME --collection $COLLECTION_NAME --file $JSON_FILE_PATH --jsonArray

echo "Data has been inserted into the database."
