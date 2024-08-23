# Custom crop

## **Scenarios**

- Web page design and production: When you design the layout of a web page, you may need to crop images to a specific size to fit web page elements, such as avatars, background images, and product displays.
- Custom image size required by social media: Different social media platforms have different size requirements for images that you want to upload, such as thumbnails, images that you can post, and story images. You need to crop source images based on the recommended size to achieve optimal display performance.
- Mobile app development: You must crop images such as icons, startup pages, and embedded pictures in an app based on specific specifications to ensure that the images are displayed as expected on devices with different resolutions and screen sizes.
- Image database management: To meet the sorting and archiving requirements of institutions that have a large number of image resources, such as libraries and archives, you may need to crop images to a preset size.

## Usage notes

- If the specified starting X coordinate and Y coordinate exceed those of the source image, the `BadRequest` error code and the Advance cut's position is out of image. error message are returned.
- If the width and height specified from the starting point exceed those of the source image, the source image is cropped to the boundaries.
- You can use object URLs, OSS SDKs, or API operations to configure image processing (IMG) parameters that are used to process images. In this topic, object URLs are used. You can use object URLs to configure IMG parameters only for public-read-write images. If you want to configure IMG parameters for private images, use OSS SDKs or call API operations. For more information, see [IMG implementation modes](https://www.alibabacloud.com/help/en/oss/user-guide/img-implementation-modes#concept-m4f-dcn-vdb).

## Parameters

Action: **crop**.

The following table describes the parameters.

| **Parameter** | **Description**                                              | **Value range**                                              |
| ------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
| **w**         | The width that you want to crop.                             | [0, image width]Default value: the maximum value.            |
| **h**         | The height that you want to crop.                            | [0, image height]Default value: the maximum value.           |
| **x**         | The X coordinate of the area that you want to crop. The default value is the X coordinate of the upper-left corner of the image. | [0, image bound]                                             |
| **y**         | The Y coordinate of the area that you want to crop. The default value is the Y coordinate of the upper-left corner of the image. | [0, image bound]                                             |
| **g**         | The position of the area that you want to crop in a 3 x 3 grid. The image is located in a 3 x 3 grid. The grid has nine tiles. | nw: upper leftnorth: upper middlene: upper rightwest: middle leftcenter: centereast: middle rightsw: lower leftsouth: lower middlese: lower rightFor more information about how to calculate the position of each tile, see the  table below. |

The following table describes how to calculate the position of each tile in a 3 x 3 grid. srcW specifies the width of the source image and srcH specifies the height of the source image.

| **Tile** | **Calculation method**     |
| -------- | -------------------------- |
| nw       | 0, 0                       |
| north    | srcW/2 - w/2, 0            |
| ne       | srcW - w, 0                |
| west     | 0, srcH/2 - h/2            |
| center   | srcW/2 - w/2, srcH/2 - h/2 |
| east     | srcW - w, srcH/2 - h/2     |
| sw       | 0, srcH - h                |
| sourth   | srcW/2 - w/2, srcH - h     |
| se       | srcW - w, srcH - h         |

## Examples

You can set image processing parameters through file URL and API. This article takes file URL as an example. This article uses examples/example.jpg.

This test is based on  [comprehensive example](../README_en.md#comprehensive)

The image access address is:

http://127.0.0.1:8080/file/example.jpg
![原图](https://help-static-aliyun-doc.aliyuncs.com/assets/img/en-US/8850665071/p674595.jpg)

- Crop an image from the starting point (800, 50) to the boundaries

  Configure the parameters based on the following requirements:

  - From the starting point (800, 50): `crop,x_800,y_50`.
  - To the boundaries: By default, the maximum values of w and h are used to crop the image. You can ignore the w and h parameters.

  The URL used to process the image: http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/crop,x_800,y_50![裁剪1](https://help-static-aliyun-doc.aliyuncs.com/assets/img/en-US/2940107071/p674602.jpg)

- Crop an area of 300 × 300 pixels from the starting point (800, 500)

  Configure the parameters based on the following requirements:

  - From the starting point (800, 500): `crop,x_800,y_500`.
  - An area of 300 × 300 pixels: `w_300,h_300`

  The URL used to process the image: http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/crop,x_800,y_500,w_300,h_300

  ![裁剪2](https://help-static-aliyun-doc.aliyuncs.com/assets/img/en-US/2940107071/p674612.jpg)

- Crop an area of 900 × 900 pixels in the lower-right corner of the source image

  Configure the parameters based on the following requirements:

  - From the starting point in the lower-right corner of the source image: `crop,g_se`
  - An area of 900 × 900 pixels: `w_900,h_900`

  The URL used to process the image: http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/crop,w_900,h_900,g_se

  ![裁剪3](https://help-static-aliyun-doc.aliyuncs.com/assets/img/en-US/2940107071/p674614.jpg)

- Crop an area of 900 × 900 pixels in the lower-right corner of an image and stretch the cropped area downward by (100, 100)

  Configure the parameters based on the following requirements:

  - From the starting point in the lower-right corner of the source image and stretch the cropped area downward by (100, 100): `crop,g_se,x_100,y_100`
  - An area of 900 × 900 pixels: `w_900,h_900`

  The URL used to process the image: http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/crop,x_100,y_100,w_900,h_900,g_se

  ![裁剪4](https://help-static-aliyun-doc.aliyuncs.com/assets/img/en-US/2940107071/p674615.jpg)