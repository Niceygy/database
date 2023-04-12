# Examples
These files are the standard for the db
## Files
### db.list
Holds all known db names
all dbs are in their own folder
the folder has the same name as is in the db.list
### colums.list
Holds all known colums in that database
for example: if a database holds users, the colums.list would hold
``` usernames ```
``` passwords ```
``` emails ```
``` ids ```
### *.list
Holds all known values in that colum
for example, a database called users would have a file called usernames.list
this would contain all known usernames
Each row has a UDID (Unique Data ID)
This is a unique number that is used to identify the row accross files
## Searching for data
### Searching for spesific value in a colum
A user would request a row by giving the db name, the colum name, and a value to search for
The server would then search the colum file for the value
It would then take the UDID and search the requested coulm file for the UDID
It would then return the data in that row
### Searching for a row
A user would request a row by giving the db name, the ( UDID or a value to search for )
The server would then search the colum file for the value
It would then take the UDID and search all colum files for the UDID
It would then return the data in that row (in an array?)