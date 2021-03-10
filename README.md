# go_lang_learning

## Possible improvements
1. Use routes(with the library gin-gonic) to divide the functionality.
2. Use different ports for different get-post to easier monitor.
3. Save IP-Mac addresses of user to make it safer for them.
4. Hash the passwords of the users.
5. Divide business logic.
6. Divide main entry function from the rest
7. divide mysql functions from the rest
8. make sql connections persistent

## General Notes
1. Didn't use gorm, they generaly make the sql requests slower and the the requests we used were easy enough to manage.
2. Using: docker exec -it golang_app tail -f /data/logs/logs.txt  to tail logs
# errors encountered:

1. need to use go env -w GO111MODULE=auto otherwise go is losing path in containers
2. better have urls within quotation marks
3. mysql default handler doesn't work due to issues with the path reading a @ randomly?!?!?!
4. need to use go version 1.14 because 1.16 cant work inside a container, so ignore number 1
5. Logrus gives the following error: 
        root@e6fefefbdc1e:/go# go run main.go 
        \# golang.org/x/sys/unix
        src/golang.org/x/sys/unix/affinity_linux.go:14:20: undefined: _CPU_SETSIZE
        src/golang.org/x/sys/unix/affinity_linux.go:14:35: undefined: _NCPUBITS
        src/golang.org/x/sys/unix/affinity_linux.go:17:25: undefined: cpuMask 
