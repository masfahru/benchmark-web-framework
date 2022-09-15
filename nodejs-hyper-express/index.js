const HyperExpress = require('hyper-express');

const app = new HyperExpress.Server();
app.get('/', (request, response) => {
  response
    .status(200)
    .header('Connection', 'Keep-Alive')
    .header('Keep-Alive', 'timeout=72')
    .header('Content-Type', 'application/json; charset=utf-8')
    .send(JSON.stringify({ status: 'OK' }));
});
app.post('/', async (request, response) => {
  const data = await request.json();
  response
    .status(200)
    .header('Connection', 'Keep-Alive')
    .header('Keep-Alive', 'timeout=72')
    .header('Content-Type', 'application/json; charset=utf-8')
    .send(JSON.stringify(data));
});

app.listen(3000);
