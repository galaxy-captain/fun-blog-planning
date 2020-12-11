> 加密技术是在互联网中传输数据最重要保护伞，当今主流的互联网服务都是在其保护下，才能够安全的发送和接收数据。而它们中有几个绝对不能泄露的密码，一旦丢失就可能造成重大的经济损失。

## 非对称加密算法
为实现在不安全的网络环境中传输数据的目的，非对称加密算法有着无法替代的必要性。 <br>
通俗的讲，非对称加密算法是一种传输数据的双方使用不同密钥进行加密和解密的算法，算法中使用了公钥和私钥，发送方将公钥发送给接收方，而接收方使用公钥加密数据并将其返回，发送方再使用自己的私钥进行解密。(注意：使用公钥加密的数据必须用私钥才能解密，加密数据无法使用公钥解密) <br>
上述过程只有公钥暴露在不安全的网络中，因此加密数据的安全性得到了保证。 <br>

#### 主要算法
* RSA算法，以三位作者的名字首字母命名，并且以该算法在2002年被授予图领奖，该算法奠定了网络安全传输的基石。
* ECC算法，椭圆曲线加密算法，该算法相较于RSA算法具有更好的安全强度和更平衡的性能开销，公钥长度显著小于RSA，加密和解密性能较为接近（RSA算法生成密钥的开销远高于解密）。
* DSA算法，该算法用于数字签名的生成和校验。
* DH算法，Diffie-Hellman算法提出了一种特别的密钥交换算法，能够在一次通信往返过程(RTT)中完成密钥的交换，有效降低了开销。

## TLS协议详解
TLS协议是最早由网景公司(NetScape)在90年代提出的SSL协议发展而来，该协议与HTTP协议共同组成了HTTPS，有效的保证了当代互联网web传输的数据安全性。 <br>
当前HTTPS主要使用的TLS协议包括v1.2和v1.3两个版本，其中v1.3基于DH算法实现了仅需一次RTT开销的密钥交换逻辑，较于v1.2的两次请求往返有了明显的提升。 <br>

#### 密钥交换过程
下图是RFC文档中描述的TLS v1.3握手动作的过程，其基于ECDHE算法(DH算法与ECC算法的结合)，可见仅需在一个RTT就可完成，并且在ServerHello之后所有的数据也同样由密钥进行了加密。
![TLS v1.3 握手过程](https://img-blog.csdnimg.cn/20201210171350692.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2R1YW54bDExMjM0,size_16,color_FFFFFF,t_70)
1. TLS协议的握手动作由client端发起，首先client端初始化DH算法产生私钥/公钥，
然后client端将公钥和DH算法相关参数发送至server端，server端根据client端发送的公钥和DH算法相关参数生成自身的私钥/公钥，
最终server端将`CA证书`和server端公钥返回给client端，client端根据server端公钥和自身私钥计算得到`传输密钥`，同样server端根据client端公钥和自身私钥计算得到与client端相同的`传输密钥`。
![Diffie-Hellman算法图解](https://img-blog.csdnimg.cn/20201210184748114.png?x-oss-process=image/watermark,type_ZmFuZ3poZW5naGVpdGk,shadow_10,text_aHR0cHM6Ly9ibG9nLmNzZG4ubmV0L2R1YW54bDExMjM0,size_16,color_FFFFFF,t_70)
2. 后续的HTTP报文将使用`传输密钥`进行对称加密后传输，通常使用AES-GCM、ChaCha20-Poly1305等AEAD算法。 <br>

#### 中间人攻击
中间人攻击是通过在client/server的通道中插入第三方，该第三方将自己伪装成client访问的server和向server发送请求的client，以此达到欺骗真正的client/server的目的，从而真正的client/server将数据发送到该第三方。 <br>
![中间人攻击描述图](https://img-blog.csdnimg.cn/20201212013855352.png) <br>
这种攻击方式利用了未知数据来源的风险，当client/server无法验证数据来源就无条件信任了收到的公钥时，便为中间人提供了可乘之机。 <br>
为了防御中间人攻击，需要提供有效的手段对数据来源进行校验。`CA证书`是由大型权威机构颁发的证明，通过ECDSA算法产生私钥/公钥的`数字签名机制`，实现了对数据来源合法性的校验。<br>

## 数字签名机制
> “完整性”是指数据在生成后从未被非法修改

`CA证书`是由可信任的机构生成的公钥/私钥的信息，公钥(VerificationKey)包含在CA证书中，私钥(SigningKey)由server端进行严密的保管。
1. 当server端将ECDHE算法生成的公钥(PublicKey)发送到client端时，首先需要使公钥(PublicKey)通过私钥(SigningKey)生成签名，然后再将公钥(PublicKey)、签名、CA证书一起发送到client端。
2. 在client端收到server端数据后，先对CA证书进行完整性校验，再根据CA证书中的公钥(VerificationKey)和server端提供的签名来校验公钥(PublicKey)的完整性，只有在通过这两项校验后，才能真正信任server端发送的数据。

#### 公钥完整性校验
前述章节已经说明过，中间人攻击是通过向client端发送非法的公钥(PublicKey)以达到骗取信任的目的，因此数字签名机制提供了针对公钥(PublicKey)的完整性进行校验的方式。<br>
因为server端发送的公钥(PublicKey)已经由中间人无法得到的私钥(SigningKey)生成签名，而签名又会在client端由CA证书中的公钥(VerificationKey)进行校验，
所以中间人无法在不篡改CA证书中公钥(VerificationKey)的前提下生成可以被client端验证通过的签名。<br>

#### CA证书完整性校验
由于中间人必须篡改CA证书中的公钥(VerificationKey)才能达到攻击的目的，所以client端必须对CA证书的完整性进行校验。<br>
CA证书在颁发时，颁发机构会使用`自有私钥`对证书进行签名，并将该签名记录在CA证书中，同时CA证书中还包括主体域名、有效时间、颁发机构信息等。
client端在收到server端发送的CA证书时，先校验主体域名是否一致、证书是否过期，再根据颁发机构信息在client端本地存储的证书链中找到与颁发机构的`自有私钥`对应的公钥，然后完成对CA证书的完整性的校验。<br>
client端本地存储的证书链是必须可信的，其通常是浏览器和操作系统内置的证书链，无法被中间人篡改(该操作只有client端已经被攻击者控制后才能实现，此时任何通信加密手段都已经失去作用)。

## 总结
通过对加密技术的了解，我们能够非常明显的发现，CA证书对应的私钥和颁发CA证书时签名使用的私钥，任何一个私钥的泄露都将使加密手段失效，从而引发巨大的风险。<br>
CA证书对应的私钥泄露仅影响一家网站，而CA机构签名使用的私钥泄露时，将对全球互联网的安全造成重大破坏，甚至引发系统性风险。<br>
因此CA机构保存私钥的手段和签发CA证书的流程便成为了守护互联网安全的基石，拥有着全球共认的国家级规范标准，和相应的法律法规来确保安全。

#### 参考链接
* [DH密钥交换算法](https://blog.csdn.net/zbw18297786698/article/details/53609794)
* [TLS 1.3 握手过程详解](https://blog.csdn.net/zk3326312/article/details/80245756)
* [TLS 1.3 vs TLS 1.2](https://www.jianshu.com/p/efe44d4a7501?utm_source=oschina-app)
* [RFC8446: TLS Protocol v1.3](https://tools.ietf.org/html/rfc8446)
* [TLS 1.3 中间人攻击和ECDSA机制](https://github.com/WeMobileDev/article/blob/master/%E5%9F%BA%E4%BA%8ETLS1.3%E7%9A%84%E5%BE%AE%E4%BF%A1%E5%AE%89%E5%85%A8%E9%80%9A%E4%BF%A1%E5%8D%8F%E8%AE%AEmmtls%E4%BB%8B%E7%BB%8D.md)
* [CA机构如何存储私钥](https://security.stackexchange.com/questions/24896/how-do-certification-authorities-store-their-private-root-keys)
