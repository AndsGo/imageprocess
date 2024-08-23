# Resize images

## Usage notes

- Limits on source images

  - The following image formats are supported: JPG, PNG, BMP, GIF, WebP, and TIFF.

  - The width or height of a source image cannot exceed 30,000 pixels. The total number of pixels of a source image cannot exceed 250 million.

    The total number of pixels of a dynamic image, such as a GIF image, is calculated by using the following formula: `Width × Height × Number of image frames`. The total number of pixels of a static image, such as a PNG image, is calculated by using the following formula: `Width × Height`.

- Limits on resized images

  The width or height of a resized image cannot exceed 16,384 pixels. The total number of pixels of a resized image cannot exceed 16,777,216.

- Resizing priorities

  If you specify parameters both for resizing based on the width and height and for proportionally resizing in a URL, the image is resized based on the specified width and height.

- Image resizing based on the specified width or height

  - The source image is proportionally resized when proportional resizing is performed. For example, if you resize the height of a source image of 200 × 100 pixels to 100 pixels, the width of the source image is resized to 50 pixels.
  - The source image is resized based on the specified width or height. For example, if you resize the height of a source image of 200 × 100 pixels to 100 pixels, the width of the source image is also resized to 100 pixels.

- If the size of the resized image is larger than the size of the source image, the source image is returned. You can add the `limit_0` parameter to enlarge the image. Example: `http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/resize,w_500,limit_0`.

## Parameter description

Action: `resize`

### Resize an image based on the specified height and width

- Parameters

  | **Parameter** | **Required**                          | **Description**                                              | **Value range**                                              |
  | ------------- | ------------------------------------- | ------------------------------------------------------------ | ------------------------------------------------------------ |
  | **m**         | Yes                                   | Specifies the type of the resize action.                     | lfit: OSS proportionally resizes the source image as large as possible in a rectangle based on the specified width and height. This is the default value.mfit: OSS proportionally resizes the source image as small as possible outside a rectangle based on the specified width and height.fill: OSS proportionally resizes the source image as small as possible outside a rectangle, and then crops the resized image from the center based on the specified width and height.pad: OSS resizes the source image as large as possible in a rectangle based on the specified width and height, and fills the empty space with a specific color.fixed: OSS forcibly resizes the source image based on the specified width and height. |
  | **w**         | No                                    | Specifies the width to which you want to resize the image.   | [1,16384]                                                    |
  | **h**         | No                                    | Specifies the height to which you want to resize the image.  | [1,16384]                                                    |
  | **color**     | Yes (only when the value of m is pad) | If you set the resizing type to pad, you can select a color to fill the empty space. | RGB color values. For example, 000000 indicates black, and FFFFFF indicates white.Default value: FFFFFF (white). |

- Examples

  The size of the source image is 200 × 100 pixels. The w parameter is set to 150, and the h parameter is set to 80. The source image is resized to different sizes when you specify different resize types.

  - Proportional resizing: The aspect ratio of the resized image must be equal to the aspect ratio of the source image. If the width of the resized image is 150 pixels, the height of the resized image is 75 pixels. If the height of the resized image is 80 pixels, the width of the resized image is 160 pixels.
  - Maximum image size in a rectangle based on the specified width and height: The width and height of the resized image cannot exceed 150 pixels and 80 pixels, respectively.

  In this case, the size of the resized image is 150 × 75 pixels.

  ![lfit](https://help-static-aliyun-doc.aliyuncs.com/assets/img/en-US/7812863061/p137017.png)

### Resize an image by percentage

| **Parameter** | **Required** | **Description**                                              | **Value range**                                              |
| ------------- | ------------ | ------------------------------------------------------------ | ------------------------------------------------------------ |
| p             | Yes          | Specifies the percentage by which you want to resize the image. | [1,1000]A value smaller than 100 specifies that the image size is reduced. A value greater than 100 specifies that the image is enlarged. |

## Examples

You can set image processing parameters through file URL and API. This article takes file URL as an example. This article uses examples/example.jpg.

This test is based on  [comprehensive example](../README_en.md#comprehensive)

The image access address is:

http://127.0.0.1:8080/file/example.jpg

![原图](https://help-static-aliyun-doc.aliyuncs.com/assets/img/en-US/6342107071/p527167.jpg)

- Proportionally resize the image

  - Based on the width or height

    Configure the following parameters to resize the image:

    - Resize the source image to a height of 100 pixels: resize,h_100
    - Set the resize type to lfit: `m_lfit`

    The URL used to apply the preceding parameters is http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/resize,h_100,w_200,m_lfit.![break](https://help-static-aliyun-doc.aliyuncs.com/assets/img/en-US/6342107071/p527171.jpg)

    The source image is proportionally resized from 400 × 300 pixels to 133 × 100 pixels.

- Resize the image based on the specified width and height

  Configure the following parameters to resize the image:

  - Resize the source image to 100 × 100 pixels: `resize,h_100,w_100`
  - Set the resize type to fixed: `m_fixed`

  The URL used to apply the preceding parameters is http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/resize,m_fixed,h_100,w_100.![宽高缩放](https://help-static-aliyun-doc.aliyuncs.com/assets/img/en-US/6342107071/p527175.jpg)

  The source image is resized from 400 × 300 pixels to 100 × 100 pixels and the image obtained after the resize operation is distorted.

- Crop the image based on the specified width and height

  Configure the following parameters to resize the image:

  - Resize the source image to 100 × 100 pixels: `resize,h_100,w_100`
  - Set the resize type to fill: `m_fill`

  The URL used to apply the preceding parameters is http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/resize,m_fill,h_100,w_100.![自动裁剪](https://help-static-aliyun-doc.aliyuncs.com/assets/img/en-US/6342107071/p527179.jpg)

  The source image is first resized from 400 × 300 pixels to 133 × 100 pixels and cropped by width from center to produce a final image that contains 100 × 100 pixels.

- Resize the source image based on the specified width and height, and fill the empty space

  Configure the following parameters to resize the image:

  - Resize the source image to 100 × 100 pixels: `resize,h_100,w_100`
  - Set the resize type to pad: `m_pad`
  - Fill the empty space with red: `color_FF0000`

  The URL used to apply the preceding parameters is http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/resize,m_pad,h_100,w_100,color_FF0000.

  ![填充红色](https://help-static-aliyun-doc.aliyuncs.com/assets/img/en-US/6342107071/p527183.jpg)

  The source image is resized from 400 × 300 pixels to 100 × 75 pixels and the height of the resulting image is enlarged from center to 100 pixels and filled with red in the empty space.

- Resize an image based on the specified percentage

  Configure the following parameters to resize the image:

  Resize the source image by 50%: `resize,p_50`

  The URL used to apply the preceding parameters is http://127.0.0.1:8080/file/example.jpg?x-oss-process=image/resize,p_50.

  ![按比例缩放](https://help-static-aliyun-doc.aliyuncs.com/assets/img/en-US/6342107071/p527188.jpg)

  The source image is proportionally resized from 400 × 300 pixels to 200 × 150 pixels. The image obtained after the resize operation is half the size of the source image.