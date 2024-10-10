# MediaServer
A media server that serves Images/videos/etc

Golang using the HTTP and template standard packages
PostgreSQL
CSS

A file distributor that shares videos, images on a local network. 

Will only work for linux filesystems, WSL2, ubuntu etc since this is only built for that filesystem.



----------------------------------------------
To build the webserver, clone the repo and run : go build -o MediaServer && ./MediaServer
The server will automatically run on your localhost:8080

- You will be able to create users with the signup, and login after creation while staying logged in with cookies created on login.

- Putting things in the media folder(static/media) will automatically render them on your webpage after a refresh (feel free to make a clone to access the folder, NOT in a onedrive/google drive folder)  

- search button is case sensitive and will show anything that has any part of the search term: e.g "terst" will find any Interstellar files.

----------------------------------------------