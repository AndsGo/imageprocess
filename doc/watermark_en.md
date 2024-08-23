# Add watermarks

## **Scenarios**

- Copyright protection: To help prevent your works from being replicated or used without your authorization, you can add watermarks to images to identify the copyright.
- Brand promotion: Enterprises or individuals can add watermarks with brand logos or names to images, videos, or documents to promote their brands or logos.
- Tamper prevention: Adding watermarks to official documents, certificates, or reports can increase the difficulty of tampering and reduce the risk of document forgery.
- Image plagiarism prevention: Images can be easily downloaded and republished by Internet users. You can add watermarks to images to show that they are copyrighted.
- Legal compliance requirement: In some cases, you must add watermarks to meet legal compliance requirements when you publish specific content of legal terms or contract terms.

## Usage notes

- Traditional Chinese characters cannot be used as text watermarks.(Will be supported later)

## Parameters

Action: **watermark**

The following tables describe the parameters that you can configure when you add watermarks to images.

- Basic parameters

  | **Parameter** | **Required** | **Description**                                              | **Valid value**                                              |
  | ------------- | ------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
  | **t**         | No           | The opacity of the watermark.                                | [0,100]Default value: 100. The value 100 specifies that the watermark is opaque. |
  | **g**         | No           | The position of the watermark on the image.                  | nw: upper left.north: upper middle.ne: upper right.west: middle left.center: center.east: middle right.sw: lower left.south: lower middle.se: lower right. This is the default value.For the precise position that is specified by each value, see the following figure. |
  | **x**         | No           | The horizontal margin, which indicates the horizontal distance between the watermark and the image edge. This parameter takes effect only when the watermark is on the upper left, middle left, lower left, upper right, middle right, or lower right of the image. | [0,4096]Default value: 10.Unit: pixel.                       |
  | **y**         | No           | The vertical margin that specifies the vertical distance between the watermark and the image edge. This parameter takes effect only when the watermark is on the upper left, upper middle, upper right, lower left, lower middle, or lower right of the image. | [0,4096]Default value: 10.Unit: pixel.                       |
  | **voffset**   | No           | The vertical offset from the middle line. When the watermark is on the middle left, center, or middle right of the image, you can specify the vertical offset of the watermark along the middle line. | [-1000,1000]Default value: 0.Unit: pixel.                    |

  You can use parameters x, y, and voffset to adjust the position of a watermark on an image. You can also use these parameters to adjust the watermark layout when the image contains multiple watermarks.

  The following figure shows the positions of watermarks based on coordinates.![origin](https://help-static-aliyun-doc.aliyuncs.com/assets/img/en-US/4956348951/p2648.png)

  

- Text watermark parameters

  | **Parameter** | **Required** | **Description**                                              | **Valid value**                                              |
  | ------------- | ------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
  | **text**      | Yes          | The content of the text watermark. The text content must be `Urlencode` | A Chinese string before Url-encoding cannot exceed 64 characters in length. |
  | **color**     | No           | The color of the text watermark. The valid values for this parameter are RGB color values. | For example, 000000 specifies black, and FFFFFF specifies white.Default value: 000000. |
  | **size**      | No           | The size of the text watermark.                              | (0,1000]Default value: 40.Unit: pixel.                       |

## Encode watermark-related parameters

1. Encode watermark-related parameters by using `Urlencode`.

**Important**

The encoding results can be used only in specific parameters in watermark operations. Do not use the encoding results in signature strings.

## Example 1: Add a text watermark to an image

You can set image processing parameters through file URL and API. This article takes file URL as an example. This article uses examples/example.jpg.

This test is based on  [comprehensive example](../README_en.md#comprehensive)

The image access address is:

http://127.0.0.1:8080/file/example.jpg

![原图](https://help-static-aliyun-doc.aliyuncs.com/assets/img/en-US/0327006071/p529184.jpg)

The following examples show how to add a text watermark to example.jpg:

- Add the string "Hello World" to the image as a text watermark

  Url-encode the string "Hello World" into a URL-safe string. The encoding result is `Hello%20World` and the URL used to process the image is http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/watermark,text_Hello%20World.

  ![Hello World](https://help-static-aliyun-doc.aliyuncs.com/assets/img/en-US/9227006071/p529185.jpg)

- Configure multiple IMG parameters when you add a text watermark to the image

  In this example, IMG parameters are configured to perform the following operations on the source image and the text watermark "Hello World":

  - Resize the source image example.jpg to 300 × 300 pixels by using `resize,w_300,h_300`
  - Add the string "Hello World" to the source image as a text watermark by using `text_Hello%20World`.
  - Set the color of the watermark text to white and the size of the text to 30 pixels by using `color_FFFFFF,size_30`.
  - Set the position of the text watermark to lower right, the horizontal margin to 10 pixels, and the vertical offset from the middle line to 10 pixels by using `g_se,x_10,y_10`.

  The following URL is used to process the image based on the preceding parameters: http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/resize,w_300,h_300/watermark,size_30,text_Hello%20World,color_FFFFFF,t_100,g_se,x_10,y_10.

  ![图片处理1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/en-US/0327006071/p529186.jpg)