const swaggerUI = require('swagger-ui-express');
const express = require('express');
const YAML = require('yamljs');
const config = require('./config');

const authDocument = YAML.load('./openapi3/auth.yaml');
const taskDocument = YAML.load('./openapi3/auth.yaml');
const billingDocument = YAML.load('./openapi3/auth.yaml');
const analyticsDocument = YAML.load('./openapi3/auth.yaml');

const app = express();

app.use('/auth', function(req, res, next){
  authDocument.servers[0].url = `http://localhost:${config.auth_port}`;
  req.swaggerDoc = authDocument;
  next();
}, swaggerUI.serve, swaggerUI.setup());

app.use('/task', function(req, res, next){
  taskDocument.servers[0].url = `http://localhost:${config.tasks_port}`;
  req.swaggerDoc = taskDocument;
  next();
}, swaggerUI.serve, swaggerUI.setup());

app.use('/billing', function(req, res, next){
  billingDocument.servers[0].url = `http://localhost:${config.billing_port}`;
  req.swaggerDoc = billingDocument;
  next();
}, swaggerUI.serve, swaggerUI.setup());

app.use('/analytics', function(req, res, next){
  analyticsDocument.servers[0].url = `http://localhost:${config.analytic_port}`;
  req.swaggerDoc = analyticsDocument;
  next();
}, swaggerUI.serve, swaggerUI.setup());

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
  