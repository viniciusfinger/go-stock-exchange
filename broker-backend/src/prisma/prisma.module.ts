import { Global, Module } from '@nestjs/common';
import { PrismaService } from './prisma/prisma.service';

@Global() //evita ter que importar esse módulo, já que será usado em vários outros módulos
@Module({
  providers: [PrismaService],
  exports: [PrismaService]
})
export class PrismaModule {}