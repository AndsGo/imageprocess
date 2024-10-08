# 图片水印

为保护OSS存储的图片或文件的所有权，防止资源未经授权被复制或使用，您可以为存储的资源增加水印。

## **使用场景**

- 版权保护：为保护自己的作品不被未授权使用或复制，需要在图片上加上水印来标识版权。
- 品牌推广：企业或个人为了宣传自己的品牌或标识，会在图片、视频或文档上加上带有品牌标志或名称的水印。
- 防止篡改：在某些官方文件、证书或报告上添加水印，可以增加篡改的难度，减少文件被伪造的风险。
- 抵制盗图：在网络环境中，图片很容易被他人下载和再次发布。加水印可以作为一种警示，减少他人直接盗用图片的情况。
- 法律要求：某些情况下，法律或合同条款可能要求在特定内容发布时必须加上水印，以符合规定。

## 注意事项

- 文字水印不要包号/=_#等字符
- text 文字最后使用`Urlencode`进行编码

## 参数说明

操作名称：**watermark**

相关参数如下：

- 基础参数

  | **参数**    | **是否必须** | **描述**                                                     | **取值范围**                                                 |
  | ----------- | ------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
  | **t**       | 否           | 指定图片水印或水印文字的透明度。                             | [0,100]默认值：100， 表示透明度100%（不透明）。              |
  | **g**       | 否           | 指定水印在图片中的位置。                                     | nw：左上north：中上ne：右上west：左中center：中部east：右中sw：左下south：中下se（默认值）：右下详情请参见下方基准点图片。 |
  | **x**       | 否           | 指定水印的水平边距， 即距离图片边缘的水平距离。这个参数只有当水印位置是左上、左中、左下、右上、右中、右下才有意义。 | [0,4096]默认值：10单位：像素（px）                           |
  | **y**       | 否           | 指定水印的垂直边距，即距离图片边缘的垂直距离， 这个参数只有当水印位置是左上、中上、右上、左下、中下、右下才有意义。 | [0,4096]默认值：10单位：像素（px）                           |
  | **voffset** | 否           | 指定水印的中线垂直偏移。当水印位置在左中、中部、右中时，可以指定水印位置根据中线往上或者往下偏移。 | [-1000,1000]默认值：0单位：像素（px）                        |

  水平边距、垂直边距、中线垂直偏移不仅可以调节水印在图片中的位置，当图片存在多重水印时，还可以调节水印在图中的布局。

  区域数值以及每个区域对应的基准点如下图所示。

  ![origin](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/2252359951/p2648.png)

- 文字水印参数

  | **参数**  | **是否必须** | **描述**                                              | **取值范围**                                                 |
  | --------- | ------------ | ----------------------------------------------------- | ------------------------------------------------------------ |
  | **text**  | 是           | 指定文字水印的文字内容，文字内容需进行urlencode编码。 | 最大字节长度为64个字符。                                     |
  | **color** | 否           | 指定文字水印的文字颜色，参数值为RGB颜色值。           | RGB颜色值，例如：000000表示黑色，FFFFFF表示白色。默认值：000000（黑色） |
  | **size**  | 否           | 指定文字水印的文字大小。                              | (0,1000]默认值：40单位：px                                   |

## 示例一：添加文字水印

您可以通过文件URL、API方式设置图片处理参数。本文以文件URL为例进行介绍。本文示例使用的examples/example.jpg，

该测试基于 

[综合示例](../README.md#comprehensive)

图片访问地址为：

http://127.0.0.1:8080/file/example.jpg

![原图](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/6929730761/p529184.jpg)

为example.jpg图片添加文字水印示例如下：

- 快速添加Hello World的文字水印

  对文字水印的内容Hello World进行图片处理URL为http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/resize,w_300,h_300/watermark,text_Hello。

  ![Hello World](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/5929730761/p529185.jpg)

- 添加文字水印时配置多个图片处理参数

  为example.jpg图片添加Hello World的文字水印的同时，需要对水印文字以及原图做如下相应处理：

  - 将example.jpg缩略为宽高300：`resize,w_300,h_300`
  - 水印内容为“Hello World”：`text_Hello%20World
  - 水印文字颜色为白色、字体大小为30：`color_FFFFFF,size_30`
  - 水印文字位置是右下、水平边距10、中线垂直偏移10：`g_se,x_10,y_10`

  图片处理的URL为：http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/resize,w_300,h_300/watermark,text_Hello%20World,size_30,color_FFFFFF,t_100,g_se,x_10,y_10

  ![图片处理1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/zh-CN/6929730761/p529186.jpg)