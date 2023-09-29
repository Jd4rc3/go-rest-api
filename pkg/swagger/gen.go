package swagger

//go:generate powershell -nop -Command {rm -Force server}
//go:generate powershell -nop -Command {mkdir server}
//go:generate swagger generate server --target server --name hello-api --spec swagger.yml --exclude-main
