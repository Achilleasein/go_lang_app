# go_lang_learning



# errors encountered:

1. need to use go env -w GO111MODULE=auto otherwise go is losing path in containers
2. better have urls within quotation marks
3. mysql default handler doesnt work due to issues with the path reading a @ randomly?!?!?!
4. need to use go version 1.14 because 1.16 cant work inside a container, so ignore number 1
5. Logrus gives the following error: 
        root@e6fefefbdc1e:/go# go run main.go 
        \# golang.org/x/sys/unix
        src/golang.org/x/sys/unix/affinity_linux.go:14:20: undefined: _CPU_SETSIZE
        src/golang.org/x/sys/unix/affinity_linux.go:14:35: undefined: _NCPUBITS
        src/golang.org/x/sys/unix/affinity_linux.go:17:25: undefined: cpuMask 
