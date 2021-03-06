# Start by building the application.                                                                                                                                                                            
FROM golang:1.10 as build                                                                                                                                                                                       
                                                                                                                                                                                                                
WORKDIR /go/src/app                                                                                                                                                                                             
COPY . . 
COPY * /usr/local/go/src/

RUN go get -d -v ./...
RUN go install -v ./...
                                                                                                                                                                                                                
# Now copy it into our base image.                                                                                                                                                                              
FROM gcr.io/distroless/base                                                                                                                                                                                     
COPY --from=build /go/bin/app /                                                                                                                                                                                 
CMD ["/app"]   
