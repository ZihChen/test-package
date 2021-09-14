
ENV="local"
WORK_PATH=$(dirname $(readlink -f $0))
PROJECT_NAME=${WORK_PATH##*/}

echo "ENV=$ENV">>.env
echo "PROJECT_NAME=$PROJECT_NAME">>.env

docker-compose up -d
