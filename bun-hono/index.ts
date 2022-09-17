import { Hono } from 'hono';
import { prettyJSON } from 'hono/pretty-json';

const app = new Hono();

app.use('*', prettyJSON());

app.get('/', (c) => {
  c.header('Connection', 'Keep-Alive');
  c.header('Keep-Alive', 'timeout=72');
  return c.json({ status: 'OK' });
});

app.post('/', async (c) => {
  const data = await c.req.json();
  c.header('Connection', 'Keep-Alive');
  c.header('Keep-Alive', 'timeout=72');
  return c.json(data);
});

export default {
  port: 3000,
  fetch: app.fetch,
};
