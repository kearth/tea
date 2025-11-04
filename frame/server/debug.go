package server

// RedocUI Redoc UI模板
const RedocUI = `<!doctype html>
<html lang="en">
  <head>
    <meta charset="UTF-8" />
    <title>Redoc UI</title>
  </head>
  <body>
    <redoc spec-url="{SwaggerUIDocUrl}"></redoc>
    <script src="https://cdn.redoc.ly/redoc/latest/bundles/redoc.standalone.js"> </script>
  </body>
</html>`
