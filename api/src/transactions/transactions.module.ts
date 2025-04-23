import { Module } from '@nestjs/common';
import { TransactionsService } from './transactions.service';
import { TransactionsController } from './transactions.controller';
import { ClientsModule, Transport } from '@nestjs/microservices';

@Module({
  providers: [TransactionsService],
  controllers: [TransactionsController],
  imports: [
    ClientsModule.register([
      {
        name: 'KAFKA_SERVICE',
        transport: Transport.KAFKA,
        options: {
          client: {
            clientId: 'API',
            brokers: ['localhost:19092', 'localhost:19093'],
          },
          consumer: {
            groupId: 'api-group',
          },
          producer: {
            allowAutoTopicCreation: true,
          },
        },
      },
    ]),
  ],
})
export class TransactionsModule {}
