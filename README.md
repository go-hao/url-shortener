# README.md

```
Usage: ./zero <option>
Options:
  set                   set go environment variables - GO111MODULE(on) and GOPROXY
  init                  initialize the project
  tidy                  tidy and format the project
  clean                 clean build content
  api <NAME>            create an api app in service/NAME/api
  rpc <NAME>            create a rpc app in service/NAME/rpc
  <MODEL OP> <PATH>     generate app models based on sql files in PATH/sql
                          - PATH      - service/NAME
                          - MODEL OP  - model generate app models with no cache
                                        modelc generate app models with cache
  gen <PATH>            generate or update app code
                          - PATH      - service/NAME/OPT
                          - OPT       - api|rpc
  build [PATH]          build app
                          - PATH      - service/NAME/OPT
                                        omit to build all apps
                          - OPT       - api|rpc
  run <PATH>            run app
                          - PATH      - service/NAME/OPT
                          - OPT       - api|rpc
  <DOCKER OP> [PATH]    shortcut of docker compose to manage docker service in docker
                          - PATH      - docker/NAME
                                        omit to manage all services
                          - DOCKER OP - up   start the service
                                        down stop the service
                                        ps   show status of the service
                                        logs show logs of the service
```