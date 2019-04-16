# Docker

+ docker images查看镜像
+ 给一个镜像名字和tag
    ```shell
    docker tag [image id] [name]:[版本]
    docker tag b03b74b01d97 docker-redis:0.0.1
    ```
+ 保存镜像
    ```shell
    docker save -O fileName imageName
    docker save -o tensorrt_19.03_py2.tar 1327897c9f27
    ```
+ 加载镜像
    ```shell
    docker load --input rocketmq.tar
    #或
    docker load < rocketmq.tar
    ```

+ 复制文件
    ```shell
    #从宿主机到容器
    docker cp fileName imageID:path
    docker cp ./fileName 30504d38a731:/root
    #从容器到宿主机
    docker cp containID:/opt/testnew/file.txt  /opt/test/
    ```
+ 容器
    + 创建容器
        ```shell
        #普通
        docker run imageId|imageName
        #交互式
        docker run -it ...
        ```
    + 删除容器
        ```shell
        docker container kill containerID
        ```
    + 查看容器
        ```shell
        docker container ls
        ```

