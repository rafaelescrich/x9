FROM golang:1.13.1

COPY .  $GOPATH/src/github.com/ResultadosDigitais/x9
WORKDIR $GOPATH/src/github.com/ResultadosDigitais/x9
CMD ["go", "run",  "main.go"] 
