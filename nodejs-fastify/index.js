const fastify = require('fastify');
const app = fastify();

app.get('/', (request, reply) => {
  reply
    .code(200)
    .header('Content-Type', 'application/json; charset=utf-8')
    .send({
      status: 'OK',
    });
});
app.post('/', (request, reply) => {
  reply
    .code(200)
    .header('Content-Type', 'application/json; charset=utf-8')
    .send(request.body);
});

app.listen({ port: 3000, host: '0.0.0.0' });
