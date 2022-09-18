Bun.serve({
  async fetch(req: Request) {
    switch (req.method) {
      case 'GET':
        return new Response(JSON.stringify({ status: 'OK' }), {
          status: 200,
          headers: {
            'Content-Type': 'application/json; charset=UTF-8',
            Connection: 'Keep-Alive',
            'Keep-Alive': 'timeout=72',
          },
        });
        break;
      case 'POST':
        const data = await req.json();
        return new Response(JSON.stringify(data), {
          status: 200,
          headers: {
            'Content-Type': 'application/json; charset=UTF-8',
            Connection: 'Keep-Alive',
            'Keep-Alive': 'timeout=72',
          },
        });
        break;
      default:
        break;
    }
  },
  port: 3000,
});
