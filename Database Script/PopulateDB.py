import csv
from io import StringIO
import psycopg2
import os
import glob

csvPath = "Database Script/test_data/"
#Change the database credentials according to your machine before running this script
conn = psycopg2.connect("host=localhost dbname=postgres user=postgres password=12345")


print("Connecting To Database...\n\n")
print("Database Connection Successful\n")

# Loop through each CSV
for filename in glob.glob(csvPath+"*.csv"):

    # Create a table name
    tablename = filename.replace("Database Script/test_data\\", "").replace(".csv", "")

    # Open file
    fileInput = open(filename, "r")

    #****NOTE******
    # PLEASE NOTE CSV FILE NAME SHOULD HAVE NO SPACES IN IT AND NAME SHOULD BE THE NAME YOU WANT IN DATABASE
    # Extract first line of file
    firstLine = fileInput.readline().strip()

    # Extract columns top row for database columns
    tablecolumns = firstLine.split(",")

    # Build SQL code to drop table if exists and create table
    sqlQueryCreate = 'DROP TABLE IF EXISTS '+ tablename + ";\n"
    sqlQueryCreate += 'CREATE TABLE '+ tablename + "("

    # Define columns for table
    for column in tablecolumns:
        sqlQueryCreate += column + " VARCHAR(64),\n"

    sqlQueryCreate = sqlQueryCreate[:-2]
    sqlQueryCreate += ");"

    cur = conn.cursor()
    cur.execute(sqlQueryCreate)
    conn.commit()
    
    print(tablename +" Table Created\n")

    #Insert Data in newly created table
    query = "COPY "+tablename+" FROM STDIN DELIMITER ',' CSV HEADER"
    
    #*****TIP******
    #I tried using copy_from but the data format in "Customer.csv" file was not compatible
    #So did some R&D and found copy_expert() works best with CSV files and PostgresSQL
    cur.copy_expert(query, open(filename, "r"))


    conn.commit()
    cur.close()
    print("Data Insertion In "+ tablename +" Complete\n")