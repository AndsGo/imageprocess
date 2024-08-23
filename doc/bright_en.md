# Brightness

## Parameters

Operation name: **bright**

The following table describes the parameters you can configure.

| Parameter   | Description                                             | Valid value                                                  |
| :---------- | :------------------------------------------------------ | :----------------------------------------------------------- |
| **[value]** | The percentage by which to adjust the image brightness. | [-100, 100]A value smaller than 0 indicates that the brightness of the image is decreased.A value of 0 indicates that the brightness of the image is not changed.A value greater than 0 indicates that the brightness of the image is increased. |

## Examples

You can set image processing parameters through file URL and API. This article takes file URL as an example. This article uses examples/example.jpg.

This test is based on  [comprehensive example](../README_en.md#comprehensive)

The image access address is:

http://127.0.0.1:8080/file/example.jpg

![1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/en-US/5630040761/p531695.jpg)



- Increase the brightness of the image by 50 percent

  The following URL is used to process the image: http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/bright,50

  ![2](https://help-static-aliyun-doc.aliyuncs.com/assets/img/en-US/5630040761/p531698.jpg)

- Decrease the brightness of the image by 50 percent

  The following URL is used to process the image: http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/bright,-50