import { Controller, Post, Body } from '@nestjs/common';
import { AssetsService } from './assets.service';

@Controller('assets')
export class AssetsController {
  constructor(private readonly assetsService: AssetsService) {
  }

  @Post()
  create(@Body() body: {id: string, symbol: string, price: number}) {
    return this.assetsService.create(body);
  }
}
