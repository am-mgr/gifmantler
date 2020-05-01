# GifMantler

[![Build Status](https://dev.azure.com/am-mgr/gifmantler/_apis/build/status/am-mgr.gifmantler?branchName=master)](https://dev.azure.com/am-mgr/gifmantler/_build/latest?definitionId=1&branchName=master)

You can use **gifmantler** to unpack a GIF to its individual images.

#### Usage:

1. To unpack a GIF. The images will be generated in a **gif_out** folder at GIF location.

    To PNG images:
    ```
    gifmantler -png -gif foo.gif
    ```

    To JPEG images:

    *The flag **jpeg-quality** controls the quality of jpeg output ranging from 1 to 100. The higher, the better quality*
    ```
    gifmantler -jpeg -jpeg-quality 80 -gif foo.gif
    ```

2. To check version
    ```
    gifmantler -version
    ```
