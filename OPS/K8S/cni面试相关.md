
# K8s通信问题

1. 容器间通信：即同一个Pod内多个容器间通信，通常使用loopback来实现。
2. Pod间通信：K8s要求,Pod和Pod之间通信必须使用Pod-IP 直接访问另一个Pod-IP
3. Pod与Service通信：即PodIP去访问ClusterIP，当然，clusterIP实际上是IPVS或iptables规则的虚拟IP，是没有TCP/IP协议栈支持的。但不影响Pod访问它.
4. Service与集群外部Client的通信，即K8s中Pod提供的服务必须能被互联网上的用户所访问到。

需要注意的是，k8s集群初始化时的service网段，pod网段，网络插件的网段，以及真实服务器的网段，都不能相同，如果相同就会出各种各样奇怪的问题，而且这些问题在集群做好之后是不方便改的，改会导致更多的问题，所以，就在搭建前将其规划好。



# 1




> From: [灵雀云k8s面试题剖析 - 脉脉 (maimai.cn)](https://maimai.cn/article/detail?fid=1555220859&efid=GH590E-KjnTlxSn1Q-0Y8Q)
> 
## 一、简述你知道的几种CNI网络插件，并详述其工作原理。K8s常用的CNI网络插件 （calico && flannel），简述一下它们的工作原理和区别.

1、calico根据 iptables规则 进行路由转发，并没有进行封包，解包的过程，这和flannel比起来效率就会快多

calico包括如下重要组件：Felix，etcd，BGP Client，BGP Route Reflector。下面分别说明一下这些组件。

Felix：主要负责路由配置以及ACLS规则的配置以及下发，它存在在每个node节点上。

etcd：分布式键值存储，主要负责网络元数据一致性，确保Calico网络状态的准确性，可以与kubernetes共用；

BGPClient(BIRD), 主要负责把 Felix写入 kernel的路由信息分发到当前 Calico网络，确保 workload间的通信的有效性；

BGPRoute Reflector(BIRD), 大规模部署时使用，摒弃所有节点互联的mesh模式，通过一个或者多个 BGPRoute Reflector 来完成集中式的路由分发

通过将整个互联网的可扩展 IP网络原则压缩到数据中心级别，Calico在每一个计算节点利用 Linuxkernel 实现了一个高效的 vRouter来负责数据转发，而每个vRouter通过 BGP协议负责把自己上运行的 workload的路由信息向整个Calico网络内传播，小规模部署可以直接互联，大规模下可通过指定的BGProute reflector 来完成。这样保证最终所有的workload之间的数据流量都是通过 IP包的方式完成互联的。

## 2、Flannel的工作原理：

Flannel实质上是一种“覆盖网络(overlay network)”，也就是将TCP数据包装在另一种网络包里面进行路由转发和通信，目前已经支持UDP、VxLAN、AWS VPC和GCE路由等数据转发方式。

默认的节点间数据通信方式是UDP转发。

工作原理：

数据从源容器中发出后，经由所在主机的docker0虚拟网卡转发到flannel0虚拟网卡（ 先可以不经过docker0网卡，使用cni模式 ），这是个P2P的虚拟网卡，flanneld服务监听在网卡的另外一端。

Flannel通过Etcd服务维护了一张节点间的路由表，详细记录了各节点子网网段 。

源主机的flanneld服务将原本的数据内容UDP封装后根据自己的路由表投递给目的节点的flanneld服务，数据到达以后被解包，然后直接进入目的节点的flannel0虚拟网卡，然后被转发到目的主机的docker0虚拟网卡，最后就像本机容器通信一下的有docker0路由到达目标容器。

flannel在进行路由转发的基础上进行了封包解包的操作，这样浪费了CPU的计算资源。