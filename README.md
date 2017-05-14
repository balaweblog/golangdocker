# golangdocker

<b>Basic golang applications in docker environment </b>

1. Build the docker image <br>
  docker build -t "golangimage" .
2. run the container with the image just created
  docker run --name=golangimagecontainer golangimage
3. start the container
  docker start golangimagecontainer
4. To check the logs inside container
  docker logs golangimagecontainer
  

<b>golang application with reading environment variables </b>
  
 1. Build the docker image 
  docker build -t "golangimage" .
 2. run the container with the image just created. make sure the environment variable is passed when running the container
  docker run --name=golangimagecontainer -i -t -e TEST_ENV=sampleEnv golangimage
 3. start the container
  docker start golangimagecontainer
