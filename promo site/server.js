const express = require('express')
const app = express()
const port = 3000
var path = require('path');

app.use(express.static(__dirname));

app.listen(port, () => {
  console.log(`Example app listening on port {port}`)
})