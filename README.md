# counter-service
counter service with realisation of pattern SAGA

### Startup
1. run command `make rebuild`
2. disable needless services in config [local_config.yaml](cmd%2Frealtime%2Flocal_config.yaml) (set `enable: false` in config blocks)

### System Design
![counter.png](files%2Fcounter.png)