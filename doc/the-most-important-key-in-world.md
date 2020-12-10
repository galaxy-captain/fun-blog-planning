> 加密技术是在互联网中传输数据最重要保护伞，当今主流的互联网服务都是在其保护下，才能够安全的发送和接收数据。

## 非对称加密算法
为实现在不安全的网络环境中传输数据的目的，非对称加密算法有着无法替代的必要性。 <br>
通俗的讲，非对称加密算法是一种传输数据的双方使用不同密钥进行加密和解密的算法，算法中使用了公钥和私钥，发送方将公钥发送给接收方，而接收方使用公钥加密数据并将其返回，发送方再使用自己的私钥进行解密。(注意：使用公钥加密的数据必须用私钥才能解密，加密数据无法使用公钥解密) <br>
上述过程只有公钥暴露在不安全的网络中，因此加密数据的安全性得到了保证。 <br>
#### 主要算法
* RSA算法，以三位作者的名字首字母命名，并且以该算法在2002年被授予图领奖，该算法奠定了网络安全传输的基石。
* ECC算法，椭圆曲线加密算法，该算法相较于RSA算法具有更好的安全强度和更平衡的性能开销，公钥长度显著小于RSA，加密和解密性能较为接近（RSA算法生成密钥的开销远高于解密）。
* DSA算法，该算法用于数字签名的生成和校验。
* DH算法，Diffie-Hellman算法提出了一种特别的密钥交换算法，能够在一次通信过程中完成密钥的交换，有效降低了TLS的开销。

## TLS协议详解
TLS协议是最早由网景公司(NetScape)在90年代提出的SSL协议发展而来，该协议与HTTP协议共同组成了HTTPS，有效的保证了当代互联网web传输的数据安全性。 <br>
当前HTTPS主要使用的TLS协议包括v1.2和v1.3，其中v1.3基于DH算法实现了仅需一次请求往返开销的密钥交换逻辑，较于v1.2的两次请求往返有了明显的提升。 <br>
#### 密钥交换过程
下图是RFC文档中描述的TLS v1.3握手动作的过程，可见其仅需在一个RTT就可完成，并且在ServerHello之后所有的数据也同样由密钥进行了加密。
![TLS v1.3 握手过程](https://img-blog.csdnimg.cn/20201210171350692.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2R1YW54bDExMjM0,size_16,color_FFFFFF,t_70)
1. TLS协议的握手动作由client端发起，首先client端初始化DH算法产生私钥/公钥，
然后client端将公钥和DH算法相关参数发送至server端，server端根据client端发送的公钥和DH算法相关参数生成自身的私钥/公钥，
最终server端将CA证书和server端公钥返回给client端，client端根据server端公钥和自身私钥计算得到"传输密钥"，同样server端根据client端公钥和自身私钥计算得到与client端相同的"传输密钥"。
![Diffie-Hellman算法图解](https://img-blog.csdnimg.cn/20201210184748114.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2R1YW54bDExMjM0,size_16,color_FFFFFF,t_70)
2. 后续的HTTP报文将使用"传输密钥"进行对称加密后传输。 <br>

#### 中间人攻击

## 证书颁发机构(CA, Certificate Authority)
CA是颁发数字证书的权威机构，

## 参考内容
* [非对称密码之DH密钥交换算法](https://blog.csdn.net/zbw18297786698/article/details/53609794)
* [TLS 1.3 握手过程详解](https://blog.csdn.net/zk3326312/article/details/80245756)
* [TLS 1.3 VS TLS 1.2 - 简书](https://www.jianshu.com/p/efe44d4a7501?utm_source=oschina-app)
* [TLS Protocol v1.3 - RFC8446](https://tools.ietf.org/html/rfc8446)
* [CA证书防御中间人攻击](https://www.sohu.com/a/376226794_100004247)

