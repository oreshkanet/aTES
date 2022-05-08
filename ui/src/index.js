const swaggerUI = require('swagger-ui-express');
const express = require('express');
const YAML = require('yamljs');
const config = require('./config');

const authDocument = YAML.load('./openapi3-auth.yaml');

const app = express()
app.use('/auth', swaggerUI.serve, swaggerUI.setup(authDocument));

const server = app.listen(config.port, () => {
  const host = server.address().address;
  const port = server.address().port;
  console.log(`Запущен сервер http://${host}:${port}`);
});

process.on('SIGTERM', () => {
  const host = server.address().address;
  const port = server.address().port;
  console.log(`Остановка сервера http://${host}:${port}`);
  server.stop();
});
  