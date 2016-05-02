# Direct Buildpack deployment to Marathon
This folder include example json file for buildpack and start.sh *

## Deplyment Steps.

1. Step to Deploy
```
./start.sh -m shipped-tx3-control-01.tx3.shipped-cisco.com -u synthetic-mon -p *** -c go-config.json 
```

2. Following are the configuration that need to update json file for the build pack for deployment.
  - "Update Host Name " : Please put active salve name.
  - "Docker Image ": Enter valid docker image "shippedrepos-docker-{project id}.bintray.io/{project name}:{github last successfull commit}"
  - "Change Host Port and Service Port" : One app can't deploy on same port number, So make sure port is avaible.
  - "Change Env Name" : Change Env Service name and also Change Id.
  - "Update Uri" : add valid dockercfg.
  
