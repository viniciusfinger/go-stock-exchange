import { Injectable } from '@nestjs/common';
import { PrismaService } from 'src/prisma/prisma/prisma.service';

@Injectable()
export class AssetsService {

    constructor(private prismaService : PrismaService) {

    }

    create(data: {id: string, symbol: string, price: number}) {
        //O prisma gera uma classe dentro do node_modules com o nome do schema que criamos (no caso "asset")
        return this.prismaService.asset.create({data});
    }

}
