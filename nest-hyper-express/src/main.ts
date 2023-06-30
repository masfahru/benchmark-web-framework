import { NestFactory } from '@nestjs/core';
import { AppModule } from './app.module';
import { HyperExpressAdapter } from './hyper-express-adapter';
import { Logger } from '@nestjs/common';

async function bootstrap() {
  const now = Date.now();
  const app = await NestFactory.create(
    AppModule,
    new HyperExpressAdapter({
      fast_buffers: true,
    }),
  );
  await app.listen(3000, '0.0.0.0');
  Logger.log(
    `Server running on port 3000, initialize time: ${Date.now() - now}ms`,
    'Bootstrap',
  );
}
bootstrap();
