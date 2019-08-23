
## Deploy
1. Start k8s cluster 
```sh
git clone https://github.com/danielinclouds/kops-istio-test-env.git
cd kops-istio-test-env
./start
cd ..
````
2. Build and push docker image
```sh
cd server
     docker build -t sli-server .
     docker push danielinclouds/sli-server:latest
cd ..
```
3. Deploy sli-server helm chart
```sh
pushd server/sli-server
    helm install --name sli-server . --values values.yaml
popd
```
4. Deploy lambda with serverless framework
```sh
pushd client
    export SERVER_URL=http://K8S_LB_IP:80/health
    make deploy
popd
```


## Cleanup 
1. Delete k8s cluster
```sh
cd kops-istio-test-env
./stop
cd ..
```
2. Delete serverless framework deployment
```sh
pushd client
    sls remove
popd
```



## Priority future improvements
- It looks like there are 2 lambda instances running at the same time
- Client should set a header so that it's easy to find client traffic in logs
- Standard sli-server log should not show connections from k8s probes
- Create debug flag for sli-server that shows connections from k8s probes


## Other future improvements
- Add other health checks like K8s API and DNS:
- Check availability of K8s API
- Check if deploying app works
- Check if External DNS can create DNS record
- Check if Cert Manager can request certificate
- Check latency of request health check connection
- Model this health check on PCF Healthwatch

