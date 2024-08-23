# Sharpen

## Parameters

Operation name: **sharpen**

The following table describes the parameters you can configure.

| Parameter   | Description              | Valid value                                                  |
| :---------- | :----------------------- | :----------------------------------------------------------- |
| **[value]** | The degree of sharpness. | [50,399]A greater value indicates a clearer image. However, an overlarge value may result in image artifacts. We recommend that you set this parameter to 100 for optimal effects. |

## Examples

You can process images by using object URLs, OSS SDKs, or API operations. In this example, object URLs are used. For more information about how to use OSS SDKs and API operations to process images, see [IMG implementation modes](https://www.alibabacloud.com/help/en/oss/user-guide/img-implementation-modes#concept-m4f-dcn-vdb).

An image in the bucket named oss-console-img-demo-cn-hangzhou in the China (Hangzhou) region is used in this example. The following URL is used to access the image over the Internet:

http://127.0.0.1:8080/file/example.jpg![1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/en-US/4450040761/p531706.jpg)

Sharpen the source image. The degree of sharpness is set to 100. The following URL is used to process the image: http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/sharpen,100

![2](https://help-static-aliyun-doc.aliyuncs.com/assets/img/en-US/4450040761/p531708.jpg)