# 说明
* blockchain
    + AtomInfo:单条交易信息,分为转账和节点发布的某人作弊的信息
    + Block:一个区块的结构
    + Chain:链的结构
    + Loader:加载本地信息到内存中,如用户信息,公钥等等
* communication
    + Matrix Problem Task: 递进关系 一个task包含一个problem,其中有两个待计算的矩阵
    + PreInfo:先导信息,等于是消息包的包头
    + Processor:处理对方节点发来信息的模块
    + Sender:发送广播信息
    + Server:开启后一直接收网络上其他节点的广播信息
* main
    + Node:主类
    + NodeInfo:储存节点信息
    + ProblemPublisher:~~发送问题的类,已经废弃~~
    + ProblemSolver:计算矩阵的类