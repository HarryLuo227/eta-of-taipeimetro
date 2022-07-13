# ETA of Taipei Metro

This is a simple web service project which implement to get some data from MongoDB for gin web framework practice.

## MongoDB data source

### Database and Collections

> <b>Database name</b>
>  
> testDB
> 
> <b>Collection name</b>
> 
> LineTransfer, S2STravelTime, StationOfLine

If you want to run on your local machine, you need to

1. Install MongoDB.
2. Get the data from [Taipei metro API](https://ptx.transportdata.tw/MOTC/?urls.primaryName=%E8%BB%8C%E9%81%93V2#/), and download the json file.

    `The collection name is used in api routes.`

3. Import json file into MongoDB, Example below.
    ```
    mongoimport --uri mongodb://localhost:27017/testDB --collection LineTransfer --jsonArray --file <your json file path>
    ```

## Features

**Goal : Calculate the time duration you take between any two stations of Taipei Metro.**

Get the time duration between any two stations on the same line.

## Data reference

[Metro Taipei API service](https://www.metro.taipei/cp.aspx?n=BDEB860F2BE3E249)
