import { serve } from 'https://deno.land/std/http/server.ts';

const headers = new Headers({
  'Content-Type': 'application/json; charset=UTF-8',
  'Connection': 'Keep-Alive',
  'Keep-Alive': 'timeout=72',
});

async function handler(req: Request) {
  switch (req.method) {
    case 'GET':
      return new Response(JSON.stringify({ status: "OK" }), {
        status: 200,
        headers
      });
      break;
    case 'POST':
      const data = await req.json()
      return new Response(JSON.stringify(data), {
        status: 200,
        headers
      });
      break;
    default:
      break;
  }
}

await serve(handler, { port: 3000 });