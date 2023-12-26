# microDMS
 A micro DMS for fun in Go with a CouchDB as backend. The main purpose is to learn a bit about Go programming and CouchDB.

## pre-requirement
 - CouchDB
 - Environmental variables
     - DMS_DB -> Name of the Database
     - DMS_HOST -> Host address of the CouchDB server
     - DMS_STORAGE -> Filesystem location where to store the documents

## Usage

 ```$>./dms add -f <local filename> -d <description> <label1> <label2> .. <labeln>
 $>./dms list -label <label1> <label2> .. <labeln>
 $>./dms list -id <id>
 $>./dms remove -id <id>
 $>./dms update -id <id> -d <description>
 $>./dms update -id <id> -addlabel <label> .. <labeln>
 $>./dms update -id <id> -rmlabel <label> .. <labeln>
 ```
