# golangdocker

<b>Basic golang applications in docker environment </b>

1. Build the docker image <br>
  docker build -t "golangimage" . <br>
2. run the container with the image just created <br>
  docker run --name=golangimagecontainer golangimage <br>
3. start the container <br>
  docker start golangimagecontainer <br>
4. To check the logs inside container <br>
  docker logs golangimagecontainer <br>
  

<b>golang application with reading environment variables </b>
  
 1. Build the docker image <br>
  docker build -t "golangimage" . <br>
 2. run the container with the image just created. make sure the environment variable is passed when running the container <br>
  docker run --name=golangimagecontainer -i -t -e TEST_ENV=sampleEnv golangimage <br>
 3. start the container <br>
  docker start golangimagecontainer <br>

<b>golang application getting github.com packages </b>


1. Build the docker image <br>
  docker build -t "golangimage" . <br>
2. run the container with the image just created <br>
  docker run --name=golangimagecontainer golangimage <br>
3. start the container <br>
  docker start golangimagecontainer <br>
4. To check the logs inside container <br>
  docker logs golangimagecontainer <br>
  

<b>golang application docker compose up </b>

Your working directory should be like $ mkdir $GOPATH/src/golangdocker <br>
Inside put your files docker-compose.yml and main.go <br>

myapp -> service name <br>
image -> we use the already created golang image <br>
volume -> Attach the volumes and working directory. <br>
command -> the initial command to run  <br>
links - relative links like connecting redis or sql <br>
environment - environment variable <br>
redis image - image to external <br>


1. $ docker-compose  <br>
or 
2. docker-compose run myapp go run main.go <br>
myapp - service name <br>

