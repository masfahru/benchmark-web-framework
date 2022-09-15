const express = require('express');
const bodyParser = require('body-parser');

const app = express();
app.set('etag', false);

app.get('/', (req, res) => {
  res.setHeader('Content-Type', 'application/json; charset=utf-8');
  res.send({ status: 'OK' });
});
app.post('/', bodyParser.json(), (req, res) => {
  res.setHeader('Content-Type', 'application/json; charset=utf-8');
  res.send(req.body);
});

app.listen(3000);
