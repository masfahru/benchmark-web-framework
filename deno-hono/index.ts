import { serve } from 'https://deno.land/std/http/server.ts';
import { Hono } from 'https://deno.land/x/hono/mod.ts';
import { prettyJSON } from 'https://deno.land/x/hono/middleware.ts';

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

await serve(app.fetch, { port: 3000 });
