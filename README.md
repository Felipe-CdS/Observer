## Observer
The idea is to write a library that watches markdowns files in a bl-
og folder. The library runs in a go routine inside the blog server
and every time the author changes a post .md file the library sees
the change and parses the .md file to html saving that on a posts db
table.

## Waiting
Right now waiting for this commit. The internal library that watches
inotify linux calls cant see CLOSE_WRITE calls, so every event is
called more than once because the go functions call write inside 
other functions.

https://github.com/fsnotify/fsnotify/pull/651

Maybe rewrite this lib...
