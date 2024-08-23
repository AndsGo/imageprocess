# Convert image formats

## **Scenarios**

- Adaptive to different devices and platforms: Different browsers, operating systems, or mobile devices may support different image formats. For example, the WebP format provides better compression efficiency on browsers. With OSS image format conversion, uploaded images can be converted into multiple formats to ensure compatibility and optimal display effects on various devices.
- Reduced storage costs: Some image formats, such as WebP, usually provide smaller sizes than other traditional formats while maintaining visual quality. You can convert image formats to reduce storage usage without compromising image quality. This way, storage costs are reduced.
- Unified resource management: In e-commerce, social networking, media, and other industries, images uploaded by a large number of users need to be standardized. You can convert uploaded images to a specific format for subsequent management and distribution.

## Usage notes

- If an image processing (IMG) request includes the format and resize parameters, we recommend that you place the format parameter at the end.

  Example: image/resize,w_100/format,jpg

- If an IMG request includes the format, resize, and watermark parameters, we recommend that you place the format parameter after the resize parameter.

  Example: `image/resize,w_100/format,jpg/watermark,...`

- If the source image does not support alpha channels, the format of the source image is converted to a format that supports alpha channels. Formats that support alpha channels include PNG, WebP, and BMP. By default, OSS fills the transparent area with white.

- You cannot use OSS to fill the transparent area with black.

## Parameter description

Parameter name: **format**.

The following table describes the valid values of the format parameter.

| **Valid value** | **Description**                                              |
| --------------- | ------------------------------------------------------------ |
| **jpg**         | Converts the format of a source image to JPG.**Important**Images in the HEIC format that support alpha channels cannot be converted to JPG images. |
| **png**         | Converts the format of a source image to PNG.                |
| **webp**        | Converts the format of a source image to WebP.               |
| **bmp**         | Converts the format of a source image to BMP.                |
| **gif**         | Converts the format of a source image to GIF. The conversion takes effect only when the source image is also a GIF image. If the source image is not in the GIF format, the processed image is stored in the original format. |
| **tiff**        | Converts the format of a source image to TIFF.               |

## Examples

You can set image processing parameters through file URL and API. This article takes file URL as an example. This article uses examples/example.jpg.

This test is based on  [comprehensive example](../README_en.md#comprehensive)

The image access address is:

http://127.0.0.1:8080/file/example.gif

![gif](https://help-static-aliyun-doc.aliyuncs.com/assets/img/en-US/5703863061/p139212.png)

- Convert the format of the source image to PNG

  The URL used to process the source image is http://127.0.0.1:8080/file/example.gif?x-oss-process=image/format,png

  ![png](https://help-static-aliyun-doc.aliyuncs.com/assets/img/en-US/6703863061/p139213.png).

- Convert the format of the source image to JPG that supports gradual display

  Configure the parameters based on the following requirements:

  - Apply gradual display: `interlace,1`
  - Convert the format of the source image to JPG: `format,jpg`

  The URL used to process the source image is http://127.0.0.1:8080/file/example.gif?x-oss-process=image/format,jpg.

  ![img](https://help-static-aliyun-doc.aliyuncs.com/assets/img/en-US/3956348951/p2555.jpg)

- Resize the image to a width of 200 pixels and convert the format of the image to WebP

  Configure the parameters based on the following requirements:

  - Resize the image to a width of 200 pixels: resize,w_200
  - Convert the format of the image to WebP: format,webp

  The URL used to process the source image is http://127.0.0.1:8080/file/example.gif?x-oss-process=image/resize,w_200,h_200/format,webp

Currently, conversion to webp does not support animated images.