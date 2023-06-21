import { Injectable, OnModuleInit } from '@nestjs/common';
import { PrismaClient } from '@prisma/client';

@Injectable()
export class PrismaService extends PrismaClient implements OnModuleInit {

    async onModuleInit() {
        await this.$connect();
    }


    //evita deixar a conexao aberta durante o hot reload no desenvolvimento
    async enableShutdownHooks(app: any) {
        this.$on('beforeExit', async () => {
            await app.close();
        })
    }
}
