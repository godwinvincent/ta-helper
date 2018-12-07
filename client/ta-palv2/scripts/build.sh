cd ./..

npm run build

docker build -t info441tapal/client .

docker push info441tapal/client