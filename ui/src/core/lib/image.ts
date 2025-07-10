export function generateThumb(file: File, width = 100, height = 100) {
  return new Promise((resolve) => {
    const reader = new FileReader()

    reader.onload = function (e) {
      const img = new Image()

      img.onload = function () {
        const canvas = document.createElement('canvas')
        const ctx = canvas.getContext('2d')
        const imgWidth = img.width
        const imgHeight = img.height

        canvas.width = width
        canvas.height = height

        ctx?.drawImage(
          img,
          imgWidth > imgHeight ? (imgWidth - imgHeight) / 2 : 0,
          0, // top aligned
          // imgHeight > imgWidth ? (imgHeight - imgWidth) / 2 : 0,
          imgWidth > imgHeight ? imgHeight : imgWidth,
          imgWidth > imgHeight ? imgHeight : imgWidth,
          0,
          0,
          width,
          height,
        )

        return resolve(canvas.toDataURL(file.type))
      }

      img.src = e.target?.result as string
    }

    reader.readAsDataURL(file)
  })
}