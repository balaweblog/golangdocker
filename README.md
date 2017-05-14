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
  

