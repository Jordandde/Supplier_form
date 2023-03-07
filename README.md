# Rundoo Supplier List
### This project was created using:
### Frontend: 
    React with Typescript
    Semantic-ui
    axios
### Backend:
    Go
    Mongodb

### How to run?
    enter the client folder and run npm install. The run npm start
    in a separate terminal enter the server folder and run go run main.go. If there are package issues just install all the required packages using go get {insert package name here}
    You will also need a mongoDB cluster/collection running at a predefined port of your choosing

### Env file? 
    Create a .env file in the base of the server directory and add these values to the file. Replace the strings with whatever your local configuration is:
    ```
    DB_URI="mongodb://localhost:27017"
    DB_NAME="test"
    DB_COLLECTION_NAME="SupplierList"
    ```

### Summary of somethings I would change if given more time?
    - Create custom components rather than using Semantic-UI
    - Write the JavaScript to replace the file picker's value on submission, note, most forms seem to clear it on entry and place a "picked files" section below the file picker to avoid this issue
    - Flesh out the Go endpoints to exapnd on features like adding more fields or formating options
    - Research how to store images more efficiently, this project converts the images to Base64 encoded strings which are easy to work with on a time constraint, but much larger than storing the actual file I believe.
    - Host the application myself on Vercel
    - A few more things that I will save to discuss later
