import { Controller, Get, Post, Req } from '@nestjs/common';

@Controller()
export class AppController {
  @Get()
  getHello() {
    return {
      status: 'OK',
    };
  }

  @Post()
  postHello(@Req() request: any) {
    return request.body.sort((a, b) => a.age - b.age);;
  }
}
