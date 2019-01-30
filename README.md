# rest-api-gpio



## Docker part : 

dl api :
* docker build -t apirestimg https://github.com/Arxsos/rest-api-gpio.git#master

launch api :
* docker run -d --privileged --publish 8001:8001 --name apirest apirestimg
